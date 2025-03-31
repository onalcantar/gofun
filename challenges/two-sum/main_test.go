package main

import (
	"testing"
	"slices"
)

func equalUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Happy path - [2,7,11,15] - 9",
			args: args{
				nums: []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Happy path - [2, 1, 15, 11, 7, 3] - 9",
			args: args{
				nums: []int{2, 1, 15, 11, 7, 3},
				target: 9,
			},
			want: []int{0, 4},
		},
		{
			name: "Happy path - [3, 2, 4] - 6",
			args: args{
				nums: []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Happy path - [3, 3] - 6",
			args: args{
				nums: []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "Happy path - [0, 4, 3, 0] - 0",
			args: args{
				nums: []int{0, 4, 3, 0},
				target: 0,
			},
			want: []int{0, 3},
		},
		{
			name: "Happy path - [-3, 4, 3, 90] - 0",
			args: args{
				nums: []int{-3, 4, 3, 90},
				target: 0,
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !equalUnordered(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
