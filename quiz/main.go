package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var correct int
var incorrect int

func main() {

	timer := time.NewTimer(30 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Quiz Time! Ready?! [y/n]: ")

	for scanner.Scan() {
		if scanner.Text() != "y" {
			fmt.Printf("%T\n", "y")
			os.Exit(3)
		} else {
			break
		}
	}

	go func() {
		fmt.Println("Timer started!")
		<-timer.C
		fmt.Println("Outta Time!!")
		fmt.Println("Correct:", correct, "Incorrect:", incorrect)
		os.Exit(0)

	}()

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
