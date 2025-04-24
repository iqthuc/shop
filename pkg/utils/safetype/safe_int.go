package safetype

import (
	"errors"
	"math"
)

var ErrOverflow = errors.New("value out of int32 bounds")

func SafeIntToInt32(i int) (int32, error) {
	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0, ErrOverflow
	}

	return int32(i), nil
}
