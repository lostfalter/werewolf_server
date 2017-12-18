package game

import "fmt"

type LiveStatus int

const (
	Live LiveStatus = iota
	Dead
)

func (s LiveStatus) String() string {
	var result string
	if s == Live {
		result = "Live"
	} else {
		result = "Dead"
	}

	return result
}

type TalkStatus int

const (
	Normal TalkStatus = iota
	Forbidden
)

func (s TalkStatus) String() string {
	var result string
	if s == Normal {
		result = "Normal"
	} else {
		result = "Forbidden"
	}

	return result
}

type RoleStatus struct {
	liveStatus LiveStatus
	talkStatus TalkStatus
}

func (s RoleStatus) String() string {
	return s.liveStatus.String() + " " + s.talkStatus.String()
}

type Role struct {
	Name string
}

func (r Role) String() string {
	return r.Name
}

type Player struct {
	role   Role
	status RoleStatus
	id     int
}

func (p Player) String() string {
	info := fmt.Sprintf("Player\nId: %d\nLive status: %s\nTalk status: %s\n",
		p.id, p.status.liveStatus.String(), p.status.talkStatus.String())
	return info
}

func (player *Player) IsValid() bool {

	return player.id != -1
}
