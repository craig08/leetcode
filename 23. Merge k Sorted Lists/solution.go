import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListPQ []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	head := new(ListNode)
	node := head
	var pq ListPQ
	for _, node := range lists {
		if node == nil {
			continue
		}
		pq = append(pq, node)
	}
	heap.Init(&pq)
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)
		if item.Next != nil {
			heap.Push(&pq, item.Next)
		}
		node.Next = item
		node = node.Next
	}
	return head.Next
}

func (pq ListPQ) Len() int           { return len(pq) }
func (pq ListPQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq ListPQ) Less(i, j int) bool { return pq[i].Val < pq[j].Val }

func (pq *ListPQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *ListPQ) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:(n - 1)]
	return x
}
