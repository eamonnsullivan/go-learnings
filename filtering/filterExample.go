package filtering

import "fmt"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func evenNumbers(min int, max int) []int {
	filter := func(number int) bool {
		return number%2 == 0
	}

	fullRange := makeRange(min, max+1)
	var filteredRange []int
	for a := range fullRange {
		if filter(a) {
			filteredRange = append(filteredRange, a)
		}
	}
	return filteredRange
}

func PrintEvens(min int, max int) {
	fmt.Println(evenNumbers(min, max))
}
