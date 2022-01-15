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
