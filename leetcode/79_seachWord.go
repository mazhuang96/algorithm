/*
 * @Author: mazhuang
 * @Date: 2020-09-13 13:27:17
 * @LastEditors: mazhuang
 * @LastEditTime: 2020-09-13 14:40:03
 * @FilePath: /algorithm/leetcode/79_seachWord.go
 * @Description:
 */
package leetcode

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	if m == 0 {
		return false
	}
	directions := [][]int{
		{0, -1}, // 向上搜素
		{-1, 0}, // 向左搜索
		{0, 1},  // 向下搜索
		{1, 0},  // 向右搜索
	}
	var marked = make([][]bool, m)
	for i := range marked {
		marked[i] = make([]bool, n)
	}
	var dfs func(cur, x, y int) bool
	dfs = func(cur, x, y int) bool {
		if board[x][y] != word[cur] {
			return false
		}
		if cur == len(word)-1 {
			return true
		}
		marked[x][y] = true
		defer func() { marked[x][y] = false }() // 回溯自动调用

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n &&
				!marked[nx][ny] &&
				dfs(cur+1, nx, ny) {
				return true
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}
