package main

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sqrt - 16 | 4",
			args: args{
				x: 16,
			},
			want: 4,
		},
		{
			name: "sqrt - 4 | 2",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "sqrt - 8 | 2",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "sqrt - 1 | 1",
			args: args{
				x: 1,
			},
			want: 1,
		},
		{
			name: "sqrt - 0 | 0",
			args: args{
				x: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
