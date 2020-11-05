package reversi

import (
	"math/rand"
)

// RandomPlaceAgent is Place a stone somewhere in the board where agent can place it.
// Return true when agent do.
func RandomPlaceAgent(data *Data, agent int, doPlace chan bool) bool {
	if agent != data.whoTurn {
		return false
	}
	type coordinate struct {
		y int
		x int
	}
	var actions []coordinate
	for y := 0; y < data.height; y++ {
		for x := 0; x < data.width; x++ {
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
	tmp := <-doPlace
	if tmp {
		data.PlaceStone(actions[sel].y, actions[sel].x, agent, false)
	}
	return true
}
