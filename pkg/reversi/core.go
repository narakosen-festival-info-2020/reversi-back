package reversi

import "fmt"

/*
Data is Reversi Core
Black Stone: 1, White Stone: 2, Nothing: 0, Wall: -1
*/
type Data struct {
	boardType int
	height    int
	width     int
	countTurn int
	whoTurn   int
	IsGameEnd bool
	board     [][]int
}

// Return (-1, -1): Cannot place, (a, b): Invert from (y, x) to (a, b)
func (data *Data) canPlaceStoneDirection(y, x, yMove, xMove, myStone int) (int, int, error) {
	if yMove == 0 && xMove == 0 {
		return -1, -1, fmt.Errorf("Cannot move (yMove: %d, xMove: %d)", yMove, xMove)
	}
	indexChcek := func(a, b int) bool {
		return (a >= 0 && a < data.height && b >= 0 && b < data.width)
	}
	if !indexChcek(y, x) {
		return -1, -1, fmt.Errorf("Index out of bounds (y: %d, x: %d)", y, x)
	}
	if data.board[y][x] != 0 {
		return -1, -1, nil
	}
	moveCnt := 0
	move := func() {
		y += yMove
		x += xMove
		moveCnt++
	}
	reverseStone := func(stone int) int {
		if stone == 1 {
			return 2
		}
		return 1
	}
	retY, retX := -1, -1
	for ; indexChcek(y, x); move() {
		if moveCnt == 0 {
			continue
		} else if moveCnt == 1 && data.board[y][x] == myStone {
			return -1, -1, nil
		}
		if data.board[y][x] != reverseStone(myStone) {
			if data.board[y][x] == myStone {
				retY, retX = y, x
			}
			break
		}
	}
	return retY, retX, nil
}

// CanPlaceStone is Determining if a stone can be placed.
func (data *Data) CanPlaceStone(y, x, myStone int) (bool, error) {
	ret := false
	for yMove := -1; yMove <= 1; yMove++ {
		for xMove := -1; xMove <= 1; xMove++ {
			if yMove == 0 && xMove == 0 {
				continue
			}
			yEnd, xEnd, err := data.canPlaceStoneDirection(y, x, yMove, xMove, myStone)
			if err != nil {
				return false, err
			}
			ret = ret || (yEnd != -1 && xEnd != -1)
		}
	}
	return ret, nil
}

// Return count of invert stone.
func (data *Data) invertStone(y, x, myStone int) (int, error) {
	cnt := 0
	invert := func(yMove, xMove, yEnd, xEnd int) {
		yNow, xNow := y+yMove, x+xMove
		for yNow != yEnd || xNow != xEnd {
			cnt++
			data.board[yNow][xNow] = myStone
			yNow += yMove
			xNow += xMove
		}
	}
	for yMove := -1; yMove <= 1; yMove++ {
		for xMove := -1; xMove <= 1; xMove++ {
			if yMove == 0 && xMove == 0 {
				continue
			}
			yEnd, xEnd, err := data.canPlaceStoneDirection(y, x, yMove, xMove, myStone)
			if err != nil {
				return -1, err
			}
			if yEnd != -1 && xEnd != -1 {
				invert(yMove, xMove, yEnd, xEnd)
			}
		}
	}
	return cnt, nil
}

func (data *Data) canPlaceByAgent(myStone int) bool {
	ret := false
	for y := 0; y < data.height; y++ {
		for x := 0; x < data.width; x++ {
			tmp, _ := data.CanPlaceStone(y, x, myStone)
			ret = ret || tmp
		}
	}
	return ret
}

// update data
func (data *Data) turnProgress() {
	changeTurn := func() {
		if data.whoTurn == 1 {
			data.whoTurn = 2
		} else {
			data.whoTurn = 1
		}
	}
	data.countTurn++
	for i := 0; i < 2; i++ {
		changeTurn()
		if data.canPlaceByAgent(data.whoTurn) {
			return
		}
	}
	data.IsGameEnd = true
}

// PlaceStone is Place a stone at the coordinates (y, x) and trun progresses
// Returns over 0 when it can be placed.
func (data *Data) PlaceStone(y, x, myStone int) int {
	if myStone != data.whoTurn {
		return -1
	}
	check, _ := data.CanPlaceStone(y, x, myStone)
	if !check {
		return -1
	}
	cnt, _ := data.invertStone(y, x, myStone)
	data.board[y][x] = myStone
	data.turnProgress()
	return cnt + 1
}
