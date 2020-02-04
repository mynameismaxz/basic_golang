package main

import (
	"reflect"
	"testing"
)

type fakeRandomer string

func (r fakeRandomer) Intn(int) int {
	return 3
}

func TestFizzBuzzByRandom(t *testing.T) {
	r := fakeRandomer("")

	want := FizzBuzzRandomResponse{
		Number:  "3",
		Message: "Fizz",
	}

	get := fizzbuzzByRandom(r)

	if !reflect.DeepEqual(want, get) {
		t.Error("not equal")
	}
}
