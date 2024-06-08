package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	catRunCnt    = 4
	catRunWidth  = 20
	catRunHeight = 17
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
	spawnLocation [][]int
}

func initPlayer(g *Game) *Player {
	return &Player{
		game:   g,
		facing: 1,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(float64(p.facing), 1)
	opt.GeoM.Translate(32, 320)
	playAni(screen, p.game.sprites["cat_run"].img, opt, catRunWidth, catRunHeight, catRunCnt, p.game.count)
}

func playAni(screen, img *ebiten.Image, opt *ebiten.DrawImageOptions, imgWidth, imgHeight, imgCnt, frameCnt int) {
	i := (frameCnt / 5) % imgCnt
	x0 := 0
	y0 := i * imgHeight
	x1 := x0 + imgWidth
	y1 := y0 + imgHeight
	screen.DrawImage(img.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image), opt)
}
