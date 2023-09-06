/**
 * @author: baowj
 * @time: 2023/6/13 9:55
 */
package main

type TreeAncestor struct {
	parent [][16]int
	n      int
}

func TreeAncestorConstructor(n int, parent []int) TreeAncestor {
	pa := make([][16]int, len(parent))
	for i := 0; i < len(parent); i++ {
		pa[i][0] = parent[i]
	}
	for i := 1; i < 16; i++ {
		for j := 0; j < len(pa); j++ {
			if pa[j][i-1] == -1 {
				pa[j][i] = -1
			} else {
				pa[j][i] = pa[pa[j][i-1]][i-1]
			}
		}
	}

	return TreeAncestor{parent: pa, n: n}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	cnt := 0
	for k > 0 {
		if k&0b1 == 1 {
			node = this.parent[node][cnt]
		}
		if node == -1 {
			return node
		}
		cnt++
		k >>= 1
	}
	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
