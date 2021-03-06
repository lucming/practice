package lc

import (
	"reflect"
	"testing"
)

func Test_node_headInsert(t *testing.T) {
	head := node{}
	for i := 0; i < 10; i++ {
		head.headInsert(i)
	}
	head.show()
	reverse(&head)
	head.show()
}

func getCycleList() *node {
	n8 := node{data: 8}
	n7 := node{data: 7, next: &n8}
	n6 := node{data: 6, next: &n7}
	n5 := node{data: 5, next: &n6}
	n4 := node{data: 4, next: &n5}
	n3 := node{data: 3, next: &n4}
	n2 := node{data: 2, next: &n3}
	n1 := node{data: 1, next: &n2}
	head := node{next: &n1}
	n8.next = &n5
	return &head
}

func Test_cycle(t *testing.T) {
	p1 := node{}
	for i := 0; i < 10; i++ {
		p1.headInsert(i)
	}

	p2 := node{}
	p2.next = &p2

	type args struct {
		head *node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "not cycle",
			args: args{head: &p1},
			want: false,
		},
		{
			name: "cycle",
			args: args{head: &p2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cycle(tt.args.head); got != tt.want {
				t.Errorf("cycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLoppNode(t *testing.T) {

	type args struct {
		head *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{
			name: "first node",
			args: args{head: getCycleList()},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLoppNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLoppNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deletePoint(t *testing.T) {
	n5 := node{data: 5}
	n4 := node{data: 4, next: &n5}
	n3 := node{data: 3, next: &n4}
	n2 := node{data: 2, next: &n3}
	n1 := node{data: 1, next: &n2}
	head := node{next: &n1}

	type args struct {
		head *node
		find *node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "delete exist node",
			args: args{
				head: &head,
				find: &n1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deletePoint(tt.args.head, tt.args.find)
			tt.args.head.show()
		})
	}
}

func Test_insertPri(t *testing.T) {
	n5 := node{data: 5}
	n4 := node{data: 4, next: &n5}
	n3 := node{data: 3, next: &n4}
	n2 := node{data: 2, next: &n3}
	n1 := node{data: 1, next: &n2}
	head := node{next: &n1}

	type args struct {
		head *node
		p    *node
		s    *node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "add",
			args: args{
				head: &head,
				p:    &n1,
				s:    &node{data: 100},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertPri(tt.args.head, tt.args.p, tt.args.s)
			tt.args.head.show()
		})
	}
}

func Test_deleteK(t *testing.T) {
	p1 := node{}
	for i := 0; i < 10; i++ {
		p1.headInsert(i)
	}

	type args struct {
		head *node
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				head: &p1,
				k:    3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteK(tt.args.head, tt.args.k)
			tt.args.head.show()
		})
	}
}

func TestSortList(t *testing.T) {
	//1???12???4???10???6???8???9
	n6 := node{data: 9}
	n5 := node{data: 8, next: &n6}
	n4 := node{data: 6, next: &n5}
	n3 := node{data: 10, next: &n4}
	n2 := node{data: 4, next: &n3}
	n1 := node{data: 12, next: &n2}
	head := node{next: &n1, data: 1}

	type args struct {
		head *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{&head},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SortList(tt.args.head)
			p.show()
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		head *node
	}
	arr := []int{1, 5, 4, 6, 3, 8, 7, 6, 2}
	var p *node
	for _, val := range arr {
		if p == nil {
			p = &node{val, nil}
		} else {
			cur := &node{val, p}
			p = cur
		}
	}
	p.show()

	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{p},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
				got.show()
			}
		})
	}
}
