package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sortTests = []struct {
	name  string
	items []int

	want []int
}{
	{"success/nil", nil, nil},
	{"success/one", []int{5}, []int{5}},
	{"success/reversed", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{"success/sorted", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{"success/duplicated", []int{1, 1, 2, 2, 3, 3, 9, 7, 8}, []int{1, 1, 2, 2, 3, 3, 7, 8, 9}},
	{
		"success/random",
		[]int{-100, 1452, 22435, 3456, 0, -1, 9, 678687, 0},
		[]int{-100, -1, 0, 0, 9, 1452, 3456, 22435, 678687},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range sortTests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BubbleSort(tt.items))
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for _, tt := range binarySearchTests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = BubbleSort(tt.items)
			}
		})
	}
}
