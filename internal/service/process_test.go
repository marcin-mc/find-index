package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name   string
		data   []int
		target int
		index  int
		value  int
	}{
		{
			name:   "value exists",
			data:   []int{0, 10, 20, 30, 50},
			target: 30,
			index:  3,
			value:  30,
		},
		{
			name:   "value exists (first value 0)",
			data:   []int{0, 10, 20, 30, 50},
			target: 0,
			index:  0,
			value:  0,
		},
		{
			name:   "value exists (last value)",
			data:   []int{0, 10, 20, 30, 999999, 1000000},
			target: 1000000,
			index:  5,
			value:  1000000,
		},
		{
			name:   "value only in 10%% range",
			data:   []int{0, 10, 20, 30, 50},
			target: 33,
			index:  3,
			value:  30,
		},
		{
			name:   "value only in 10%% range",
			data:   []int{0, 100, 200, 270, 500},
			target: 300,
			index:  3,
			value:  270,
		},
		{
			name:   "value only in 10%% range higher than last number in array",
			data:   []int{0, 10, 20, 30, 50},
			target: 54,
			index:  4,
			value:  50,
		},
		{
			name:   "value not in 10%% range",
			data:   []int{0, 100, 200, 540, 100000},
			target: 601,
			index:  -1,
			value:  -1,
		},
		{
			name:   "value only not in 10%% range, higher than the last number in array",
			data:   []int{0, 100, 200, 500, 100000},
			target: 200000,
			index:  -1,
			value:  -1,
		},
		{
			name:   "select exact value if many values in 10%% range",
			data:   []int{0, 895, 890, 900, 905, 910, 915, 920, 925},
			target: 900,
			index:  3,
			value:  900,
		},
		{
			name:   "select proper value if many values in 10%% range",
			data:   []int{0, 890, 895, 900, 905, 910, 915, 920, 925},
			target: 892,
			index:  2,
			value:  895,
		},
		{
			name:   "one value only",
			data:   []int{0},
			target: 0,
			index:  0,
			value:  0,
		},
		{
			name:   "empty list",
			data:   []int{},
			target: 1,
			index:  -1,
			value:  -1,
		},
	}
	for _, tt := range tests {
		index, value := FindIndex(tt.data, tt.target)
		assert.Equal(t, tt.index, index)
		assert.Equal(t, tt.value, value)
	}

}
