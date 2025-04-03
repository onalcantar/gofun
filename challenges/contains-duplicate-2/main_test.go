package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - nums = [1,2,3,1], k = 3",
			args: args{
				nums: []int{1, 2, 3, 1},
				k: 3,
			},
			want: true,
		},
		{
			name: "Happy path - [1,0,1,1], k = 1",
			args: args{
				nums: []int{1, 0, 1, 1},
				k: 1,
			},
			want: true,
		},
		{
			name: "Sad path - [1,2,3,1,2,3], k = 2",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
