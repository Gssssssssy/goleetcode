package l1185

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_dayOfTheWeek(t *testing.T) {
	tests := []struct {
		name string

		want string
		args struct {
			day   int
			month int
			year  int
		}
	}{
		{
			want: "Saturday",
			args: struct {
				day   int
				month int
				year  int
			}{
				day:   31,
				month: 8,
				year:  2019,
			},
		},
		{
			want: "Sunday",
			args: struct {
				day   int
				month int
				year  int
			}{
				day:   18,
				month: 7,
				year:  1999,
			},
		},
		{
			want: "Sunday",
			args: struct {
				day   int
				month int
				year  int
			}{
				day:   15,
				month: 8,
				year:  1993,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := dayOfTheWeek(tt.args.day, tt.args.month, tt.args.year)
			assert.Equal(t, tt.want, output)
		})
	}
}
