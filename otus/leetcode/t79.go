package leetcode

// 79. Word Search
// https://leetcode.com/problems/word-search/

func exist(board [][]byte, word string) bool {

	type key struct {
		i, j int
	}

	visited := make(map[key]struct{})

	var backtrack func(i, j, k int) bool

	backtrack = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}

		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}

		if board[i][j] != word[k] {
			return false
		}

		key := key{i, j}

		if _, ok := visited[key]; ok {
			return false
		}
		visited[key] = struct{}{}

		res := backtrack(i-1, j, k+1) || backtrack(i+1, j, k+1) || backtrack(i, j-1, k+1) || backtrack(i, j+1, k+1)

		delete(visited, key)

		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res := backtrack(i, j, 0)
			if res {
				return true
			}
		}
	}

	return false
}
