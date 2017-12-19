package game

import "fmt"

// Start game procedure
func Start() bool {
	players := GeneratePlayers()

	var context Context
	context.alivePlayers = players

	Election(context)

	for isGameEnd(&context) == 3 {
		announcePublicStatus(players)

		Exile(&context)

		if isGameEnd(&context) != 3 {
			break
		}

		nightOperation(&context)

		//Exile(context)
	}

	return isGameEnd(&context) == 2
}

func nightOperation(context *Context) {

	WolveOperation(context)

	//WitchOperation(context.alivePlayers)

	//GuardOperation(context.alivePlayers)

	//FarseerOperation(context.alivePlayers)

	//HunterOperation(context.alivePlayers)
}

func isGameEnd(context *Context) int {
	players := context.alivePlayers
	wolveCount := countRole(players, "wolve")
	citizenCount := countRole(players, "citizen")
	godCount := len(players) - wolveCount - citizenCount
	if wolveCount > len(players)-wolveCount ||
		citizenCount == 0 || godCount == 0 {
		fmt.Println("Bad guys win!")
		return 1
	} else if wolveCount == 0 {
		fmt.Println("Good guys win!")
		return 2
	}
	return 3
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

func (context *Context) GetPlayersByRole(name string) []Player {
	players := make([]Player, 0)
	for _, p := range context.alivePlayers {
		if p.role.Name == name {
			players = append(players, p)
		}
	}

	return players
}

func (context *Context) GetWolve() []Player {
	return context.GetPlayersByRole("wolve")
}

func (context *Context) GetGodGuys() []Player {
	players := make([]Player, 0)
	players = append(players, context.GetPlayersByRole("citizen")...)
	players = append(players, context.GetPlayersByRole("farseer")...)
	players = append(players, context.GetPlayersByRole("idiom")...)
	players = append(players, context.GetPlayersByRole("witch")...)
	players = append(players, context.GetPlayersByRole("hunter")...)
	return players
}
