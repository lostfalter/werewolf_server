package game

import "fmt"

// Election for captatin
func Election(context Context) Player {
	fmt.Print("\n-------------\nStart Election\n")

	players := context.alivePlayers

	Talk(players)

	votes := CollectVote(players, players)
	result := GetVoteResult(votes)
	if len(result.target) > 1 {
		// invoke another talk session
		Talk(players)

		votes = CollectVote(players, result.target)
		result = GetVoteResult(votes)
	}

	// if len(result.target) > 1 {
	// 	// invoke another talk session
	// 	Talk(players)

	// 	votes = CollectVote(players)
	// 	result = GetVoteResult(votes)
	// }

	fmt.Print("End Election\n-------------\n")

	return result.target[0]
}
