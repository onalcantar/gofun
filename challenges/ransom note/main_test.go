package main

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - aa | aab",
			args: args{
				ransomNote: "aa",
				magazine: "aab",
			},
			want: true,
		},
		{
			name: "Sad path - a | b",
			args: args{
				ransomNote: "a",
				magazine: "b",
			},
			want: false,
		},
		{
			name: "Sad path - aa | ab",
			args: args{
				ransomNote: "aa",
				magazine: "ab",
			},
			want: false,
		},
		{
			name: "Sad path - aaa | aa",
			args: args{
				ransomNote: "aaa",
				magazine: "aa",
			},
			want: false,
		},
		{
			name: "Happy path - amor | roma",
			args: args{
				ransomNote: "amor",
				magazine: "roma",
			},
			want: true,
		},
		{
			name: "Happy path - amor | mora",
			args: args{
				ransomNote: "amor",
				magazine: "mora",
			},
			want: true,
		},
		{
			name: "Happy path - simio | miopesi",
			args: args{
				ransomNote: "simio",
				magazine: "miopesi",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
