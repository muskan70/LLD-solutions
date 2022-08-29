package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_AddNodeAtEnd(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "test1",
			fields: fields{head: nil, tail: nil},
			args:   args{value: "vipul"},
			want:   &Node{val: "vipul", next: nil, prev: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := d.AddNodeAtEnd(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoublyLinkedList.AddNodeAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveNode(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	type args struct {
		node *Node
	}
	node := &Node{val: "vipul", next: nil, prev: nil}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name:   "test2",
			fields: fields{head: node, tail: node},
			args:   args{node: node},
			want:   fields{head: nil, tail: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			d.RemoveNode(tt.args.node)
			if !reflect.DeepEqual(d.head, tt.want.head) {
				t.Errorf("DoublyLinkedList.RemoveNode() = %v, want %v", d.head, tt.want.head)
			}
		})
	}
}

func TestDoublyLinkedList_RemoveNodeAtHead(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
	}
	node := &Node{val: "vipul", next: nil, prev: nil}
	res := "vipul"
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name:   "test3",
			fields: fields{head: node, tail: node},
			want:   &res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := d.RemoveNodeAtHead(); got != tt.want {
				t.Errorf("DoublyLinkedList.RemoveNodeAtHead() = %v, want %v", got, tt.want)
			}
		})
	}
}
