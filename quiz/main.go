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

// Scores to return
var correct int
var incorrect int

func main() {
	// Timer flag, defaults to 30 seconds
	var nFlag = flag.Int("n", 30, "seconds for timer")
	flag.Parse()
	timer := time.NewTimer(time.Duration(*nFlag) * time.Second)

	fmt.Print("Quiz Time! Ready?! [y/n]: ")
	scanner := bufio.NewScanner(os.Stdin)

	// Read stdin, Exits when n is passed, break loop otherwise
	for scanner.Scan() {
		if scanner.Text() == "n" {
			fmt.Println("Goodbye")
			os.Exit(3)
		} else {
			break
		}
	}
	// Timer starts in seperate goroutine. Running <-timer.C w/o caused various problems.
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
	/*
	 Because there is no modification to the file and is only used for reading/scoring anweres, The file is
	 is not loaded into memory, as we only need to keep score.
	*/
	defer f.Close()

	csvReader := csv.NewReader(f)
	// for loop through questions.
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("What is %v?\n", rec[0])
		// Replace previous input with new at same address
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
