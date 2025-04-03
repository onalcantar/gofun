package main

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - numer 19 is happy",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "Happy path - numer 1111111 is happy",
			args: args{
				n: 1111111,
			},
			want: true,
		},
		{
			name: "Sad path - numer 2 is unhappy",
			args: args{
				n: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
