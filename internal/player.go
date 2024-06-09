package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"meows-adventure/package/util"
)

const (
	catRun       = "cat_run"
	catRunCnt    = 4
	catRunWidth  = 20
	catRunHeight = 17
)

type Player struct {
	game          *Game
	x             float64
	y             float64
	xVelocity     float64
	yVelocity     float64
	facing        int
	topSpeed      float64
	size          int
	landed        bool
	jumpVelocity  float64
	maxFallSpeed  int
	acceleration  float64
	climbingSpeed int
	climbing      bool
	touchingVine  bool
	anis          map[string]*Animation
}

type Animation struct {
	width, height int
	imgCnt        int
	count         int
	img           *ebiten.Image
}

func initPlayer(g *Game) *Player {
	anis := map[string]*Animation{}

	catRunAni := &Animation{
		width:  catRunWidth,
		height: catRunHeight,
		imgCnt: catRunCnt,
		img:    g.sprites[catRun].img,
	}

	anis[catRun] = catRunAni
	return &Player{
		game:          g,
		facing:        1,
		topSpeed:      2,
		size:          16,
		landed:        true,
		jumpVelocity:  4.2,
		maxFallSpeed:  8,
		acceleration:  0.4,
		climbingSpeed: 1,
		anis:          anis,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	if p.facing < 0 {
		opt.GeoM.Scale(-1, 1)
		opt.GeoM.Translate(float64(p.size), 1)
	}
	opt.GeoM.Translate(32+p.x, 320+p.y)
	playAni(screen, opt, p.anis[catRun])
}

func (p *Player) Update() {
	// collision
	//mi := p.game.maps["level1"].info
	//oldTop := math.Floor((p.y+128+4)/float64(mi.BlockHeight)+1-float64(mi.Height)) * float64(mi.Width)
	//oldBottom := math.Floor((p.y+128-6)/float64(mi.BlockHeight)+1-float64(mi.Height)) * float64(mi.Width)
	//oldLeft := math.Floor()

	// moving
	var keys []ebiten.Key
	keys = inpututil.AppendPressedKeys(keys)

	if len(keys) > 0 {
		k := keys[0]
		if k == ebiten.KeyArrowLeft {
			p.facing = -1
			p.xVelocity -= p.acceleration
		} else if k == ebiten.KeyArrowRight {
			p.facing = 1
			p.xVelocity += p.acceleration
		}
	}

	p.xVelocity = util.Clamp(p.xVelocity, -p.topSpeed, p.topSpeed)

	p.xVelocity *= 0.85
	if p.xVelocity > -0.05 && p.xVelocity < 0.05 {
		p.xVelocity = 0
	}

	p.x += p.xVelocity

}

func playAni(screen *ebiten.Image, opt *ebiten.DrawImageOptions, ani *Animation) {
	ani.count++
	i := (ani.count / 15) % ani.imgCnt
	x0 := 0
	y0 := i * ani.height
	x1 := x0 + ani.width
	y1 := y0 + ani.height
	if i == ani.imgCnt-1 {
		ani.count = 0
	}
	screen.DrawImage(ani.img.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), opt)
}
