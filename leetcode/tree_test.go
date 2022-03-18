package leetcode

import (
	"reflect"
	"testing"
)

func Test_getCommonAncestor(t *testing.T) {
	n10 := node{value: 10}
	n9 := node{value: 9}
	n8 := node{value: 8}
	n7 := node{value: 7}
	n6 := node{value: 6}
	n5 := node{value: 5, left: &n10}
	n4 := node{value: 4, left: &n8, right: &n9}
	n3 := node{value: 3, left: &n6, right: &n7}
	n2 := node{value: 2, left: &n4, right: &n5}
	n1 := node{value: 1, left: &n2, right: &n3}

	type args struct {
		root *node
		p    *node
		q    *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				root: &n1,
				p:    &n2,
				q:    &n3,
			},
			want: &n1,
		},
		{
			name: "test2",
			args: args{
				root: &n1,
				p:    &n6,
				q:    &n7,
			},
			want: &n3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
