package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type country struct {
	country   string
	latitude  float64
	longitude float64
	name      string
}

func main() {
	countries := readCsvFile("./countries.csv")
	fmt.Println(countries)
}

func readCsvFile(filePath string) map[string]country {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var countries = make(map[string]country)
	for _, record := range records {
		var country country
		country.country = record[0]
		country.latitude, _ = strconv.ParseFloat(record[1], 64)
		country.longitude, _ = strconv.ParseFloat(record[2], 64)
		country.name = record[3]
		countries[country.name] = country
	}

	return countries
}
