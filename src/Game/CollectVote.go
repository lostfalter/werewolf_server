package game

import (
	"fmt"
	"math/rand"
	"time"
)

// CollectVote from players
func CollectVote(players []Player, targets []Player) []Vote {
	return getRandomVote(players, targets)
}

func getVotesSample(players []Player) []Vote {
	votes := make([]Vote, len(players))

	for i, v := range votes {
		v.player = players[i]
		v.target = players[(i+1)%len(players)]
	}

	return votes
}

func getRandomVote(players []Player, targets []Player) []Vote {
	votes := make([]Vote, len(players))

	rand.Seed(time.Now().UnixNano())

	fmt.Println("\nOriginal votes:")

	for i, v := range votes {
		v.player = players[i]
		targetIndex := rand.Intn(len(targets))

		fmt.Printf("%d->%d, ", players[i].id, targets[targetIndex].id)
		votes[i].player = players[i]
		votes[i].target = targets[targetIndex]

	}

	fmt.Println()

	return votes
}
