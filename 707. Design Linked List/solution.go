type MyLinkedList struct {
	Dummy *MyLinkedNode
	Len   int
}

type MyLinkedNode struct {
	Val        int
	Prev, Next *MyLinkedNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	dummy := new(MyLinkedNode)
	dummy.Prev, dummy.Next = dummy, dummy
	return MyLinkedList{Dummy: dummy}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Len {
		return -1
	}
	return this.MoveToIdx(index).Val
}

func (this *MyLinkedList) MoveToIdx(index int) *MyLinkedNode {
	node := this.Dummy
	if index < this.Len/2 {
		for i := 0; i <= index; i++ {
			node = node.Next
		}
	} else {
		for i := 0; i < this.Len-index; i++ {
			node = node.Prev
		}
	}
	return node
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	InsertAfter(this.Dummy, &MyLinkedNode{Val: val})
	this.Len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	InsertAfter(this.Dummy.Prev, &MyLinkedNode{Val: val})
	this.Len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Len {
		return
	} else if index < 0 {
		this.AddAtHead(val)
		return
	}
	node := this.MoveToIdx(index - 1)
	InsertAfter(node, &MyLinkedNode{Val: val})
	this.Len++
}

func InsertAfter(node, toBeInserted *MyLinkedNode) {
	toBeInserted.Prev, toBeInserted.Next = node, node.Next
	node.Next.Prev = toBeInserted
	node.Next = toBeInserted
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Len {
		return
	}
	node := this.MoveToIdx(index)
	node.Next.Prev, node.Prev.Next = node.Prev, node.Next
	this.Len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
