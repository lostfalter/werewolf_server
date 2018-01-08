package game

import (
	"reflect"
	"testing"
)

func TestRemovePlayer(t *testing.T) {
	type args struct {
		players []Player
		t       Player
	}

	type testcase struct {
		name string
		args args
		want []Player
	}
	var tests []testcase

	var player1 Player
	player1.id = 1
	player1.role.Name = "villager"

	var player2 Player
	player2.id = 2
	player2.role.Name = "wolf"

	var players []Player
	players = append(players, player1)
	players = append(players, player2)

	var basicCase testcase
	basicCase.name = "basic case"
	basicCase.args.players = players
	basicCase.args.t = player1
	basicCase.want = append(basicCase.want, player2)

	tests = append(tests, basicCase)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovePlayer(tt.args.players, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemovePlayers(t *testing.T) {
	type args struct {
		players []Player
		targets []Player
	}

	type testcase struct {
		name string
		args args
		want []Player
	}
	var tests []testcase

	var player1 Player
	player1.id = 1
	player1.role.Name = "villager"

	var player2 Player
	player2.id = 2
	player2.role.Name = "wolf"

	var players []Player
	players = append(players, player1)
	players = append(players, player2)

	var removeTarget []Player
	removeTarget = append(removeTarget, player1)

	var basicCase testcase
	basicCase.name = "basic case"
	basicCase.args.players = players
	basicCase.args.targets = removeTarget
	basicCase.want = append(basicCase.want, player2)

	tests = append(tests, basicCase)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovePlayers(tt.args.players, tt.args.targets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovePlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}
