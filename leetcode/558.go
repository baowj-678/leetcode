/**
 * @author: baowj
 * @time: 2022/7/15 12:02:54
 */
package main

/**
 * Definition for a QuadTree node.
 */

type Node558 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node558
	TopRight    *Node558
	BottomLeft  *Node558
	BottomRight *Node558
}

func intersect(quadTree1 *Node558, quadTree2 *Node558) *Node558 {
	if quadTree1.IsLeaf && quadTree2.IsLeaf {
		quadTree1.Val = quadTree1.Val || quadTree2.Val
		return quadTree1
	} else if quadTree1.IsLeaf {
		quadTree2.TopLeft = intersect(quadTree2.TopLeft, quadTree1)
		quadTree2.TopRight = intersect(quadTree2.TopRight, quadTree1)
		quadTree2.BottomLeft = intersect(quadTree2.BottomLeft, quadTree1)
		quadTree2.BottomRight = intersect(quadTree2.BottomRight, quadTree1)
		isMerge(quadTree2)
		return quadTree2
	} else if quadTree2.IsLeaf {
		quadTree1.TopLeft = intersect(quadTree1.TopLeft, quadTree2)
		quadTree1.TopRight = intersect(quadTree1.TopRight, quadTree2)
		quadTree1.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2)
		quadTree1.BottomRight = intersect(quadTree1.BottomRight, quadTree2)
		isMerge(quadTree1)
		return quadTree1
	} else {
		quadTree1.TopLeft = intersect(quadTree1.TopLeft, quadTree2.TopLeft)
		quadTree1.TopRight = intersect(quadTree1.TopRight, quadTree2.TopRight)
		quadTree1.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
		quadTree1.BottomRight = intersect(quadTree1.BottomRight, quadTree2.BottomRight)
		isMerge(quadTree1)
		return quadTree1
	}
}

func isMerge(quadTree *Node558) bool {
	a := quadTree.TopLeft
	b := quadTree.TopRight
	c := quadTree.BottomLeft
	d := quadTree.BottomRight
	if a.IsLeaf && b.IsLeaf && c.IsLeaf && d.IsLeaf {
		if a.Val == b.Val && a.Val == c.Val && a.Val == d.Val {
			quadTree.IsLeaf = true
			quadTree.Val = a.Val
			quadTree.TopLeft = nil
			quadTree.TopRight = nil
			quadTree.BottomLeft = nil
			quadTree.BottomRight = nil
			return true
		}
	}
	return false
}
