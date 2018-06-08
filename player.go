package main

import "sort"

var faceValueToIndex map[string]int
var suitToIndex map[string]int

func init() {
	faceValueToIndex = map[string]int{
		"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
		"8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14,
	}
	suitToIndex = map[string]int{"C": 1, "D": 2, "H": 3, "S": 4}
}

// FindBestHand returns highest possible high-hand
// category from the given hand and deck sets
func FindBestHand(hand string, deck string) string {
	handCards := *NewCards(hand)
	deckCards := *NewCards(deck)

	return HandCategoryName(lookup(handCards, deckCards))
}

// lookup does look up of the best possible hand combination
func lookup(cards Cards, deckCards Cards) int {
	var newRank int
	bestRank := 9

	// assume initially the player has a good hand
	newRank = FindHandCategory(cards)
	if HandStraightFlush == bestRank {
		return HandStraightFlush
	}
	if newRank < bestRank {
		bestRank = newRank
	}

	// check all deck cards
	// copy to new slice to avoid shuffle
	tmp := make(Cards, len(deckCards))
	copy(tmp, deckCards)
	newRank = FindHandCategory(tmp)
	if HandStraightFlush == newRank {
		return HandStraightFlush
	}

	if newRank < bestRank {
		bestRank = newRank
	}

	n := len(cards)
	for j := 1; j < len(deckCards); j++ {
		r := len(deckCards) - j
		data := make(Cards, r)
		cardCombinations(cards, data, deckCards[:j], &bestRank, 0, n-1, 0, r)
	}

	if newRank < bestRank {
		bestRank = newRank
	}

	return bestRank
}

// cardCombinations implements look up of +r+ possible
// combimations of elements in the set with additional
// ordered elements from other set to be used to find
// the best poker hand rank among resulted combinations.
func cardCombinations(inputSet Cards, result Cards, orderedSet Cards, handRank *int, start, end, index, r int) {
	if index == r {
		candidateCards := Cards{}
		candidateCards = append(candidateCards, result[0:r]...)
		candidateCards = append(candidateCards, orderedSet...)
		newRank := FindHandCategory(candidateCards)

		if newRank < *handRank {
			*handRank = newRank
		}
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		result[index] = inputSet[i]
		cardCombinations(inputSet, result, orderedSet, handRank, i+1, end, index+1, r)
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
		countPerFaceValues[cards[i].FaceValue]++

		if i == nCards-1 {
			break
		}

		if cards[i].FaceValue != cards[i+1].FaceValue+1 {
			gapCount++
		}
	}

	suitsCount := Count(countPerSuit)

	if suitsCount == 1 {
		if gapCount == 0 {
			return HandStraightFlush
		}

		return HandFlush
	} else if len(countPerFaceValues) == 2 && Included(countPerFaceValues, 4) {
		return HandFourOfAKind
	} else if Included(countPerFaceValues, 2) && Included(countPerFaceValues, 3) {
		return HandFullHouse
	} else if gapCount == 0 {
		return HandStraight
	} else if len(countPerFaceValues) == 3 && Included(countPerFaceValues, 3) && !Included(countPerFaceValues, 2) {
		return HandThreeOfAKind
	} else if len(countPerFaceValues) == 3 && Included(countPerFaceValues, 2) && Included(countPerFaceValues, 1) {
		return HandTwoPairs
	} else if len(countPerFaceValues) == 4 && Included(countPerFaceValues, 2) && Included(countPerFaceValues, 1) {
		return HandOnePair
	} else {
		return HandHighestCard
	}

}
