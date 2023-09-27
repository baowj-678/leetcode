/**
 * @author: baowj
 * @time: 2023/9/4 9:50
 */
package main

import (
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor449() Codec {
	return Codec{}
}

func (this *Codec) putInt32(x int32) []byte {
	buf := make([]byte, 4)
	i := 0
	for x >= 0x80 {
		buf[i] = byte(x) | 0x80
		x >>= 7
		i++
	}
	buf[i] = byte(x)
	return buf
}

func (this *Codec) getInt32(buf []byte) int {
	res := 0
	res += int(buf[0])
	res += int(buf[1]) << 7
	res += int(buf[2]) << 14
	res += int(buf[3]) << 21
	return res
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	start := 0
	var mid func(root *TreeNode, builder *strings.Builder)
	var pre func(root *TreeNode, builder *strings.Builder)
	mid = func(root *TreeNode, builder *strings.Builder) {
		if root == nil {
			return
		}
		builder.Write(this.putInt32(int32(root.Val)))
		mid(root.Left, builder)
		mid(root.Right, builder)
	}
	pre = func(root *TreeNode, builder *strings.Builder) {
		if root == nil {
			return
		}
		pre(root.Left, builder)
		builder.Write(this.putInt32(int32(root.Val)))
		pre(root.Right, builder)
	}
	builder := &strings.Builder{}
	mid(root, builder)
	start = builder.Len()
	pre(root, builder)
	builder.Write(this.putInt32(int32(start)))
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	byteData := []byte(data)
	newData := make([]int, len(data)/4)
	for i := range newData {
		newData[i] = this.getInt32(byteData[i*4 : i*4+4])
	}

	start := newData[len(newData)-1] / 4

	mid := newData[:start]
	pre := newData[start : len(newData)-1]

	var decode func(mid []int, pre []int) *TreeNode
	decode = func(mid []int, pre []int) *TreeNode {
		if len(mid) == 0 {
			return nil
		}
		val := mid[0]
		m := 0
		for i := range pre {
			if pre[i] == val {
				m = i
				break
			}
		}
		root := &TreeNode{
			Val:   int(val),
			Left:  nil,
			Right: nil,
		}
		root.Left = decode(mid[1:m+1], pre[:m])
		root.Right = decode(mid[m+1:], pre[m+1:])
		return root
	}
	return decode(mid, pre)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
