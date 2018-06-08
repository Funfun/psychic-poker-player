package main

import "sort"

var faceValueToIndex map[string]int
var suitToIndex map[string]int
var handCategories map[string]int

func init() {
	faceValueToIndex = map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
		"8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}
	suitToIndex = map[string]int{"C": 1, "D": 2, "H": 3, "S": 4}

	handCategories = make(map[string]int, 9)
	handCategories["straight-flush"] = 1
	handCategories["four-of-a-kind"] = 2
	handCategories["full-house"] = 3
	handCategories["flush"] = 4
	handCategories["straight"] = 5
	handCategories["three-of-a-kind"] = 6
	handCategories["two-pairs"] = 7
	handCategories["one-pair"] = 8
	handCategories["highest-card"] = 9
}

// FindBestHand returns highest possible high-hand
// category from the given hand and deck sets
func FindBestHand(hand string, deck string) string {
	handCards := *NewCards(hand)
	deckCards := *NewCards(deck)

	return HandCategoryName(Lookup(handCards, deckCards))
}

func Lookup(cards Cards, deckCards Cards) int {
	var newRank int
	bestRank := 9

	// assume initially the player has a good hand
	newRank = FindHandCategory(cards)
	if HandCategoryRank("straight-flush") == bestRank {
		return HandCategoryRank("straight-flush")
	}
	if newRank < bestRank {
		bestRank = newRank
	}

	// check all deck cards
	// copy to new slice to avoid shuffle
	tmp := make(Cards, len(deckCards))
	copy(tmp, deckCards)
	newRank = FindHandCategory(tmp)
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
		combinationUtil(cards, data, deckCards[:j], &bestRank, 0, n-1, 0, r)
	}

	if newRank < bestRank {
		bestRank = newRank
	}

	return bestRank
}

func combinationUtil(inputArray Cards, data Cards, deckArray Cards, handRank *int, start, end, index, r int) {
	if index == r {
		candidateCards := Cards{}
		candidateCards = append(candidateCards, data[0:r]...)
		candidateCards = append(candidateCards, deckArray...)
		newRank := FindHandCategory(candidateCards)

		if newRank < *handRank {
			*handRank = newRank
		}
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = inputArray[i]
		combinationUtil(inputArray, data, deckArray, handRank, i+1, end, index+1, r)
	}
}

// FindHandCategory returns best possible hand category
func FindHandCategory(cards Cards) int {
	sort.Sort(cards) // mutation!

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

	suitsCount := Count(&countPerSuit)

	if suitsCount == 1 {
		if gapCount == 0 {
			return HandCategoryRank("straight-flush")
		}

		return HandCategoryRank("flush")
	} else if len(countPerFaceValues) == 2 && Filter(&countPerFaceValues, 4) {
		return HandCategoryRank("four-of-a-kind")
	} else if Filter(&countPerFaceValues, 2) && Filter(&countPerFaceValues, 3) {
		return HandCategoryRank("full-house")
	} else if gapCount == 0 {
		return HandCategoryRank("straight")
	} else if len(countPerFaceValues) == 3 && Filter(&countPerFaceValues, 3) && !Filter(&countPerFaceValues, 2) {
		return HandCategoryRank("three-of-a-kind")
	} else if len(countPerFaceValues) == 3 && Filter(&countPerFaceValues, 2) && Filter(&countPerFaceValues, 1) {
		return HandCategoryRank("two-pairs")
	} else if len(countPerFaceValues) == 4 && Filter(&countPerFaceValues, 2) && Filter(&countPerFaceValues, 1) {
		return HandCategoryRank("one-pair")
	} else {
		return HandCategoryRank("highest-card")
	}

}

// HandCategoryRank returns
func HandCategoryRank(name string) int {
	return handCategories[name]
}

func HandCategoryName(rank int) string {
	for k, v := range handCategories {
		if v == rank {
			return k
		}
	}

	return "not_found"
}
