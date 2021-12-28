package stack

import (
	"fmt"
	"testing"
)

func Test_queue_Push(t *testing.T) {
	type fields struct {
		stack1 []int
		stack2 []int
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
			name:   "normal",
			fields: fields{stack1: []int{}, stack2: []int{}},
			args:   args{val: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			for i := 0; i < 5; i++ {
				q.Push(i)
			}
			q.Pop()
			fmt.Println("peek:", q.Peek())
			for i := 0; i < 5; i++ {
				fmt.Println(q.Pop())
			}
		})
	}
}
