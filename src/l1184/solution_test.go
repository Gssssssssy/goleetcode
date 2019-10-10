package l1184

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_distanceBetweenBusStops(t *testing.T) {
	tests := []struct {
		name string

		want int
		args struct {
			distance    []int
			start       int
			destination int
		}
	}{
		{
			want: 1,
			args: struct {
				distance    []int
				start       int
				destination int
			}{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 1,
			},
		},
		{
			want: 3,
			args: struct {
				distance    []int
				start       int
				destination int
			}{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 2,
			},
		},
		{
			want: 4,
			args: struct {
				distance    []int
				start       int
				destination int
			}{
				distance:    []int{1, 2, 3, 4},
				start:       0,
				destination: 3,
			},
		},
		//extra test case1
		{
			want: 17,
			args: struct {
				distance    []int
				start       int
				destination int
			}{
				distance:    []int{7, 10, 1, 12, 11, 14, 5, 0},
				start:       7,
				destination: 2,
			},
		},
		//extra test case2: 超出时间限制
		{
			want: 1,
			args: struct {
				distance    []int
				start       int
				destination int
			}{
				distance:    []int{6, 47, 48, 1},
				start:       3,
				destination: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := distanceBetweenBusStops(tt.args.distance, tt.args.start, tt.args.destination)
			assert.Equal(t, output, tt.want)
		})
	}
}
