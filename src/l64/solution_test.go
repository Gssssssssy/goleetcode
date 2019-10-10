package l64

import (
	"testing"

	"gotest.tools/assert"
)

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string

		want  int
		input [][]int
	}{
		{
			want: 7,
			input: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		output := minPathSum(tt.input)
		assert.Equal(t, output, tt.want)
	}
}
