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
