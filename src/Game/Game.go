package game

import "fmt"

// Start game procedure
func Start() {
	players := GeneratePlayers()

	var context Context
	context.alivePlayers = players

	Election(context)

	for !isGameEnd(context.alivePlayers) {
		announcePublicStatus(players)

		Exile(&context)

		if isGameEnd(context.alivePlayers) {
			break
		}

		nightOperation(&context)

		//Exile(context)
	}

}

func nightOperation(context *Context) {

	WolveOperation(context)

	//WitchOperation(context.alivePlayers)

	//GuardOperation(context.alivePlayers)

	//FarseerOperation(context.alivePlayers)

	//HunterOperation(context.alivePlayers)
}

func isGameEnd(players []Player) bool {
	wolveCount := countRole(players, "wolve")
	if wolveCount > len(players)-wolveCount {
		fmt.Println("Bad guys win!")
		return true
	} else if wolveCount == 0 {
		fmt.Println("Good guys win!")
		return true
	}
	return false
}

func countRole(players []Player, name string) int {
	count := 0
	for _, p := range players {
		if p.role.Name == name {
			count++
		}
	}

	return count
}

func announcePublicStatus(players []Player) {

}

// Context for game
type Context struct {
	alivePlayers []Player
	deadPlayers  []Player
}

func (context *Context) GetWolve() []Player {
	wolves := make([]Player, 0)
	for _, p := range context.alivePlayers {
		if p.role.Name == "wolve" {
			wolves = append(wolves, p)
		}
		fmt.Println(p.role.Name)
	}

	fmt.Println("wolves:")
	fmt.Println(wolves)
	return wolves
}
