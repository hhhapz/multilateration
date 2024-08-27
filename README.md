# Multilateration Library

This Go library provides functionality for **multilateration**, a technique used to determine the position of an object based on time difference of arrival (TDOA) measurements from multiple known positions (sensors).

For more information on multilateration, see the [Wikipedia article on Multilateration](https://en.wikipedia.org/wiki/Multilateration).

The api is based on meters for positions, and has nanosecond precision.

## Installation

```bash
go get github.com/hhhapz/multilateration
```

## Example

```go
import multilat "github.com/hhhapz/multilateration"

func main() {
	n := time.Now() // source emission time
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
	fmt.Printf("multilaterated: (%.2f, %.2f) with time diff %s\n", source.X, source.Y, n.Sub(source.T))
	// Output:
	// source: (50.00, 50.00)
	// station: (0.00, 0.00)
	// station: (0.00, 100.00)
	// station: (100.00, 0.00)
	// multilaterated: (50.00, 50.00) with time diff 0s
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
