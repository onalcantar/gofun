package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid pattern",
			args: args{
				pattern: "abba",
				s: "cat dog dog cat",
			},
			want: true,
		},
		{
			name: "Invalid pattern - invalid sequence",
			args: args{
				pattern: "abba",
				s: "cat dog cat dog",
			},
			want: false,
		},
		{
			name: "Invalid pattern - same pattern invalid words 1",
			args: args{
				pattern: "aaaa",
				s: "dog cat cat dog",
			},
			want: false,
		},
		{
			name: "Invalid pattern - same pattern invalid words 2",
			args: args{
				pattern: "abba",
				s: "dog dog dog dog",
			},
			want: false,
		},
		{
			name: "Invalid pattern - same pattern invalid words 3",
			args: args{
				pattern: "aaa",
				s: "aa aa aa aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("test: %s, wordPattern() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
