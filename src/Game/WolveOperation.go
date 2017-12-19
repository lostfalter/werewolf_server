package game

// WolveOperation kill people
func WolveOperation(context *Context) Player {

	votes := CollectVote(context.GetWolve(), context.alivePlayers)

	result := GetVoteResult(votes)

	RemovePlayer(context.alivePlayers, result.target[0])

	return result.target[0]
}
