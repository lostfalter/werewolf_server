package game

import (
	"fmt"
)

// Start game procedure
func Start() {
	players := GeneratePlayers()

	var context Context
	context.alivePlayers = players

	Election(context)

	for !isGameEnd(players) {
		announcePublicStatus(players)

		context = Exile(context)

		fmt.Println(context.alivePlayers)
		context = Exile(context)

		if isGameEnd(players) {
			break
		}

		// nightOperation(players)

		//Exile(context)

		break
	}

}

func nightOperation(players []Player) {

	WolveOperation(players)

	WitchOperation(players)

	GuardOperation(players)

	FarseerOperation(players)

	HunterOperation(players)
}

func isGameEnd(players []Player) bool {
	return false
}

func announcePublicStatus(players []Player) {

}

// Context for game
type Context struct {
	alivePlayers []Player
	deadPlayers  []Player
}
