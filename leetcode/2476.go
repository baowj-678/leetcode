/**
 * @author: baowj
 * @time: 2022/11/22 23:50:47
 * @email: bwj_678@qq.com
 */
package main

func closestNodes(root *TreeNode, queries []int) [][]int {
	res := make([][]int, len(queries))
	for i, query := range queries {
		l, r := f2376(root, query)
		res[i] = []int{l, r}
	}
	return res
}

func f2376(root *TreeNode, x int) (l, r int) {
	return l2376(root, x), r2376(root, x)
}

func l2376(root *TreeNode, x int) int {
	if root == nil {
		return -1
	}
	if root.Val == x {
		return x
	} else if root.Val > x {
		return l2376(root.Left, x)
	} else {
		tmp := l2376(root.Right, x)
		if tmp == -1 {
			return root.Val
		} else {
			return tmp
		}
	}
}

func r2376(root *TreeNode, x int) int {
	if root == nil {
		return -1
	}
	if root.Val == x {
		return x
	} else if root.Val < x {
		return r2376(root.Right, x)
	} else {
		tmp := r2376(root.Left, x)
		if tmp == -1 {
			return root.Val
		} else {
			return tmp
		}
	}
}
