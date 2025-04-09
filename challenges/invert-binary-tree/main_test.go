package main

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root = [4,2,7,1,3,6,9]",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{Val: 6},
						Right: &TreeNode{Val: 9},
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "root = [2,1,3]",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name: "root = []",
			args: args{
				root: &TreeNode{},
			},
			want: &TreeNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
