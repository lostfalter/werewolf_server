package game

// GuardOperation guard people
func GuardOperation(players []Player) Player {
	votes := CollectVote(players, players)

	result := GetVoteResult(votes)

	return result.target[0]
}
