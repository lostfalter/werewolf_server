package game

// Start game procedure
func Start() {
	players := GeneratePlayers()

	Election(players)

	for !isGameEnd(players) {
		announcePublicStatus(players)

		Exile(players)

		if !isGameEnd(players) {
			break
		}

		nightOperation(players)

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
