package main

func minHeapify(minHeap []*Node, index int) {
	smallest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < len(minHeap) && minHeap[left].frequency < minHeap[smallest].frequency {
		smallest = left
	}
	if right < len(minHeap) && minHeap[right].frequency < minHeap[smallest].frequency {
		smallest = right
	}
	if smallest != index {
		swapNodes(&minHeap[smallest], &minHeap[index])
		minHeapify(minHeap, smallest)
	}
}

func minHeapExtract(minHeap *[]*Node) *Node {
	if len(*minHeap) == 0 {
		return nil
	}
	temp := (*minHeap)[0]
	(*minHeap)[0] = (*minHeap)[len(*minHeap)-1]

	*minHeap = (*minHeap)[:len(*minHeap)-1]
	minHeapify(*minHeap, 0)

	return temp
}

func minHeapInsert(minHeap *[]*Node, minHeapNode *Node) {
	*minHeap = append(*minHeap, minHeapNode)
	i := len(*minHeap) - 1

	for i > 0 && ((*minHeapNode).frequency < (*minHeap)[(i-1)/2].frequency) {
		(*minHeap)[i] = (*minHeap)[(i-1)/2]
		i = (i - 1) / 2
	}
	(*minHeap)[i] = minHeapNode
}

func minHeapInit(data []rune, freq []int) []*Node {
	size := len(data)
	minHeap := make([]*Node, size)

	for i := 0; i < size; i++ {
		minHeap[i] = newNode(data[i], freq[i])
	}
	for i := (size - 2) / 2; i >= 0; i-- {
		minHeapify(minHeap, i)
	}
	return minHeap
}
