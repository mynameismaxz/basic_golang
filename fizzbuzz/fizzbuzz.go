package fizzbuzz

import (
	"strconv"
)

func New(n int) Object {
	return Object{
		n: n,
	}
}

type Object struct {
	n int
}

// receiver method
func (o Object) String() string {
	return FizzBuzz(o.n)
}

func FizzBuzz(n int) string {
	var s string
	switch {
	case n%15 == 0:
		s = "FizzBuzz"
	case n%3 == 0:
		s = "Fizz"
	case n%5 == 0:
		s = "Buzz"
	default:
		s = strconv.Itoa(n)
	}
	return s
}
