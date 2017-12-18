package game

import (
	"fmt"
	"math/rand"
	"time"
)

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
	personCount := len(players)

	votes := make([]Vote, personCount)

	rand.Seed(time.Now().UnixNano())

	for i, v := range votes {
		v.player = players[i]
		targetIndex := rand.Intn(personCount)

		fmt.Printf("%d->%d,\t", i, targetIndex)
		votes[i].player = players[i]
		votes[i].target = players[targetIndex]

	}

	fmt.Println()

	return votes
}
