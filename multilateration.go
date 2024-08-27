package multilateration

import (
	"errors"
	"time"
)

const c = 299_792_458 // m/s

// meters to ns returns the duration it takes light to travel x meters in nanoseconds (time.Duration)
func mToD(x float64) time.Duration {
	return time.Duration(x / c * 1e9)
}

var ErrNotEnoughPoints = errors.New("not enough points to multilaterate")
