package game

import (
	"fmt"
)

// Exile procedure
func Exile(context Context) Context {
	players := context.alivePlayers

	fmt.Print("\n-------------\nStart Exile\n")

	Talk(players)

	voters := players
	votes := CollectVote(voters, players)
	result := GetVoteResult(votes)
	if len(result.target) > 1 {
		// invoke another talk session
		fmt.Println("Multi candidates")

		Talk(result.target)

		voters = RemovePlayers(players, result.target)
		votes = CollectVote(voters, result.target)
		result = GetVoteResult(votes)
	}

	// if len(result.target) > 1 {
	// 	// invoke another talk session
	// 	fmt.Println("Multi candidates")
	// 	fmt.Println("Start Talk")
	// 	Talk(players)

	// 	votes = CollectVote(players)
	// 	result = GetVoteResult(votes)
	// }

	target := result.target[0]

	players = RemovePlayer(players, target)

	target.status.liveStatus = Dead
	context.deadPlayers = append(context.deadPlayers, target)
	context.alivePlayers = players

	fmt.Print("End Exile\n-------------\n")

	return context
}
