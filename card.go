package main

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
