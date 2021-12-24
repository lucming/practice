package search

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 5, 2, 3, 2, 6, 54},
				val:  2,
			},
			want: 8,
		},
		{
			name: "not exitst",
			args: args{
				nums: []int{1, 2, 3, 3, 4, 5, 2, 3, 2, 6, 54},
				val:  100,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.nums)
		})
	}
}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{[]int{-3, -2, 0, 1, 5}},
			want: []int{0, 1, 4, 9, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		nums []int
		s    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 5, 2, 1, 4, 2, 1},
				s:    10,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.nums, tt.args.s); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
