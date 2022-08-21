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

func (d *DoublyLinkedList) Add(key string) *Node {
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

func (d *DoublyLinkedList) Remove(node *Node) {
	if node == nil {
		return
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
}

func (d *DoublyLinkedList) Evict() *string {
	if d.head != nil {
		tmp := d.head
		d.head = d.head.next
		d.head.prev = nil
		return &tmp.val
	}
	return nil
}
