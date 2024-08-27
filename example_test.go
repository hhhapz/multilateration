package multilateration_test

import (
	"fmt"

	multilat "github.com/hhhapz/multilateration"
)

func Example2D() {
	n := 0.0 // source emission time
	source := multilat.TimePos2D{T: n, X: 50, Y: 50}

	// generate positions based on source and stations. This simulates when stations will
	// receive the signal, which we can then use multilateration to retrieve the source.
	positions := multilat.Simulate2D(source, multilat.Pos2D{0, 0}, multilat.Pos2D{0, 100}, multilat.Pos2D{100, 0})

	fmt.Printf("source: (%.2f, %.2f)\n", source.X, source.Y)
	for _, p := range positions {
		fmt.Printf("station: (%.2f, %.2f)\n", p.X, p.Y)
	}

	source, err := multilat.Multilaterate2D(positions...)
	if err != nil {
		fmt.Printf("error: could not multilaterate: %v\n", err)
		return
	}
	fmt.Printf("multilaterated: (%.2f, %.2f)\n", source.X, source.Y)
	// Output:
	// source: (50.00, 50.00)
	// station: (0.00, 0.00)
	// station: (0.00, 100.00)
	// station: (100.00, 0.00)
	// multilaterated: (50.00, 50.00)
}
