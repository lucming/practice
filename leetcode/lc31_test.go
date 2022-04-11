package leetcode

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[]int{1, 2, 3}},
		},
		{
			name: "test2",
			args: args{[]int{3, 2, 1}},
		},
		{
			name: "test3",
			args: args{[]int{1, 1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
