package internal

const (
	frameWidth    = 19
	frameHeight   = 17
	runFrameCount = 4
	frame0X       = 0
	frame0Y       = 0
)

type Player struct {
	game          *Game
	x             int
	y             int
	xVelocity     int
	yVelocity     int
	facing        int
	topSpeed      int
	size          int
	landed        bool
	jumpVelocity  int
	maxFallSpeed  int
	acceleration  int
	climbingSpeed int
	climbing      bool
	touchingVine  bool
}

func (p *Player) Draw() {
	//aniSpeed := p.game.count / 5 // 12 FPS

}
