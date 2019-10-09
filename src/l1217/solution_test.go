package l1217

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_minCostToMoveChips(t *testing.T) {
	tests := []struct {
		name  string
		chips []int
		want  int
	}{
		{
			chips: []int{1, 2, 3},
			want:  1,
		},
		{
			chips: []int{2, 2, 2, 3, 3},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := minCostToMoveChips(tt.chips)
			assert.Equal(t, tt.want, output)
		})
	}
}
