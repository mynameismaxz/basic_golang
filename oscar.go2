package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./testdata/oscar_age_male.csv")
	if err != nil {
		log.Fatal("Cannot open this file.")
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read file in format .csv")
	}

	actorPopularName := make(map[string]bool)

	for _, v := range records {
		// show name of Actor Oscar
		name := v[3]
		if _, exist := actorPopularName[name]; exist {
			continue
		} else {
			actorPopularName[name] = true
		}
	}
}
