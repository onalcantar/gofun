package main

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy path - root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val: 4,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
		{
			name: "Sad path - root = [1,2,3], targetSum = 5",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "Sad path - root = [], targetSum = 0",
			args: args{
				root: nil,
				targetSum: 0,
			},
			want: false,
		},
		{
			name: "Sad path - root = [1,2], targetSum = 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{Val: 2},
				},
				targetSum: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
