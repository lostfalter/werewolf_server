package game

import "fmt"

// Start game procedure
func Start() bool {
	players := GeneratePlayers()

	var context Context
	context.alivePlayers = players

	Election(context)

	for isGameEnd(&context) == 3 {
		announcePublicStatus(&context)

		Exile(&context)

		if isGameEnd(&context) != 3 {
			break
		}

		nightOperation(&context)

		//Exile(context)
	}

	announcePublicStatus(&context)

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
	wolveCount := countRole(players, "wolf")
	citizenCount := countRole(players, "villager")
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

func announcePublicStatus(context *Context) {
	fmt.Println("alive players:")
	for _, v := range context.alivePlayers {
		fmt.Println(v)
	}

	fmt.Println("dead players:")
	for _, v := range context.deadPlayers {
		fmt.Println(v)
	}
}

// Context for game
type Context struct {
	alivePlayers []Player
	deadPlayers  []Player
}

func (context *Context) getPlayersByRole(name string) []Player {
	players := make([]Player, 0)
	for _, p := range context.alivePlayers {
		if p.role.Name == name {
			players = append(players, p)
		}
	}

	return players
}

func (context *Context) GetWolve() []Player {
	return context.getPlayersByRole("wolf")
}

func (context *Context) GetGodGuys() []Player {
	players := make([]Player, 0)
	players = append(players, context.getPlayersByRole("villager")...)
	players = append(players, context.getPlayersByRole("seer")...)
	players = append(players, context.getPlayersByRole("idiom")...)
	players = append(players, context.getPlayersByRole("witch")...)
	players = append(players, context.getPlayersByRole("hunter")...)
	return players
}
