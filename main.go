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
	var filename string
	// read input file from argument
	if len(os.Args) < 2 {
		// when no file than use default
		filename = "./sample_input"
	} else {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		hand := string(line[:14])
		deck := string(line[15:])

		bestHand := FindBestHand(hand, deck)
		fmt.Printf("Hand: %s Deck: %s Best hand: %s\n", hand, deck, bestHand)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
