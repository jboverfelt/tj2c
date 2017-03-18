package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
)

type timeSeries struct {
	ActivitiesCalories []struct {
		DateTime string `json:"dateTime"`
		Value    string `json:"value"`
	} `json:"activities-calories"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatalln("Usage: tj2c <path to json file>")
	}

	f, err := os.Open(args[0])

	checkErr(err)

	var series timeSeries

	err = json.NewDecoder(f).Decode(&series)

	checkErr(err)

	w := csv.NewWriter(os.Stdout)

	err = w.Write([]string{"date", "value"})

	checkErr(err)

	for _, rec := range series.ActivitiesCalories {
		err = w.Write([]string{rec.DateTime, rec.Value})
		checkErr(err)
	}

	w.Flush()

	checkErr(w.Error())
}
