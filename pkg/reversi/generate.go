package reversi

import (
	"fmt"
)

// CustomBoard is custom board type string
const CustomBoard = "custom"

// NormalBoard is normal board type string
const NormalBoard = "normal"

// CircleBoard is circle board type string
const CircleBoard = "circle"

func createNormalReversi() Data {
	ret := Data{
		boardType: NormalBoard,
		height:    8,
		width:     8,
		countTurn: 1,
		whoTurn:   1,
		isGameEnd: false,
		canPlace:  true,
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

func createCircleReversi() Data {
	ret := Data{
		boardType: CircleBoard,
		height:    8,
		width:     8,
		countTurn: 1,
		whoTurn:   1,
		isGameEnd: false,
		canPlace:  true,
		board:     make([][]int, 8),
	}
	left, right := 3, 4
	for i := 0; i < 8; i++ {
		ret.board[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			if j < left || right < j {
				ret.board[i][j] = -1
			}
		}
		if i == 3 {
			ret.board[3][3] = 1
			ret.board[3][4] = 2
		} else if i == 4 {
			ret.board[4][3] = 2
			ret.board[4][4] = 1
		}
		if i < 3 {
			left--
			right++
		} else if i >= 4 {
			left++
			right--
		}
	}
	return ret
}

func createCustomReversi(gen *GenerateData) (Data, error) {
	ret := Data{
		boardType: CustomBoard,
		height:    gen.Height,
		width:     gen.Width,
		countTurn: 1,
		whoTurn:   1,
		isGameEnd: false,
		canPlace:  true,
		board:     gen.Board,
	}
	if ret.height > 20 || ret.width > 20 || ret.height < 4 || ret.width < 4 {
		return Data{}, fmt.Errorf("Invalid Board Size")
	}
	if ret.height != len(ret.board) {
		return Data{}, fmt.Errorf("Height not match")
	}
	check := true
	for i := 0; i < ret.height; i++ {
		if ret.width != len(ret.board[i]) {
			check = false
		}
	}
	if !check {
		return Data{}, fmt.Errorf("Width not match")
	}
	if !ret.canPlaceByAgent(1) {
		return Data{}, fmt.Errorf("Black must can place")
	}
	return ret, nil
}

// GenerateData is base data of generate reversi board
type GenerateData struct {
	BoardType string  `json:"board_type"`
	Height    int     `json:"height"`
	Width     int     `json:"width"`
	Board     [][]int `json:"board"`
}

// Create is create board based on GenerateData
func (data *GenerateData) Create() (Data, error) {
	switch data.BoardType {
	case NormalBoard:
		return createNormalReversi(), nil
	case CircleBoard:
		return createCircleReversi(), nil
	case CustomBoard:
		return createCustomReversi(data)
	}
	return Data{}, fmt.Errorf("Invalid Board Type")
}
