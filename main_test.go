package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBestHand(t *testing.T) {
	hand := "TH JH QC QD QS"
	deck := "QH KH AH 2S 6S"
	assert.Equal(t, FindBestHand(hand, deck), "straight-flush")
}

func TestFindHandCategory(t *testing.T) {
	cards := []*Card{NewCard("JH"), NewCard("TH"), NewCard("QH"), NewCard("9H"), NewCard("8H")}
	assert.Equal(t, HandCategoryRank("straight-flush"), FindHandCategory(cards))

	cards = []*Card{NewCard("TH"), NewCard("TC"), NewCard("TD"), NewCard("TS"), NewCard("JS")}
	assert.Equal(t, HandCategoryRank("four-of-a-kind"), FindHandCategory(cards))

	cards = []*Card{NewCard("3S"), NewCard("3H"), NewCard("3D"), NewCard("6C"), NewCard("6H")}
	assert.Equal(t, HandCategoryRank("full-house"), FindHandCategory(cards))

	cards = []*Card{NewCard("JH"), NewCard("TH"), NewCard("QH"), NewCard("7H"), NewCard("2H")}
	assert.Equal(t, HandCategoryRank("flush"), FindHandCategory(cards))

	cards = []*Card{NewCard("7S"), NewCard("6S"), NewCard("5S"), NewCard("4H"), NewCard("3H")}
	assert.Equal(t, HandCategoryRank("straight"), FindHandCategory(cards))

	cards = []*Card{NewCard("2D"), NewCard("2S"), NewCard("2S"), NewCard("KS"), NewCard("6H")}
	assert.Equal(t, HandCategoryRank("three-of-a-kind"), FindHandCategory(cards))

	cards = []*Card{NewCard("JH"), NewCard("JS"), NewCard("4S"), NewCard("4H"), NewCard("9H")}
	assert.Equal(t, HandCategoryRank("two-pairs"), FindHandCategory(cards))

	cards = []*Card{NewCard("4H"), NewCard("4S"), NewCard("KS"), NewCard("TD"), NewCard("5S")}
	assert.Equal(t, HandCategoryRank("one-pair"), FindHandCategory(cards))

	cards = []*Card{NewCard("KH"), NewCard("JH"), NewCard("8S"), NewCard("7D"), NewCard("4S")}
	assert.Equal(t, HandCategoryRank("highest-card"), FindHandCategory(cards))
}

func TestSortByFaceValue(t *testing.T) {
	cards := Cards{
		NewCard("KC"), NewCard("2S"), NewCard("3H"), NewCard("4H"),
		NewCard("5H"), NewCard("6H"), NewCard("JS"), NewCard("7H"),
		NewCard("TH"), NewCard("9H"), NewCard("8H"), NewCard("AD"),
		NewCard("QH"),
	}
	sortedCards := Cards{
		NewCard("AD"), NewCard("KC"), NewCard("QH"), NewCard("JS"),
		NewCard("TH"), NewCard("9H"), NewCard("8H"), NewCard("7H"), NewCard("6H"),
		NewCard("5H"), NewCard("4H"), NewCard("3H"), NewCard("2S"),
	}
	sort.Sort(cards)
	assert.Equal(t, sortedCards, cards)
}

func TestLookupDeckSizeOne(t *testing.T) {
	cards := *NewCards("JH TH 3C 5S 2D")
	deckCards := *NewCards("AH")
	assert.Equal(t, HandCategoryRank("highest-card"), Lookup(cards, deckCards))

	cards = *NewCards("JH TH TC 5S 2D")
	deckCards = *NewCards("AH")
	assert.Equal(t, HandCategoryRank("one-pair"), Lookup(cards, deckCards))

	cards = *NewCards("JH TH TC JS 2D")
	deckCards = *NewCards("AH")
	assert.Equal(t, HandCategoryRank("two-pairs"), Lookup(cards, deckCards))

	cards = *NewCards("2H 2S TC JS 4D")
	deckCards = *NewCards("2D")
	assert.Equal(t, HandCategoryRank("three-of-a-kind"), Lookup(cards, deckCards))

	cards = *NewCards("3S KS 5C 6S 4D")
	deckCards = *NewCards("7D")
	assert.Equal(t, HandCategoryRank("straight"), Lookup(cards, deckCards))

	cards = *NewCards("TD 7C KC 4C 6C")
	deckCards = *NewCards("TC")
	assert.Equal(t, HandCategoryRank("flush"), Lookup(cards, deckCards))

	cards = *NewCards("3C 3S 3D 7C 6H")
	deckCards = *NewCards("6C")
	assert.Equal(t, HandCategoryRank("full-house"), Lookup(cards, deckCards))

	cards = *NewCards("3C 3S 3D 4C 6H")
	deckCards = *NewCards("3H")
	assert.Equal(t, HandCategoryRank("four-of-a-kind"), Lookup(cards, deckCards))

	cards = *NewCards("JH TH QH 9H 8H")
	deckCards = *NewCards("AH")
	assert.Equal(t, HandCategoryRank("straight-flush"), Lookup(cards, deckCards))
}

func TestLookupDeckSizeTwo(t *testing.T) {
	cards := *NewCards("JH TH 5S 2D")
	deckCards := *NewCards("AH 3C")
	assert.Equal(t, HandCategoryRank("highest-card"), Lookup(cards, deckCards))

	cards = *NewCards("JH TH 9H 8H")
	deckCards = *NewCards("AH QH")
	assert.Equal(t, HandCategoryRank("straight-flush"), Lookup(cards, deckCards))
}

func TestLookupDeckSizeThree(t *testing.T) {
	cards := *NewCards("JH TH 2D")
	deckCards := *NewCards("AH 3C 5S")
	assert.Equal(t, HandCategoryRank("highest-card"), Lookup(cards, deckCards))

	cards = *NewCards("AH 9H 8H")
	deckCards = *NewCards("JH QH TH")
	assert.Equal(t, HandCategoryRank("straight-flush"), Lookup(cards, deckCards))
}

func TestLookupDeckSizeFour(t *testing.T) {
	cards := *NewCards("JH TH")
	deckCards := *NewCards("AH 3C 5S 2D")
	assert.Equal(t, HandCategoryRank("highest-card"), Lookup(cards, deckCards))

	cards = *NewCards("AH 9H")
	deckCards = *NewCards("JH QH TH 8H")
	// assert.Equal(t, HandCategoryRank("straight-flush"), Lookup(cards, deckCards))
}
