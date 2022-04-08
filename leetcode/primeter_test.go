package leetcode

import "testing"

func Test_perimeter(t *testing.T) {
	type args struct {
		arr [][]int
		p   point
	}
	arr1 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0}}
	arr2 := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0}}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{ //所求的点在黑色区域
			name: "point in colored gird",
			args: args{
				arr: arr1,
				p: point{
					x: 1,
					y: 2,
				},
			},
			want: 12,
		},
		{
			name: "point in colored gird",
			args: args{
				arr: arr2,
				p: point{
					x: 2,
					y: 2,
				},
			},
			want: 16,
		},
		{
			name: "point not in colored gird",
			args: args{
				arr: arr2,
				p: point{
					x: 0,
					y: 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := perimeter(tt.args.arr, tt.args.p); got != tt.want {
				t.Errorf("perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
