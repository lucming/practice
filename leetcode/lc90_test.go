package leetcode

import (
	"reflect"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[]int{1}},
			want: [][]int{{}, {1}},
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 2}},
			want: [][]int{{}, {2}, {2, 2}, {1}, {1, 2}, {1, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
