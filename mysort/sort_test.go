package mysort

import (
	"fmt"
	"testing"
)

func Test_bible(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "firt",
			args: args{
				[]int{1, 8, 43, 2, 12, 35, 67},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.arr, 0, len(tt.args.arr)-1)
			fmt.Println(tt.args.arr)
		})
	}
}

func Test_heapSort(t *testing.T) {
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
			args: args{[]int{1, 9, 6, 3, 7, 5}},
		},
		{
			name: "test2",
			args: args{[]int{1, 5, 4, 3, 5, 10, 6, 7, 2}},
		},
		{
			name: "test3",
			args: args{[]int{3, 21, 3, 45, 6, 7, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
		heapSort(tt.args.nums)
		fmt.Println(tt.args.nums)
	}
}
