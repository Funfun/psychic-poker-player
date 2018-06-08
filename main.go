package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	var inputFileName string
	var outputFileName string

	// read input file from argument
	if len(os.Args) < 2 {
		// when no file than use default
		inputFileName = "./sample_input"
		outputFileName = "./output"
	} else {
		inputFileName = os.Args[1]
		outputFileName = os.Args[2]
	}

	fi, err := os.Open(inputFileName)
	check(err)

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	fo, err := os.Create(outputFileName)
	check(err)
	w := bufio.NewWriter(fo)

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		hand := string(line[:14])
		deck := string(line[15:])

		bestHand := FindBestHand(hand, deck)
		output := fmt.Sprintf("Hand: %s Deck: %s Best hand: %s\n", hand, deck, bestHand)
		if _, err := w.WriteString(output); err != nil {
			log.Fatalf("Error while writing to a file. Error: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}

}
