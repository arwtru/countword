package algorithms

// Subset test if given one element is subset of an array.
func Subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value]++
	}
	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}
