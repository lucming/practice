package leetcode

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
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
				board: [][]byte{
					{'a', 'b', 'c', 'd'},
					{'e', 'f', 'c', 'g'},
					{'x', 's', 'i', 'i'},
					{'a', 'u', 'k', 'r'},
				},
				word: "abfci",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{'a', 'b', 'c', 'd'},
					{'e', 'f', 'c', 'g'},
					{'x', 's', 'i', 'i'},
					{'a', 'u', 'k', 'r'},
				},
				word: "aaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
