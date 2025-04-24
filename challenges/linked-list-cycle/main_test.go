package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "head = [3,2,0,-4], pos = 1",
			args: args{
				head: func() *ListNode {
					n1 := &ListNode{Val: 3}
					n2 := &ListNode{Val: 2}
					n3 := &ListNode{Val: 0}
					n4 := &ListNode{Val: -4}
					n1.Next = n2
					n2.Next = n3
					n3.Next = n4
					n4.Next = n2
					return n1
				}(),
			},
			want: true,
		},
		{
			name: "head = [1,2], pos = 0",
			args: args{
				head: func() *ListNode {
					n1 := &ListNode{Val: 1}
					n2 := &ListNode{Val: 2}
					n1.Next = n2
					n2.Next = n1
					return n1
				}(),
			},
			want: true,
		},
		{
			name: "head = [1], pos = -1",
			args: args{
				head: func() *ListNode {
					n1 := &ListNode{Val: 1}
					return n1
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
