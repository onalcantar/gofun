package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		},
		{
			name: "Sad path - foo, bar",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		},
		{
			name: "Happy path - paper, title",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
		{
			name: "Sad path - pizza, taco",
			args: args{
				s: "pizza",
				t: "taco",
			},
			want: false,
		},
		{
			name: "Sad path - peperonni, tataronni",
			args: args{
				s: "peperonni",
				t: "tataronni",
			},
			want: true,
		},
		{
			name: "Sad path - badc, baba",
			args: args{
				s: "badc",
				t: "baba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
