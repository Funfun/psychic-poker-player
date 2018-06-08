package main

func Count(hash *map[int]int) int {
	count := 0
	for _, v := range *hash {
		if v > 0 {
			count++
		}
	}

	return count
}

func Filter(hash *map[int]int, elem int) bool {
	found := false
	for _, v := range *hash {
		if v == elem {
			found = true
			break
		}
	}

	return found
}
