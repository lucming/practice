package leetcode

import "testing"

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"tree"},
			want: "eert",
		},
		{
			name: "test2",
			args: args{"cccaaa"},
			want: "aaaccc",
		},
		{
			name: "test3",
			args: args{"Aabb"},
			want: "bbAa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
