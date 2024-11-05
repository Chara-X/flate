package flate

import (
	"container/heap"
	"math/bits"
	"slices"

	"github.com/Chara-X/priority"
)

type Compressor struct {
	*BitWriter
}

func (c *Compressor) Write(b []byte) (n int, err error) {
	c.WriteBits(2, 3)
	return 0, nil
}

func (c *Compressor) WriteX(b []byte) (n int, err error) {
	var freqs = make([]int, 257)
	for _, v := range b {
		freqs[v]++
	}
	freqs[256] = 1
	var pq = &priority.Queue[*huffmanNode]{}
	heap.Init(pq)
	for char, freq := range freqs {
		heap.Push(pq, &huffmanNode{char: char, freq: freq})
	}
	for pq.Len() > 1 {
		var left = heap.Pop(pq).(*huffmanNode)
		var second = heap.Pop(pq).(*huffmanNode)
		heap.Push(pq, &huffmanNode{freq: left.freq + second.freq, left: left, right: second})
	}
	var root = heap.Pop(pq).(*huffmanNode)
	var codeLens, bitCounts = make([]int, 257), map[int]int{}
	dfs(root, 0, codeLens, bitCounts)
	var code = uint16(0)
	var codes = map[int]huffmanCode{}
	for nb, n := range bitCounts {
		code <<= 1
		if nb == 0 || n == 0 {
			continue
		}
		var chunk = make([]*huffmanNode, 0, n)
		for i := 0; i < n; i++ {
			chunk = append(chunk, pq.Pop().(*huffmanNode))
		}
		slices.SortFunc(chunk, func(a, b *huffmanNode) int { return a.char - b.char })
		for _, node := range chunk {
			codes[node.char] = huffmanCode{code: bits.Reverse16(code << (16 - nb)), len: uint16(nb)}
			code++
		}
	}
	return 0, nil
}

type huffmanCode struct {
	code uint16
	len  uint16
}

func dfs(node *huffmanNode, depth int, codeLens []int, bitCounts map[int]int) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		bitCounts[depth]++
		codeLens[node.char] = depth
		return
	}
	dfs(node.left, depth+1, codeLens, bitCounts)
	dfs(node.right, depth+1, codeLens, bitCounts)
}
func (c *Compressor) Close() error { return nil }

type huffmanNode struct {
	char  int
	freq  int
	left  *huffmanNode
	right *huffmanNode
}

func (n *huffmanNode) Less(v priority.Ordered) bool {
	return n.freq > v.(*huffmanNode).freq
}
