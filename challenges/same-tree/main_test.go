package main

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - p = [1,2,3], q = [1,2,3]",
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				q: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: true,
		},
		{
			name: "Sad path - p = [1,2], q = [1,null,2]",
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
				},
				q: &TreeNode{
					Val: 1,
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
		{
			name: "Sad path - p = [1,2,1], q = [1,1,2]",
			args: args{
				p: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{Val: 1},
				},
				q: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
