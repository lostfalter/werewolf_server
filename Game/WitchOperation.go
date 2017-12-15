package game

// WitchOperation kill people
func WitchOperation(players []Player) Player {
	votes := CollectVote(players)

	result := GetVoteResult(votes)

	return result.target[0]
}
