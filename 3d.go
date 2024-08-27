package multilateration

import (
	"fmt"
	"math"

	"github.com/davidkleiven/gononlin/nonlin"
	"gonum.org/v1/gonum/stat/combin"
)

var solver3D = nonlin.NewtonKrylov{
	Maxiter:  1000,
	StepSize: 0.01,
	Tol:      0.001, // 1 mm
}

type pos3D interface {
	XYZ() Pos3D
}

type Pos3D struct {
	X, Y, Z float64 // in meters
}

func (p Pos3D) XYZ() Pos3D {
	return p
}

type TimePos3D struct {
	T       float64 // in seconds
	X, Y, Z float64 // in meters
}

func (t TimePos3D) XYZ() Pos3D {
	return Pos3D{t.X, t.Y, t.Z}
}

func Multilaterate3D(pos ...TimePos3D) (TimePos3D, error) {
	if len(pos) < 4 {
		return TimePos3D{}, ErrNotEnoughPoints
	}
	combs := combin.Combinations(len(pos), 4)
	ref := pos[0].T
	var averaged TimePos3D
	var averagedT float64
	for _, comb := range combs {
		res, err := multilaterate3D(pos[comb[0]], pos[comb[1]], pos[comb[2]], pos[comb[3]])
		if err != nil {
			return TimePos3D{}, fmt.Errorf("could not calculate with points (%d, %d, %d, %d): %v", comb[0], comb[1], comb[2], comb[3], err)
		}
		averaged.X += res.X
		averaged.Y += res.Y
		averaged.Z += res.Z
		averagedT += ref - res.T
	}
	averaged.X /= float64(len(combs))
	averaged.Y /= float64(len(combs))
	averaged.Z /= float64(len(combs))
	averagedT /= float64(len(combs))
	averaged.T = ref - averagedT
	return averaged, nil
}

func multilaterate3D(p1, p2, p3, p4 TimePos3D) (TimePos3D, error) {
	t_12 := p1.T - p2.T
	t_13 := p1.T - p3.T
	t_14 := p1.T - p4.T
	p := nonlin.Problem{
		F: func(out, x []float64) {
			p := Pos3D{x[0], x[1], x[2]}
			out[0] = dist3(p1, p) - dist3(p2, p) - t_12*c
			out[1] = dist3(p1, p) - dist3(p3, p) - t_13*c
			out[2] = dist3(p1, p) - dist3(p4, p) - t_14*c
		},
	}

	avgX := (p1.X + p2.X + p3.X) / 3
	avgY := (p1.Y + p2.Y + p3.Y) / 3
	avgZ := (p1.Z + p2.Z + p3.Z) / 3
	res, err := solver3D.Solve(p, []float64{avgX, avgY, avgZ})
	if err != nil {
		return TimePos3D{}, err
	}
	pos := TimePos3D{X: res.X[0], Y: res.X[1], Z: res.X[2]}
	dist := dist3(pos, p1)
	t := p1.T - mToD(dist)
	pos.T = t
	return pos, nil
}

func Simulate3D(source TimePos3D, stations ...Pos3D) []TimePos3D {
	res := make([]TimePos3D, len(stations))
	for i, s := range stations {
		res[i] = TimePos3D{
			T: source.T + mToD(dist3(source, s)),
			X: s.X,
			Y: s.Y,
			Z: s.Z,
		}
	}
	return res
}

func dist3(p1, p2 pos3D) float64 {
	pp1, pp2 := p1.XYZ(), p2.XYZ()
	return math.Sqrt(math.Pow(pp1.X-pp2.X, 2) + math.Pow(pp1.Y-pp2.Y, 2) + math.Pow(pp1.Z-pp2.Z, 2))
}
