package main

var handCategories map[string]int

func init() {
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

// HandCategoryRank acts as getter for handCategories
func HandCategoryRank(name string) int {
	return handCategories[name]
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
