package mysort

import (
	"fmt"
	"testing"
)

func Test_bible(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "firt",
			args: args{
				[]int{1,8,43,2,12,35,67},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.arr,0,len(tt.args.arr)-1)
			fmt.Println(tt.args.arr)
		})
	}
}
