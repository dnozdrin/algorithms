package algorithms

// BinarySearchRecursive accepts a sorted input
func BinarySearchRecursive(items []int, target int) (int, error) {
	length := len(items)
	if length == 0 {
		return 0, ErrEmpty
	}

	return binarySearchRecursive(items, target, 0, length)
}

func binarySearchRecursive(items []int, target, low, high int) (int, error) {
	if high == low {
		return 0, ErrNotFound
	}

	middle := (low + high - 1) >> 1
	if items[middle] == target {
		return middle, nil
	}

	if items[middle] < target {
		return binarySearchRecursive(items, target, middle+1, high)
	}

	return binarySearchRecursive(items, target, low, middle)
}

// BinarySearch accepts a sorted input
func BinarySearch(items []int, target int) (int, error) {
	length := len(items)
	if length == 0 {
		return 0, ErrEmpty
	}

	low := 0
	high := length - 1

	for low <= high {
		middle := (low + high) >> 1
		val := items[middle]

		if val < target {
			low = middle + 1
		} else if val > target {
			high = middle - 1
		} else {
			return middle, nil
		}
	}

	result := -(low + 1)
	if result < 0 {
		return 0, ErrNotFound
	}

	return result, nil
}
