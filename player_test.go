package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBestHand(t *testing.T) {
	hand := "TH JH QC QD QS"
	deck := "QH KH AH 2S 6S"
	assert.Equal(t, FindBestHand(hand, deck), "straight-flush")

	hand = "2H 2S 3H 3S 3C"
	deck = "2D 3D 6C 9C TH"
	assert.Equal(t, FindBestHand(hand, deck), "four-of-a-kind")

	hand = "2H 2S 3H 3S 3C"
	deck = "2D 9C 3D 6C TH"
	assert.Equal(t, FindBestHand(hand, deck), "full-house")

	hand = "2H AD 5H AC 7H"
	deck = "AH 6H 9H 4H 3C"
	assert.Equal(t, FindBestHand(hand, deck), "flush")

	hand = "AC 2D 9C 3S KD"
	deck = "5S 4D KS AS 4C"
	assert.Equal(t, FindBestHand(hand, deck), "straight")

	hand = "KS AH 2H 3C 4H"
	deck = "KC 2C TC 2D AS"
	assert.Equal(t, FindBestHand(hand, deck), "three-of-a-kind")

	hand = "AH 2C 9S AD 3CC"
	deck = "QH KS JS JD KD"
	assert.Equal(t, FindBestHand(hand, deck), "two-pairs")

	hand = "6C 9C 8C 2D 7C"
	deck = "2H TC 4C 9S AH"
	assert.Equal(t, FindBestHand(hand, deck), "one-pair")

	hand = "3D 5S 2H QD TD"
	deck = "6S KH 9H AD QH"
	assert.Equal(t, FindBestHand(hand, deck), "highest-card")
}

func TestFindHandCategory(t *testing.T) {
	cards := []*Card{NewCard("JH"), NewCard("TH"), NewCard("QH"), NewCard("9H"), NewCard("8H")}
	assert.Equal(t, HandStraightFlush, FindHandCategory(cards))

	cards = []*Card{NewCard("TH"), NewCard("TC"), NewCard("TD"), NewCard("TS"), NewCard("JS")}
	assert.Equal(t, HandFourOfAKind, FindHandCategory(cards))

	cards = []*Card{NewCard("3S"), NewCard("3H"), NewCard("3D"), NewCard("6C"), NewCard("6H")}
	assert.Equal(t, HandFullHouse, FindHandCategory(cards))

	cards = []*Card{NewCard("JH"), NewCard("TH"), NewCard("QH"), NewCard("7H"), NewCard("2H")}
	assert.Equal(t, HandFlush, FindHandCategory(cards))

	cards = []*Card{NewCard("7S"), NewCard("6S"), NewCard("5S"), NewCard("4H"), NewCard("3H")}
	assert.Equal(t, HandStraight, FindHandCategory(cards))

	cards = []*Card{NewCard("2D"), NewCard("2S"), NewCard("2S"), NewCard("KS"), NewCard("6H")}
	assert.Equal(t, HandThreeOfAKind, FindHandCategory(cards))

	cards = []*Card{NewCard("JH"), NewCard("JS"), NewCard("4S"), NewCard("4H"), NewCard("9H")}
	assert.Equal(t, HandTwoPairs, FindHandCategory(cards))

	cards = []*Card{NewCard("4H"), NewCard("4S"), NewCard("KS"), NewCard("TD"), NewCard("5S")}
	assert.Equal(t, HandOnePair, FindHandCategory(cards))

	cards = []*Card{NewCard("KH"), NewCard("JH"), NewCard("8S"), NewCard("7D"), NewCard("4S")}
	assert.Equal(t, HandHighestCard, FindHandCategory(cards))
}
