package leetcode

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{2, 4, 5, 6, 7, 8, 9, 19},
		{6, 8, 9, 10, 18, 20, 25, 50},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				matrix: arr,
				target: 8,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				matrix: arr,
				target: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxNumByDel3(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{13938},
			want: 1938,
		},
		{
			name: "test2",
			args: args{19873},
			want: 1987,
		},
		{
			name: "test3",
			args: args{-13938},
			want: -1398,
		},
		{
			name: "test4",
			args: args{-19383},
			want: -1938,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxNumByDel3(tt.args.num); got != tt.want {
				t.Errorf("MaxNumByDel3() = %v, want %v", got, tt.want)
			}
		})
	}
}
