package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestUnivaluePath(&TreeNode{Val: 1}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Temp struct {
	Max int
}

func longestUnivaluePath(root *TreeNode) int {
	temp := &Temp{Max: 0}
	if root == nil {
		return 0
	}
	dfs(root, root.Val, temp)
	return temp.Max
}

func dfs(root *TreeNode, val int, temp *Temp) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, root.Val, temp)
	r := dfs(root.Right, root.Val, temp)

	// 若以 root 為當前節點作計算，直接把左右邊長度給加總起來跟 max 比
	temp.Max = int(math.Max(float64(l+r), float64(temp.Max)))

	// 若當前節點不為上一個節點，代表連不起來，直接回傳 0，否則就回傳看左右邊誰比較長 + 1
	if val != root.Val {
		return 0
	}
	return int(math.Max(float64(l), float64(r))) + 1
}
