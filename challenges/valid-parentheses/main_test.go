package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - ()",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "Happy path - ()[]{}",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "Sad path - (]",
			args: args{
				s: "(]",
			},
			want: false,
		},
		{
			name: "Happy path - ([])",
			args: args{
				s: "([])",
			},
			want: true,
		},
		{
			name: "Sad path - ([)]",
			args: args{
				s: "([)]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
