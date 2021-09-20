package word_search

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, count int) bool {
	if count == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != word[count] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = byte(' ')
	found := dfs(board, word, i+1, j, count+1) || dfs(board, word, i-1, j, count+1) || dfs(board, word, i, j+1, count+1) || dfs(board, word, i, j-1, count+1)

	board[i][j] = tmp

	return found
}
