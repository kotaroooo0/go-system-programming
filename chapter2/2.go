package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		[]string{"名前", "年齢", "身長", "体重"},
		[]string{"Tanaka", "31", "190cm", "97kg"},
		[]string{"Suzuki", "46", "180cm", "79kg"},
		[]string{"Matsui", "45", "188cm", "95kg"},
	}

	file, err := os.Create("sample.csv")
	if err != nil {
		panic(err)
	}

	fileWriter := csv.NewWriter(file)
	stdWriter := csv.NewWriter(os.Stdout)

	for _, record := range records {
		err = fileWriter.Write(record)
		err = stdWriter.Write(record)
		if err != nil {
			panic(err)
		}
	}

	fileWriter.Flush()
	stdWriter.Flush()
}
