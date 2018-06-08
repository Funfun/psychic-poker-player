package main

import (
	"fmt"
	"sort"
	"strings"
)

// FindBestHand returns highest possible high-hand
// category from the given hand and deck sets
func FindBestHand(hand string, deck string) string {
	handCards := *NewCards(hand)
	deckCards := *NewCards(deck)

	return HandCategoryName(Lookup(handCards, deckCards))
}

func Lookup(cards Cards, deckCards Cards) int {
	var newRank int
	var bestRank int

	// assume initially the player has a good hand
	bestRank = FindHandCategory(cards)
	if HandCategoryRank("straight-flush") == bestRank {
		return HandCategoryRank("straight-flush")
	}

	// check all deck cards
	newRank = FindHandCategory(deckCards)
	if HandCategoryRank("straight-flush") == newRank {
		return HandCategoryRank("straight-flush")
	}

	if newRank < bestRank {
		bestRank = newRank
	}

	n := len(cards)
	for j := 1; j < len(deckCards); j++ {
		r := len(deckCards) - j
		data := make(Cards, r)
		combinationUtil(cards, data, deckCards[:j], &newRank, 0, n-1, 0, r)
	}

	if newRank < bestRank {
		bestRank = newRank
	}

	return bestRank
}

// FindHandCategory returns
func FindHandCategory(cards Cards) int {
	sort.Sort(cards)

	countPerSuit := map[int]int{1: 0, 2: 0, 3: 0, 4: 0}
	gapCount := 0
	nCards := len(cards)
	countPerFaceValues := map[int]int{}

	for i := 0; i < nCards; i++ {
		countPerSuit[cards[i].Suit]++
		countPerFaceValues[cards[i].Position]++

		if i == nCards-1 {
			break
		}

		if cards[i].Position != cards[i+1].Position+1 {
			gapCount++
		}
	}

	suitsCount := numberOfNonZeroElements(&countPerSuit)
	if suitsCount == 1 {
		if gapCount == 0 {
			return HandCategoryRank("straight-flush")
		}

		return HandCategoryRank("flush")
	} else if len(countPerFaceValues) == 2 && hasElement(&countPerFaceValues, 4) {
		return HandCategoryRank("four-of-a-kind")
	} else if hasElement(&countPerFaceValues, 2) && hasElement(&countPerFaceValues, 3) {
		return HandCategoryRank("full-house")
	} else if gapCount == 0 {
		return HandCategoryRank("straight")
	} else if len(countPerFaceValues) == 3 && hasElement(&countPerFaceValues, 3) && !hasElement(&countPerFaceValues, 2) {
		return HandCategoryRank("three-of-a-kind")
	} else if len(countPerFaceValues) == 3 && hasElement(&countPerFaceValues, 2) && hasElement(&countPerFaceValues, 1) {
		return HandCategoryRank("two-pairs")
	} else if len(countPerFaceValues) == 4 && hasElement(&countPerFaceValues, 2) && hasElement(&countPerFaceValues, 1) {
		return HandCategoryRank("one-pair")
	} else {
		return HandCategoryRank("highest-card")
	}

}

func numberOfNonZeroElements(suitsWithCount *map[int]int) int {
	count := 0
	for _, v := range *suitsWithCount {
		if v > 0 {
			count++
		}
	}

	return count
}

func hasElement(suitsWithCount *map[int]int, elem int) bool {
	found := false
	for _, v := range *suitsWithCount {
		if v == elem {
			found = true
			break
		}
	}

	return found
}

// Card is abstruction of real world card from deck
type Card struct {
	Code     string
	Position int
	Suit     int
}

// NewCard returns ref to new instance of the Card
// by parsing the code
func NewCard(code string) *Card {
	parsableCode := []rune(code)
	position := faceValueToIndex[string(parsableCode[0])]
	suit := suitToIndex[string(parsableCode[1])]
	return &Card{code, position, suit}
}

// Cards is collection of ref to Card items
type Cards []*Card

func NewCards(rawInput string) *Cards {
	cards := Cards{}
	cardCodes := strings.Split(rawInput, " ")
	for i := range cardCodes {
		cards = append(cards, NewCard(cardCodes[i]))
	}

	return &cards
}

// Len is the number of elements in the collection.
func (c Cards) Len() int {
	return len(c)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (c Cards) Less(i, j int) bool {
	return c[i].Position > c[j].Position
}

// Swap swaps the elements with indexes i and j.
func (c Cards) Swap(i, j int) {
	t := c[j]
	c[j] = c[i]
	c[i] = t
}

var faceValueToIndex map[string]int
var suitToIndex map[string]int
var HandCategories map[string]int

// HandCategoryRank returns
func HandCategoryRank(name string) int {
	return HandCategories[name]
}

func HandCategoryName(rank int) string {
	for k, v := range HandCategories {
		if v == rank {
			return k
		}
	}

	return "not_found"
}

func init() {
	faceValueToIndex = map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
		"T": 10, "J": 11, "Q": 12, "K": 13, "A": 1}
	suitToIndex = map[string]int{"C": 1, "D": 2, "H": 3, "S": 4}

	HandCategories = make(map[string]int, 9)
	HandCategories["straight-flush"] = 1
	HandCategories["four-of-a-kind"] = 2
	HandCategories["full-house"] = 3
	HandCategories["flush"] = 4
	HandCategories["straight"] = 5
	HandCategories["three-of-a-kind"] = 6
	HandCategories["two-pairs"] = 7
	HandCategories["one-pair"] = 8
	HandCategories["highest-card"] = 9
}

func combinationUtil(inputArray Cards, data Cards, deckArray Cards, handRank *int, start, end, index, r int) {
	if index == r {
		// fmt.Printf("%v + %v\n", data[0:r], deckArray)
		candidateCards := Cards{}
		candidateCards = append(candidateCards, data[0:r]...)
		candidateCards = append(candidateCards, deckArray...)
		newRank := FindHandCategory(candidateCards)
		if newRank < *handRank {
			*handRank = newRank
			// fmt.Printf("newRank %d, handRank %d\n", newRank, handRank)
		}
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = inputArray[i]
		combinationUtil(inputArray, data, deckArray, handRank, i+1, end, index+1, r)
	}
}

func main() {
	deckCards := *NewCards("QH KH AH 2S 6S")
	arr := *NewCards("TH JH QC QD QS")

	fmt.Printf("bestRank = %s", HandCategoryName(Lookup(arr, deckCards)))
}
