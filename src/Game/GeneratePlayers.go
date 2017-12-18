package game

import (
	"fmt"
)

// GeneratePlayers generate players
func GeneratePlayers() []Player {
	var roles = basicConfiguration()
	roles = shuffle(roles)

	fmt.Println(roles)

	players := make([]Player, len(roles))
	for i, r := range roles {
		players[i].role = r
		players[i].id = i + 1
	}

	return players
}

// shuffle role list
func shuffle(roles []Role) []Role {
	return roles
}

// basicConfiguration baisc twelve people configuration
func basicConfiguration() []Role {
	roles := make([]Role, 0)

	wolve := Role{"wolve"}
	citizen := Role{"citizen"}
	farseer := Role{"farseer"}
	witch := Role{"witch"}
	hunter := Role{"hunter"}
	idiom := Role{"idiom"}

	for i := 0; i < 4; i++ {
		roles = append(roles, wolve)
	}

	for i := 0; i < 4; i++ {
		roles = append(roles, citizen)
	}

	roles = append(roles, farseer)
	roles = append(roles, witch)
	roles = append(roles, hunter)
	roles = append(roles, idiom)

	return roles
}
