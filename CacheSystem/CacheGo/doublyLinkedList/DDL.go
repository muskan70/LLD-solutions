package doublylinkedlist

type Node struct {
	val  string
	next *Node
	prev *Node
}

func newNode(x string) *Node {
	return &Node{val: x, next: nil, prev: nil}
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (d *DoublyLinkedList) AddNodeAtEnd(key string) *Node {
	n := newNode(key)
	if d.head == nil {
		d.head = n
		d.tail = d.head
	} else {
		d.tail.next = n
		n.prev = d.tail
		d.tail = n
	}
	return n
}

func (d *DoublyLinkedList) RemoveNode(node *Node) {
	if node == d.head && node == d.tail {
		d.head = nil
		d.tail = nil
	} else if node == d.head {
		d.head = node.next
		d.head.prev = nil
	} else if node == d.tail {
		d.tail = node.prev
		d.tail.next = nil
	} else if node.next != nil && node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}

func (d *DoublyLinkedList) RemoveNodeAtHead() *string {
	if d.head != nil {
		tmp := d.head
		d.head = d.head.next
		d.head.prev = nil
		return &tmp.val
	}
	return nil
}
