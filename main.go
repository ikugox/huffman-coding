package main

func main() {
	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	freq := []int{5, 9, 12, 13, 16, 45, 2, 12}

	root := generateHuffmanTree(arr, freq)
	codes := generateHuffmanCodes(root)

	PrintHuffmanCodes(codes)
}
