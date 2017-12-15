package game

// Exile procedure
func Exile(players []Player) Player {
	votes := CollectVote(players)

	result := GetVoteResult(votes)
	if len(result.target) > 1 {
		votes = CollectVote(players)
		result = GetVoteResult(votes)
	}

	if len(result.target) > 1 {
		// invoke another talk session

		votes = CollectVote(players)
		result = GetVoteResult(votes)
	}

	return result.target[0]
}
