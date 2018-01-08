package game

import (
	"fmt"
	"math/rand"
)

// GeneratePlayers generate players
func GeneratePlayers() []Player {
	var roles = basicConfiguration()
	roles = shuffle(roles)

	players := make([]Player, len(roles))
	for i, r := range roles {
		players[i].role = r
		players[i].id = i + 1
	}

	for _, v := range players {
		fmt.Println(v)
	}

	return players
}

// shuffle role list
func shuffle(src []Role) []Role {
	dest := make([]Role, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

// basicConfiguration baisc twelve people configuration
func basicConfiguration() []Role {
	roles := make([]Role, 0)

	wolf := Role{"wolf"}
	villager := Role{"villager"}
	seer := Role{"seer"}
	witch := Role{"witch"}
	hunter := Role{"hunter"}
	idiom := Role{"idiom"}

	for i := 0; i < 4; i++ {
		roles = append(roles, wolf)
	}

	for i := 0; i < 4; i++ {
		roles = append(roles, villager)
	}

	roles = append(roles, seer)
	roles = append(roles, witch)
	roles = append(roles, hunter)
	roles = append(roles, idiom)

	return roles
}
