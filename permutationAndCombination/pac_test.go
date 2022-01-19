package permutationAndCombination

import (
	"reflect"
	"testing"
)

func Test_getCombinations(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				n: 3,
				k: 2,
			},
			want: [][]int{
				{1, 2},
				{1, 3},
				{2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCombinations(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationsSum(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				n: 7,
				k: 3,
			},
			want: [][]int{{1, 2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationsSum(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationsSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationSum(t *testing.T) {
	type args struct {
		condidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				condidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "normal1",
			args: args{
				condidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum(tt.args.condidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationSum1(t *testing.T) {
	type args struct {
		condidates []int
		target     int
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
				condidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			want: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name: "test2",
			args: args{
				condidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			want: [][]int{{1, 2, 2}, {5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum1(tt.args.condidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombinationSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{"aab"},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestoreIp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"25525511135"},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "test2",
			args: args{"0000"},
			want: []string{"0.0.0.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RestoreIp(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestoreIp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
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
			name: "normal",
			args: args{[]int{1, 2, 3}},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubsequences(t *testing.T) {
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
			name: "normal",
			args: args{[]int{4, 6, 7, 7}},
			want: [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubsequences(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
