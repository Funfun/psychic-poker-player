package main

// Card is abstruction of real world card from deck
type Card struct {
	Code      string // two-character code combination of face-value & suit
	FaceValue int    // face-value (A=Ace, 2-9, T=10, J=Jack, Q=Queen, K=King)
	Suit      int    // suit (C=Clubs, D=Diamonds, H=Hearts, S=Spades)
}

// NewCard returns ref to new instance of the Card
// by parsing the code
func NewCard(code string) *Card {
	parsableCode := []rune(code)
	faceValue := faceValueToIndex[string(parsableCode[0])]
	suit := suitToIndex[string(parsableCode[1])]
	return &Card{code, faceValue, suit}
}
