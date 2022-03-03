package leetcode

import "testing"

func Test_isCycle(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{121},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCycle(tt.args.data); got != tt.want {
				t.Errorf("isCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
