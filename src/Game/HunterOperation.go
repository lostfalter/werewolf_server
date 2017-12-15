package game

// HunterOperation view hunter status
func HunterOperation(players []Player) Player {
	votes := CollectVote(players)

	result := GetVoteResult(votes)

	return result.target[0]
}
