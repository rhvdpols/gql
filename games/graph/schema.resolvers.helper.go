package graph

func determineOrder(i int) []int {
	order := []int{0, 1, 2}

	if i > 10 {
		order = []int{1, 0, 2}
	}
	if i > 20 {
		order = []int{1, 2, 0}
	}
	if i > 30 {
		order = []int{2, 1, 0}
	}
	return order
}
