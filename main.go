package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mynameismaxz/basic_golang/fizzbuzz"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fizzbuzz.FizzBuzz(number))
}
