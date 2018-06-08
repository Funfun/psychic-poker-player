package main

// Count returns number of positive ints
func Count(hash map[int]int) int {
	count := 0
	for _, v := range hash {
		if v > 0 {
			count++
		}
	}

	return count
}

// Included returns true if target
// found in hash table
func Included(vs map[int]int, t int) bool {
	found := false
	for _, v := range vs {
		if v == t {
			found = true
			break
		}
	}

	return found
}
