package leetcode

import (
	"reflect"
	"testing"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		a [][]int
		b [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				a: [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
				b: [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			},
			want: [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			name: "test2",
			args: args{
				a: [][]int{{1, 3}, {5, 9}},
				b: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "test3",
			args: args{
				a: [][]int{},
				b: [][]int{{4, 8}, {10, 12}},
			},
			want: [][]int{},
		},
		{
			name: "test4",
			args: args{
				a: [][]int{{1, 7}},
				b: [][]int{{3, 10}},
			},
			want: [][]int{{3, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalIntersection(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intervalIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
