package game

// CollectVote from players
func CollectVote(players []Player) []Vote {
	return getRandomVote(players)
}

func getVotesSample(players []Player) []Vote {
	votes := make([]Vote, len(players))

	for i, v := range votes {
		v.player = players[i]
		v.target = players[(i+1)%len(players)]
	}

	return votes
}

func getRandomVote(players []Player) []Vote {
	votes := make([]Vote, len(players))

	for i, v := range votes {
		v.player = players[i]
		v.target = players[(i+1)%len(players)]
	}

	return votes
}
