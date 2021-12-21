package search

import "testing"

func Test_binary(t *testing.T) {
	type args struct {
		arr  []int
		data int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "first",
			args: args{
				arr:  []int{1, 3, 5, 6, 7, 8, 9, 10},
				data: 6,
			},
			want: 3,
		},
		{
			name: "no data",
			args: args{
				arr:  []int{1, 3, 5, 6, 7, 8, 9, 10},
				data: 12123,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binary(tt.args.arr, tt.args.data); got != tt.want {
				t.Errorf("binary() = %v, want %v", got, tt.want)
			}
		})
	}
}
