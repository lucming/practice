package lc

import (
	"testing"
)

func Test_node_headInsert(t *testing.T) {
	head := node{}
	for i := 0; i < 10; i++ {
		head.headInsert(i)
	}
	head.show()
	reverse2(&head)
	head.show()
}
