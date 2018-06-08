package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
