/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-26 09:14:24
 * @LastEditTime: 2020-09-26 10:17:45
 * @Description:
 **************************************/
package leetcode

import "fmt"

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

*/

/**
在 root==nil 返回结果会返回两次， 叶子节点要走左右都为 nil
*/

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	var path []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum -= root.Val
		fmt.Println("root.val:", root.Val)
		path = append(path, root.Val)
		fmt.Println("path:", path)
		if root.Left == nil && root.Right == nil && sum == 0 {
			res = append(res, append([]int(nil), path...))
		}
		dfs(root.Left)
		dfs(root.Right)
		sum += root.Val
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}
