package stack

import (
	"fmt"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	type fields struct {
		queue1 []int
		queue2 []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			fields: fields{
				queue1: []int{},
				queue2: []int{},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				queue1: tt.fields.queue1,
				queue2: tt.fields.queue2,
			}
			for i := 0; i < 5; i++ {
				s.Push(i)
			}
			for i := 0; i < 5; i++ {
				fmt.Println(s.Pop())
			}
		})
	}
}
