package dataStructures

type Node struct {
	val  string
	next *Node
	prev *Node
}

func newNode(x string) *Node {
	return &Node{val: x, next: nil, prev: nil}
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
	}
}

func (d *DoubleLinkedList) AddNodeAtEnd(x string) *Node {
	n := newNode(x)
	if d.head == nil {
		d.head = n
		d.tail = n
	} else {
		d.tail.next = n
		n.prev = d.tail
		d.tail = d.tail.next
	}
	return n
}

func (d *DoubleLinkedList) RemoveNodeAtHead() string {
	if d.head == d.tail {
		n := d.head
		d.head = nil
		d.tail = nil
		return n.val
	} else if d.head != nil {
		n := d.head
		d.head = d.head.next
		n.next = nil
		d.head.prev = nil
		return n.val
	}
	return ""
}

func (d *DoubleLinkedList) RemoveNode(n *Node) {
	if d.head == n && d.tail == n {
		d.head = nil
		d.tail = nil
	} else if d.head == n {
		d.head = d.head.next
		d.head.prev = nil
	} else if d.tail == n {
		d.tail = d.tail.prev
		d.tail.next = nil
	} else {
		n.next.prev = n.prev
		n.prev.next = n.next
	}
}

func (d *DoubleLinkedList) RemoveKey(key string) {
	if d.head.val == key && d.tail.val == key {
		d.head = nil
		d.tail = nil
	} else if d.head.val == key {
		d.head = d.head.next
		d.head.prev = nil
	} else if d.tail.val == key {
		d.tail = d.tail.prev
		d.tail.next = nil
	} else {
		for temp := d.head; temp != nil; temp = temp.next {
			if temp.val == key {
				temp.next.prev = temp.prev
				temp.prev.next = temp.next
			}
		}
	}
}

func (d *DoubleLinkedList) IsEmpty() bool {
	return d.head == d.tail && d.head == nil
}
