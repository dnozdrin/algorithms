package algorithms

import "github.com/pkg/errors"

func Flatten(input [][]int) ([]int, error) {
	var index, x, y, min, max, length, capacity int

	if input == nil {
		return nil, errors.WithMessage(ErrInvalidInput, "got nil, a square matrix expected")
	}

	length = len(input)
	for _, row := range input {
		if len(row) != length {
			return nil, errors.WithMessage(ErrInvalidInput, "a square matrix (n*n) expected")
		}
	}

	capacity = length * length
	output := make([]int, capacity)

	max = length - 1
	last := capacity - 1

	for index <= last {
		for x == min && y < max {
			output[index] = input[x][y]
			y++
			index++
		}
		for y == max && x < max {
			output[index] = input[x][y]
			x++
			index++
		}
		for x == max && (y > min || index == last) {
			output[index] = input[x][y]
			y--
			index++
		}
		for y == min && x > min+1 {
			output[index] = input[x][y]
			x--
			index++
		}
		min++
		max--
	}

	return output, nil
}
