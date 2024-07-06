package jumperFactory

import "snakesNLadders/player"

type IJumper interface {
	IsValid() error
	Move(p *player.Player)
	GetJumperType() int
}

type Jumper struct {
	StartPoint int
	EndPoint   int
}
