package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var correct int
var incorrect int

func main() {

	f, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("What is %v?\n", rec[0])

		var input string
		fmt.Scanln(&input)

		if input == rec[1] {
			correct += 1
		} else {
			incorrect += 1
		}

	}
	fmt.Println("Correct:", correct, "Incorrect:", incorrect)

}
