package reversi

// GenerateNormalReversi is Return the data of a normal Reversi board.
func GenerateNormalReversi() Data {
	ret := Data{
		boardType: 1,
		height:    8,
		width:     8,
		countTurn: 1,
		whoTurn:   1,
		IsGameEnd: false,
		board:     make([][]int, 8),
	}
	for i := 0; i < 8; i++ {
		ret.board[i] = make([]int, 8)
		if i == 3 {
			ret.board[3][3] = 1
			ret.board[3][4] = 2
		} else if i == 4 {
			ret.board[4][3] = 2
			ret.board[4][4] = 1
		}
	}
	return ret
}
