package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	in := [][]string{{"id", "name"}, {"1", "daimori"}, {"2", "daimori2"}}
	file, err := os.Create("text.csv")
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(file)
	w.WriteAll(in)
	w.Flush()
}
