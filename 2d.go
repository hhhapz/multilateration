package multilateration

import (
	"fmt"
	"math"

	"github.com/davidkleiven/gononlin/nonlin"
	"gonum.org/v1/gonum/stat/combin"
)

var solver2D = nonlin.NewtonKrylov{
	Maxiter:  1000,
	StepSize: 0.01,
	Tol:      0.01, // 1 cm
}

type pos2D interface {
	XY() Pos2D
}

type Pos2D struct {
	X, Y float64 // in meters
}

func (p Pos2D) XY() Pos2D {
	return p
}

type TimePos2D struct {
	T    float64 // in seconds
	X, Y float64 // in meters
}

func (t TimePos2D) XY() Pos2D {
	return Pos2D{t.X, t.Y}
}

func Multilaterate2D(pos ...TimePos2D) (TimePos2D, error) {
	if len(pos) < 3 {
		return TimePos2D{}, ErrNotEnoughPoints
	}
	combs := combin.Combinations(len(pos), 3)
	ref := pos[0].T
	averaged := TimePos2D{X: 0, Y: 0}
	var averagedT float64
	for _, comb := range combs {
		res, err := multilaterate2D(pos[comb[0]], pos[comb[1]], pos[comb[2]])
		if err != nil {
			return TimePos2D{}, fmt.Errorf("could not calculate with points (%d, %d, %d): %v", comb[0], comb[1], comb[2], err)
		}
		averaged.X += res.X
		averaged.Y += res.Y
		averagedT += ref - res.T
	}
	averaged.X /= float64(len(combs))
	averaged.Y /= float64(len(combs))
	averagedT /= float64(len(combs))
	averaged.T = ref - averagedT
	return averaged, nil
}

func multilaterate2D(p1, p2, p3 TimePos2D) (TimePos2D, error) {
	t_12 := p1.T - p2.T
	t_13 := p1.T - p3.T
	p := nonlin.Problem{
		F: func(out, x []float64) {
			p := Pos2D{x[0], x[1]}
			out[0] = dist2(p1, p) - dist2(p2, p) - t_12*c
			out[1] = dist2(p1, p) - dist2(p3, p) - t_13*c
		},
	}

	avgX := (p1.X + p2.X + p3.X) / 3
	avgY := (p1.Y + p2.Y + p3.Y) / 3
	res, err := solver2D.Solve(p, []float64{avgX, avgY})
	if err != nil {
		return TimePos2D{}, err
	}
	pos := TimePos2D{X: res.X[0], Y: res.X[1]}
	dist := dist2(pos, p1)
	t := p1.T - mToD(dist)
	pos.T = t
	return pos, nil
}

func Simulate2D(source TimePos2D, stations ...Pos2D) []TimePos2D {
	res := make([]TimePos2D, len(stations))
	for i, s := range stations {
		res[i] = TimePos2D{
			T: source.T + mToD(dist2(source, s)),
			X: s.X,
			Y: s.Y,
		}
	}
	return res
}

func dist2(p1, p2 pos2D) float64 {
	pp1, pp2 := p1.XY(), p2.XY()
	return math.Sqrt(math.Pow(pp1.X-pp2.X, 2) + math.Pow(pp1.Y-pp2.Y, 2))
}
