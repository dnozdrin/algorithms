package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var binarySearchTests = []struct {
	name   string
	items  []int
	target int

	want int
	err  error
}{
	{"error/nil", nil, 5, 0, ErrEmpty},
	{"error/empty_slice", []int{}, 5, 0, ErrEmpty},
	{"error/not_found/less", []int{1}, 5, 0, ErrNotFound},
	{"error/not_found/greater", []int{7}, 5, 0, ErrNotFound},
	{"error/not_found", []int{1, 6}, 5, 0, ErrNotFound},
	{"success/one_item", []int{5}, 5, 0, nil},
	{"success/even_length", []int{1, 2, 3, 4, 5, 6, 7, 8}, 5, 4, nil},
	{"success/odd_length", []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 5, 5, nil},
	{"success/random", []int{55, 123, 231, 232, 967, 34534}, 232, 3, nil},
	{"success/duplicates", []int{1, 2, 3, 4, 4, 4, 5, 6}, 2, 1, nil},
	{"success/many", []int{
		0, 2, 4, 6, 8, 10, 11, 12, 13, 14, 15, 21, 23, 25, 28, 31, 34, 38, 60, 68, 70,
		100, 102, 104, 106, 108, 110, 111, 112, 113, 114, 115, 121, 123, 125, 128, 131, 134, 238, 360, 468, 570,
		901, 1056, 1057, 1058, 1059, 1070, 1090, 1095, 2131, 2323, 2425, 2626, 2666, 4555, 10000, 99999, 999999,
	}, 2131, 50, nil},
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range binarySearchTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearch(tt.items, tt.target)
			assert.Equal(t, tt.want, got)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range binarySearchTests {
			tt := tt
			_, _ = BinarySearch(tt.items, tt.target)
		}
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	for _, tt := range binarySearchTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinarySearchRecursive(tt.items, tt.target)
			assert.Equal(t, tt.want, got)
			if tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func BenchmarkBinarySearchRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range binarySearchTests {
			tt := tt
			_, _ = BinarySearchRecursive(tt.items, tt.target)
		}
	}
}
