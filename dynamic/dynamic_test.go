package dynamic

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[]int{10, 15, 20}},
			want: 15,
		},
		{
			name: "test2",
			args: args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs1(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "test2",
			args: args{
				m: 2,
				n: 3,
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "test4",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsHaveStone(t *testing.T) {
	type args struct {
		girds [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: 2,
		},
		{
			name: "test2",
			args: args{[][]int{{0, 1}, {0, 0}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsHaveStone(tt.args.girds); got != tt.want {
				t.Errorf("uniquePathsHaveStone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_integerBreak(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{2},
			want: 1,
		},
		{
			name: "test2",
			args: args{10},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerBreak(tt.args.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bagProblem01(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		bagweight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				weight:    []int{1, 3, 4},
				value:     []int{15, 20, 30},
				bagweight: 4,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagProblem01(tt.args.weight, tt.args.value, tt.args.bagweight); got != tt.want {
				t.Errorf("bagProblem01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bagProblem(t *testing.T) {
	type args struct {
		weight    []int
		value     []int
		bagweight int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				weight:    []int{1, 3, 4},
				value:     []int{15, 20, 30},
				bagweight: 4,
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagProblem(tt.args.weight, tt.args.value, tt.args.bagweight); got != tt.want {
				t.Errorf("bagProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[]int{1, 5, 11, 5}},
			want: true,
		},
		{
			name: "test2",
			args: args{[]int{1, 2, 3, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
