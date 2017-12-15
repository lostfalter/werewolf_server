package game

// FarseerOperation view people role
func FarseerOperation(players []Player) Player {
	votes := CollectVote(players)

	result := GetVoteResult(votes)

	return result.target[0]
}
