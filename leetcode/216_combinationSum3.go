/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 10:57:56
 * @LastEditTime: 2020-09-11 17:05:11
 * @Description:
 **************************************/
package leetcode

import (
	"fmt"
)

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	list := []int{}
	findNum := func(start int, n int) {}
	findNum = func(start int, n int) {
		fmt.Printf("res: %v, &res: %p, start: %2d,  n: %3d, &list:%p,  list %v ,\n", res, res, start, n, list, list)
		// 结束条件
		if len(list) == k || n <= 0 {
			if len(list) == k && n == 0 {
				// err: 直接 append ， res[]内元素内存地址是list的地址不变, list值变化后会被替换
				// res = append(res, list)
				// 需要重新分配内存
				res = append(res, append([]int(nil), list...))
			}
			return
		}
		for i := start; i <= 9; i++ {
			// 选择当前数字
			list = append(list, i)
			// 开始递归
			findNum(i+1, n-i)
			// 回溯（取消选择）
			list = list[:len(list)-1]
		}
	}

	findNum(1, n)

	return res
}
