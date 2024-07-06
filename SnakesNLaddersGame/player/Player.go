package player

var id = 0

type Player struct {
	userId           int
	name             string
	position         int
	noofTurnsBlocked int
}

type Players []*Player

func NewPlayer(nm string, pos int) *Player {
	id += 1
	return &Player{userId: id, name: nm, position: pos, noofTurnsBlocked: 0}
}

func (p *Player) GetId() int {
	return p.userId
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetCurrentPosition() int {
	return p.position
}

func (p *Player) GetBlockedTurns() int {
	return p.noofTurnsBlocked
}

func (p *Player) SetBlockTurns(turns int) {
	p.noofTurnsBlocked = turns
}

func (p *Player) Move(pos int) {
	p.position = pos
}
