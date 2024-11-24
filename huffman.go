package main

import "fmt"

func generateHuffmanTree(data []rune, freq []int) *Node {
	minHeap := minHeapInit(data, freq)

	for len(minHeap) != 1 {
		smallest1 := minHeapExtract(&minHeap)
		smallest2 := minHeapExtract(&minHeap)

		top := newNode('/', smallest1.frequency+smallest2.frequency)
		top.left = smallest1
		top.right = smallest2

		minHeapInsert(&minHeap, top)
	}
	return minHeapExtract(&minHeap)
}

func generateHuffmanCodes(root *Node) map[rune]string {
	codes := make(map[rune]string)
	huffmanEncoder(root, "", codes)
	return codes
}

func huffmanEncoder(root *Node, code string, codes map[rune]string) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		codes[root.data] = code
	}
	huffmanEncoder(root.left, code+"0", codes)
	huffmanEncoder(root.right, code+"1", codes)
}

func PrintHuffmanCodes(codes map[rune]string) {
	for char, code := range codes {
		fmt.Printf("%c : %s\n", char, code)
	}
}
