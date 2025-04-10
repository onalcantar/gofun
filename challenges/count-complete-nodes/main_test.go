package main

import "testing"

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [1,2,3,4,5,6]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{Val: 6},
					},
				},
			},
			want: 6,
		},
		{
			name: "root = []",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "root = [1]",
			args: args{
				root: &TreeNode{Val: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
