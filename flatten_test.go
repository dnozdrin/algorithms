package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	name  string
	input [][]int
	want  []int
	err   error
}{
	{

		"success/n = 0",
		[][]int{},
		[]int{},
		nil,
	},
	{
		"success/n = 1",
		[][]int{{3}},
		[]int{3},
		nil,
	},
	{
		"success/n = 2",
		[][]int{{9, 2}, {6, 4}},
		[]int{9, 2, 4, 6},
		nil,
	},
	{
		"success/n = 3",
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		nil,
	},
	{
		"success/n = 4",
		[][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}},
		[]int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
		nil,
	},
	{
		"success/n = 5",
		[][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		[]int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13},
		nil,
	},
	{
		"success/n = 20",
		[][]int{
			{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
			{20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39},
			{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
			{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79},
			{80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
			{100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119},
			{120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139},
			{140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159},
			{160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179},
			{180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
			{200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219},
			{220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239},
			{240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259},
			{260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279},
			{280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 298, 299},
			{300, 301, 302, 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, 315, 316, 317, 318, 319},
			{320, 321, 322, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 335, 336, 337, 338, 339},
			{340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359},
			{360, 361, 362, 363, 364, 365, 366, 367, 368, 369, 370, 371, 372, 373, 374, 375, 376, 377, 378, 379},
			{380, 381, 382, 383, 384, 385, 386, 387, 388, 389, 390, 391, 392, 393, 394, 395, 396, 397, 398, 399},
		},
		[]int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 39, 59, 79, 99, 119, 139, 159,
			179, 199, 219, 239, 259, 279, 299, 319, 339, 359, 379, 399, 398, 397, 396, 395, 394, 393, 392, 391, 390,
			389, 388, 387, 386, 385, 384, 383, 382, 381, 380, 360, 340, 320, 300, 280, 260, 240, 220, 200, 180, 160,
			140, 120, 100, 80, 60, 40, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 58,
			78, 98, 118, 138, 158, 178, 198, 218, 238, 258, 278, 298, 318, 338, 358, 378, 377, 376, 375, 374, 373, 372,
			371, 370, 369, 368, 367, 366, 365, 364, 363, 362, 361, 341, 321, 301, 281, 261, 241, 221, 201, 181, 161,
			141, 121, 101, 81, 61, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 77, 97, 117,
			137, 157, 177, 197, 217, 237, 257, 277, 297, 317, 337, 357, 356, 355, 354, 353, 352, 351, 350, 349, 348,
			347, 346, 345, 344, 343, 342, 322, 302, 282, 262, 242, 222, 202, 182, 162, 142, 122, 102, 82, 62, 63, 64,
			65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 96, 116, 136, 156, 176, 196, 216, 236, 256, 276, 296, 316,
			336, 335, 334, 333, 332, 331, 330, 329, 328, 327, 326, 325, 324, 323, 303, 283, 263, 243, 223, 203, 183,
			163, 143, 123, 103, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 115, 135, 155, 175, 195, 215, 235,
			255, 275, 295, 315, 314, 313, 312, 311, 310, 309, 308, 307, 306, 305, 304, 284, 264, 244, 224, 204, 184,
			164, 144, 124, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 134, 154, 174, 194, 214, 234, 254,
			274, 294, 293, 292, 291, 290, 289, 288, 287, 286, 285, 265, 245, 225, 205, 185, 165, 145, 125, 126, 127,
			128, 129, 130, 131, 132, 133, 153, 173, 193, 213, 233, 253, 273, 272, 271, 270, 269, 268, 267, 266, 246,
			226, 206, 186, 166, 146, 147, 148, 149, 150, 151, 152, 172, 192, 212, 232, 252, 251, 250, 249, 248, 247,
			227, 207, 187, 167, 168, 169, 170, 171, 191, 211, 231, 230, 229, 228, 208, 188, 189, 190, 210, 209,
		},
		nil,
	},
	{
		"error/not_square_matrix",
		[][]int{{9}, {6, 4}},
		nil,
		ErrInvalidInput,
	},
	{
		"error/nil",
		nil,
		nil,
		ErrInvalidInput,
	},
}

func TestFlatten(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Flatten(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.err != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	for _, tt := range tests {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = Flatten(tt.input)
			}
		})
	}
}
