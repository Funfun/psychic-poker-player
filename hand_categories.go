package main

const (
	HandStraightFlush = 1
	HandFourOfAKind   = 2
	HandFullHouse     = 3
	HandFlush         = 4
	HandStraight      = 5
	HandThreeOfAKind  = 6
	HandTwoPairs      = 7
	HandOnePair       = 8
	HandHighestCard   = 9
)

var handCategories map[string]int

func init() {
	handCategories = make(map[string]int, 9)
	handCategories["straight-flush"] = HandStraightFlush
	handCategories["four-of-a-kind"] = HandFourOfAKind
	handCategories["full-house"] = HandFullHouse
	handCategories["flush"] = HandFlush
	handCategories["straight"] = HandStraight
	handCategories["three-of-a-kind"] = HandThreeOfAKind
	handCategories["two-pairs"] = HandTwoPairs
	handCategories["one-pair"] = HandOnePair
	handCategories["highest-card"] = HandHighestCard
}

// HandCategoryName does reverse look up of value of handCategories
func HandCategoryName(rank int) string {
	for k, v := range handCategories {
		if v == rank {
			return k
		}
	}

	return "not_found"
}
