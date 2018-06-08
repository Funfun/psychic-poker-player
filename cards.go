package main

import "strings"

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
