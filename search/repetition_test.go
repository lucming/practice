package search

import (
	"testing"
)

func Test_duplicate(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "in",
			args: args{arr: []int{2, 3, 1, 0, 2, 5, 3}},
			want: 2,
		},
		{
			name: "not in",
			args: args{arr: []int{2, 3, 1, 0, 4, 6, 5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicate(tt.args.arr); got != tt.want {
				t.Errorf("duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBuplicate(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "in",
			args: args{arr: []int{2, 2, 2, 1, 2, 5, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBuplicate(tt.args.arr); got != tt.want {
				t.Errorf("getBuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
