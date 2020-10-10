package reversi

// JSONData is JSON of Data
type JSONData struct {
	BoardType string  `json:"board_type"`
	Height    int     `json:"height"`
	Width     int     `json:"width"`
	CountTurn int     `json:"count_turn"`
	WhoTurn   int     `json:"who_turn"`
	IsGameEnd bool    `json:"is_game_end"`
	Board     [][]int `json:"board"`
}

// GetJSON is convert Data to JSONData
func (data *Data) GetJSON() JSONData {
	return JSONData{
		BoardType: data.boardType,
		Height:    data.height,
		Width:     data.width,
		CountTurn: data.countTurn,
		WhoTurn:   data.whoTurn,
		IsGameEnd: data.isGameEnd,
		Board:     data.board,
	}
}
