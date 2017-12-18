package game

import "fmt"

// Vote for action
type Vote struct {
	target Player
	player Player
}

// VoteResult from votes
type VoteResult struct {
	target []Player

	originVotes []Vote
}

// GetVoteResult for votes
func GetVoteResult(votes []Vote) VoteResult {
	var result VoteResult
	result.originVotes = votes

	voteTable := make(map[Player][]Vote)

	for _, v := range votes {
		voteTable[v.target] = append(voteTable[v.target], v)
	}

	fmt.Println("Vote summary:")
	for k, v := range voteTable {
		fmt.Printf("Player %3d <--", k.id)
		for _, v2 := range v {
			fmt.Printf(" %d", v2.player.id)
		}
		fmt.Println()
	}

	fmt.Println()

	candidates := make([]Player, 0)
	count := 0
	for k, v := range voteTable {
		if len(candidates) == 0 {
			candidates = append(candidates, k)
			count = len(v)
		} else {
			if len(v) > count {
				candidates = nil
				candidates = append(candidates, k)
				count = len(v)
			} else if len(v) == count {
				candidates = append(candidates, k)
			}
		}
	}

	result.target = candidates
	fmt.Println("Candidate:")
	for _, v := range result.target {
		fmt.Printf("Player %d\n", v.id)
	}

	return result
}
