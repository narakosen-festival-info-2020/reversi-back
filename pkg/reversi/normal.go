package reversi

import (
	"math/rand"
)

// NormalBoard is normal board type string
const NormalBoard = "normal"

// GenerateNormalReversi is Return the data of a normal Reversi board.
func GenerateNormalReversi() Data {
	ret := Data{
		boardType: NormalBoard,
		height:    8,
		width:     8,
		countTurn: 1,
		whoTurn:   1,
		isGameEnd: false,
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

// RandomPlaceAgent is Place a stone somewhere in the board where agent can place it.
// Return true when agent do.
func RandomPlaceAgent(data *Data, agent int) bool {
	if agent != data.whoTurn {
		return false
	}
	type coordinate struct {
		y int
		x int
	}
	var actions []coordinate
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			tmp, _ := data.CanPlaceStone(y, x, agent)
			if tmp {
				actions = append(actions, coordinate{y, x})
			}
		}
	}
	if len(actions) == 0 {
		return false
	}
	sel := rand.Intn(len(actions))
	data.PlaceStone(actions[sel].y, actions[sel].x, agent)
	return true
}
