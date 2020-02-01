package fizzbuzz

import (
	"testing"
)

func TestGivenOneSayOne(t *testing.T) {
	want := "1"
	given := 1

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenTwoSayTwo(t *testing.T) {
	want := "2"
	given := 2

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenThreeSayFizz(t *testing.T) {
	want := "Fizz"
	given := 3

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenFourSayFour(t *testing.T) {
	want := "4"
	given := 4

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenFiveSayBuzz(t *testing.T) {
	want := "Buzz"
	given := 5

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenSixSayFizz(t *testing.T) {
	want := "Fizz"
	given := 6

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenSevenSaySeven(t *testing.T) {
	want := "7"
	given := 7

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenEightSayEight(t *testing.T) {
	want := "8"
	given := 8

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenNineSayFizz(t *testing.T) {
	want := "Fizz"
	given := 9

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGivenTenSayBuzz(t *testing.T) {
	want := "Buzz"
	given := 10

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGiven12SayFizz(t *testing.T) {
	want := "Fizz"
	given := 12

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGiven15SayFizzBuzz(t *testing.T) {
	want := "FizzBuzz"
	given := 15

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGiven18SayFizz(t *testing.T) {
	want := "Fizz"
	given := 18

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGiven20SayBuzz(t *testing.T) {
	want := "Buzz"
	given := 20

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}

func TestGiven30SayFizzBuzz(t *testing.T) {
	want := "FizzBuzz"
	given := 30

	get := FizzBuzz(given)

	if want != get {
		t.Errorf("given %d want %q but get %q", given, want, get)
	}
}
