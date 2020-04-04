package simplemath

import (
	"errors"
)

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Sum(values ...float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	}
	answer = p1 / p2
	return
}
