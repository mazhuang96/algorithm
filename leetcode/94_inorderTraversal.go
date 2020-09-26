/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-14 10:36:01
 * @LastEditTime: 2020-09-14 11:33:42
 * @Description: 中序遍历
 **************************************/
package leetcode

/**
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/

// 递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 迭代
func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 获取栈顶元素
		root = stack[len(stack)-1]
		// 移除栈顶元素
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
