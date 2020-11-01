package algorithms

func BubbleSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	for sweepNum := 0; sweepNum < len(items); sweepNum++ {
		swapped := false
		for i := 0; i < len(items)-1-sweepNum; i++ {
			if items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return items
}

func SelectionSort(items []int) []int {
	panic("not implemented")
}

func QuickSort(items []int) []int {
	panic("not implemented")
}

func MergeSort(items []int) []int {
	panic("not implemented")
}
