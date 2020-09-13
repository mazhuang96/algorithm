/*
 * @Author: mazhuang
 * @Date: 2020-09-13 13:27:17
 * @LastEditors: mazhuang
 * @LastEditTime: 2020-09-13 14:27:58
 * @FilePath: /algorithm/leetcode/79_seachWord.go
 * @Description:
 */
package leetcode

func exist(board [][]byte, word string) bool {

	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	directions := [][]int{
		{0, -1}, // 向上搜素
		{-1, 0}, // 向左搜索
		{0, 1},  // 向下搜索
		{1, 0},  // 向右搜索
	}
	var marked [][]bool
	for i := 0; i < m; i++ {
		var mark []bool
		for j := 0; j < n; j++ {
			mark = append(mark, false)
		}
		marked = append(marked, mark)
	}
	var dfs func(cur, x, y int) bool
	dfs = func(cur, x, y int) bool {
		if cur == len(word)-1 {
			return board[x][y] == word[cur]
		}
		if board[x][y] == word[cur] {
			marked[x][y] = true

			for _, dir := range directions {
				nx := x + dir[0]
				ny := y + dir[1]
				if nx >= 0 && nx < m && ny < n && ny >= 0 &&
					!marked[nx][ny] &&
					dfs(cur+1, nx, ny) {
					return true
				}
			}
			marked[x][y] = false
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
