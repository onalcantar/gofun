package main

import "testing"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root = [1,2,2,3,4,4,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{Val: 4},
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: true,
		},
		{
			name: "root = [1,2,2,null,3,null,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{Val: 3},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
