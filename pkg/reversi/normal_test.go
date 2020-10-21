package reversi

import (
	"math/rand"
	"testing"
	"time"
)

// WARN: I'm not going to test it at all !!

func TestNormalReversi(t *testing.T) {
	reveriState := createNormalReversi()
	for !reveriState.isGameEnd {
		time.Sleep(time.Second / 2)
		t.Logf("Turn: %d\n", reveriState.countTurn)
		t.Logf("Now Agent: %d\n", reveriState.whoTurn)
		if reveriState.whoTurn != 1 || !reveriState.canPlace {
			continue
		}
		rand.Seed(time.Now().Unix())
		cnt := rand.Intn(7)
		yPlace, xPlace := 0, 0
		for y := 0; y < reveriState.height; y++ {
			for x := 0; x < reveriState.width; x++ {
				tmp, _ := reveriState.CanPlaceStone(y, x, reveriState.whoTurn)
				if tmp {
					yPlace, xPlace = y, x
					if cnt == 0 {
						break
					}
					cnt--
				}
			}
		}
		t.Logf("Place Coordinate (%d, %d)\n", yPlace, xPlace)
		tmp := reveriState.PlaceStone(yPlace, xPlace, reveriState.whoTurn, true)
		t.Logf("Move Stones: %d\n", tmp)
		for _, tmp := range reveriState.board {
			t.Log(tmp)
		}
		if tmp == 0 || tmp == 1 {
			t.Fatalf("\nactuality： %d\nideal： over 2", tmp)
		}
	}
}
