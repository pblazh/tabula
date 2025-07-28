package functions

import (
	"math"
)

func sum[T Number](values ...T) T {
	var result T
	for _, n := range values {
		result += n
	}
	return result
}

func product[T Number](values ...T) T {
	result := T(1)
	for _, n := range values {
		result *= n
	}
	return result
}

func average[T Number](values ...T) T {
	if len(values) == 0 {
		return T(0)
	}

	total := T(0)
	for _, n := range values {
		total += n
	}
	return T(total / T(len(values)))
}

func abs[T Number](values ...T) T {
	return T(math.Abs(float64(values[0])))
}

func roundPrecise(up bool, value, precision float64) float64 {
	if up {
		return math.Ceil(value/precision) * precision
	}

	return math.Floor(value/precision) * precision
}
