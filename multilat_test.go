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

func Test3D(t *testing.T) {
	n := 0.0 // time of emission
	var tests = []struct {
		source   Pos3D
		stations []Pos3D
	}{
		{
			Pos3D{50, 50, 50},
			[]Pos3D{{0, 0, 0}, {0, 100, 0}, {100, 0, 0}, {150, 50, 100}},
		},
		{
			Pos3D{70, 30, 20},
			[]Pos3D{{0, 0, 0}, {0, 100, 0}, {100, 0, 0}, {80, 60, 50}},
		},
		{
			Pos3D{-50, 0, 25},
			[]Pos3D{{-100, -100, -100}, {-100, 100, 50}, {100, -100, 75}, {-150, 50, -50}},
		},
		{
			Pos3D{50, 0, 100},
			[]Pos3D{{0, 0, 0}, {100, 0, 0}, {50, 86.6, 200}, {75, -50, 150}},
		},
		{
			Pos3D{10, 10, 10},
			[]Pos3D{{0, 0, 0}, {100, 0, 0}, {50, 100, 100}, {20, 50, 75}},
		},
		{
			Pos3D{1000, 1000, 500},
			[]Pos3D{{0, 0, 0}, {100, 0, 0}, {1500, 1500, 50}, {2000, 1500, 800}},
		},
		{
			Pos3D{1000, 1000, 1000},
			[]Pos3D{{0, 0, 0}, {100, 0, 0}, {50, 100, 0}, {500, 200, 1800}},
		},
		{
			Pos3D{1000, 1000, 2000},
			[]Pos3D{{0, 0, 0}, {100, 0, 0}, {0, 3500, -5000}, {3500, 2500, 1500}},
		},
		{
			Pos3D{932, 932, 932},
			[]Pos3D{{0, 0, 0}, {1800, 220, 300}, {500, 2500, 700}, {1500, 1500, 1500}},
		},
		{
			Pos3D{25000, 25000, 15000},
			[]Pos3D{{300, 5000, 1000}, {53125, 0, 2000}, {30123, 26294, 3000}, {40000, 20000, 25000}},
		},
		{
			Pos3D{20000, 15000, 10000},
			[]Pos3D{{300, 5000, 500}, {53125, 0, 2500}, {30123, 26294, 1500}, {25000, 10000, 12000}},
		},
	}
	for _, test := range tests {
		source := TimePos3D{T: n, X: test.source.X, Y: test.source.Y, Z: test.source.Z}
		positions := Simulate3D(source, test.stations...)
		t.Logf("source: (%.2f, %.2f, %.2f)", source.X, source.Y, source.Z)
		var stationsLog []string
		for _, p := range positions {
			stationsLog = append(stationsLog, fmt.Sprintf("(%.2f, %.2f, %.2f)", p.X, p.Y, p.Z))
		}
		t.Logf("stations: %s", strings.Join(stationsLog, ", "))
		source, err := Multilaterate3D(positions...)
		if err != nil {
			t.Errorf("could not multilaterate: %v", err)
		}
		t.Logf("multilaterated: (%.2f, %.2f, %.2f) with time diff %v", source.X, source.Y, source.Z, n-source.T)
		t.Log()
	}
}
