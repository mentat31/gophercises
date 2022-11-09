package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var correct int
var incorrect int

func main() {

	var nFlag = flag.Int("n", 30, "seconds for timer")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*nFlag) * time.Second)

	fmt.Print("Quiz Time! Ready?! [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "n" {
			fmt.Println("Goodbye")
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
