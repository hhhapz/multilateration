package multilateration

import (
	"fmt"
	"strings"
	"testing"
)

func Test2D(t *testing.T) {
	n := 0.0 // time of emission
	var tests = []struct {
		source   Pos2D
		stations []Pos2D
	}{
		{
			Pos2D{50, 50},
			[]Pos2D{{0, 0}, {0, 100}, {100, 0}},
		},
		{
			Pos2D{70, 30},
			[]Pos2D{{0, 0}, {0, 100}, {100, 0}},
		},
		{
			Pos2D{-50, 0},
			[]Pos2D{{-100, -100}, {-100, 100}, {100, -100}},
		},
		{
			Pos2D{50, 0},
			[]Pos2D{{0, 0}, {100, 0}, {50, 86.6}},
		},
		{
			Pos2D{10, 10},
			[]Pos2D{{0, 0}, {100, 0}, {50, 100}},
		},
		{
			Pos2D{1000, 1000},
			[]Pos2D{{0, 0}, {100, 0}, {50, 100}},
		},
		{
			Pos2D{1000, 1000},
			[]Pos2D{{0, 0}, {100, 0}, {50, 100}, {600, 200}},
		},
		{
			Pos2D{1000, 1000},
			[]Pos2D{{0, 0}, {100, 0}, {50, 100}, {3000, 2000}},
		},
		{
			Pos2D{932, 932},
			[]Pos2D{{0, 0}, {1800, 220}, {500, 2500}},
		},
		{
			Pos2D{25000, 25000},
			[]Pos2D{{300, 5000}, {53125, 0}, {30123, 26294}},
		},
		{
			Pos2D{20000, 15000},
			[]Pos2D{{300, 5000}, {53125, 0}, {30123, 26294}},
		},
	}
	for _, test := range tests {
		source := TimePos2D{T: n, X: test.source.X, Y: test.source.Y}
		positions := Simulate2D(source, test.stations...)
		t.Logf("source: (%.2f, %.2f)", source.X, source.Y)
		var stationsLog []string
		for _, p := range positions {
			stationsLog = append(stationsLog, fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y))
		}
		t.Logf("stations: %s", strings.Join(stationsLog, ", "))
		source, err := Multilaterate2D(positions...)
		if err != nil {
			t.Errorf("could not multilaterate: %v", err)
		}
		t.Logf("multilaterated: (%.2f, %.2f) with time diff %v", source.X, source.Y, n-source.T)
	}

}
