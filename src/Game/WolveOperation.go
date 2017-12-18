package game

// WolveOperation kill people
func WolveOperation(players []Player) Player {
	votes := CollectVote(players, players)

	result := GetVoteResult(votes)

	return result.target[0]
}
