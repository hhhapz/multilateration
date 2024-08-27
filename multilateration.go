package multilateration

import (
	"errors"
)

const c = 299_792_458 // m/s

// meters to s returns the duration it takes light to travel x meters in seconds
func mToD(x float64) float64 {
	return x / c
}

var ErrNotEnoughPoints = errors.New("not enough points to multilaterate")
