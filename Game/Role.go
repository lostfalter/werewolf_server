package game

type LiveStatus int

const (
	Live LiveStatus = iota
	Dead
)

type TalkStatus int

const (
	Normal TalkStatus = iota
	Forbidden
)

type RoleStatus struct {
	liveStatus LiveStatus
	talkStatus TalkStatus
}

type Role struct {
	Name string
}

type Player struct {
	role   Role
	status RoleStatus
	id     int
}

func (player *Player) IsValid() bool {

	return player.id != -1
}
