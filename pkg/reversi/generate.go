package reversi

import "fmt"

// CustomBoard is custom board type string
const CustomBoard = "custom"

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
		return CreateNormalReversi(), nil
	case CustomBoard:
		if data.Height > 20 || data.Width > 20 || data.Height < 4 || data.Width < 4 {
			break
		}
		return Data{
			boardType: CustomBoard,
			height:    data.Height,
			width:     data.Width,
			countTurn: 1,
			whoTurn:   1,
			isGameEnd: false,
			canPlace:  true,
			board:     data.Board,
		}, nil
	}
	return Data{}, fmt.Errorf("Invalid Board Type")
}
