/**
 * @author: baowj
 * @time: 2023/5/30 9:55
 */
package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delS := map[int]struct{}{}
	for _, v := range to_delete {
		delS[v] = struct{}{}
	}
	res := []*TreeNode{}
	if _, ok := delS[root.Val]; !ok {
		res = append(res, root)
	}
	tmp := dfs1110(root, delS)
	res = append(res, tmp...)
	return res
}

func dfs1110(root *TreeNode, delS map[int]struct{}) []*TreeNode {
	res := []*TreeNode{}
	if _, ok := delS[root.Val]; ok {
		// delete
		delete(delS, root.Val)
		if root.Left != nil {
			if _, ok = delS[root.Left.Val]; !ok {
				res = append(res, root.Left)
			}
			tmp := dfs1110(root.Left, delS)
			res = append(res, tmp...)
		}

		if root.Right != nil {
			if _, ok = delS[root.Right.Val]; !ok {
				res = append(res, root.Right)
			}
			tmp := dfs1110(root.Right, delS)
			res = append(res, tmp...)
		}
	} else {
		// not delete
		if root.Left != nil {
			_, ok = delS[root.Left.Val]
			tmp := dfs1110(root.Left, delS)
			if ok {
				root.Left = nil
			}
			res = append(res, tmp...)
		}

		if root.Right != nil {
			_, ok = delS[root.Right.Val]
			tmp := dfs1110(root.Right, delS)
			if ok {
				root.Right = nil
			}
			res = append(res, tmp...)
		}
	}
	return res
}
