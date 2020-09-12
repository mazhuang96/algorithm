/*
 * @Author: mazhuang
 * @Date: 2020-09-12 10:34:35
 * @LastEditors: mazhuang
 * @LastEditTime: 2020-09-12 12:38:16
 * @FilePath: /algorithm/leetcode/637.go
 * @Description:
 */

package leetcode

/*
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

示例 1：

输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。

提示：节点值的范围在32位有符号整数范围内。
*/

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

func averageOfLevels(root *TreeNode) (averages []float64) {
	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		sum := 0
		curLevel := nextLevel
		nextLevel = nil
		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		averages = append(averages, float64(sum)/float64(len(curLevel)))
	}
	return

}

func dfsAverageOfLevels(root *TreeNode) []float64 {
	type data struct {
		sum   int
		count int
	}
	d := []data{}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(d) {
			d[level].sum += node.Val
			d[level].count++
		} else {
			d = append(d, data{node.Val, 1})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	averages := []float64{}
	for _, data := range d {
		averages = append(averages, float64(data.sum)/float64(data.count))
	}
	return averages
}
