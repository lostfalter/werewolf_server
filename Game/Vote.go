package game

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

	voteTable := make(map[Player]int)

	for _, vote := range votes {
		voteTable[vote.target]++
	}

	candidates := make([]Player, 0)
	count := 0
	for k, v := range voteTable {
		if len(candidates) == 0 {
			candidates = append(candidates, k)
			count = v
		} else {
			if v > count {
				candidates = nil
				candidates = append(candidates, k)
				count = v
			} else if v == count {
				candidates = append(candidates, k)
			}
		}
	}

	result.target = candidates

	return result
}
