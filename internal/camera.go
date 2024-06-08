package internal

import (
	"math"
	"meows-adventure/package/util"
)

type Camera struct {
	player    *Player
	lvl       *Level
	x         float64
	y         float64
	targetX   int
	targetY   int
	xDrag     int
	yDrag     int
	smoothing float64
}

func (c *Camera) Update() {
	c.targetX = util.Clamp(c.targetX, c.player.x-c.xDrag, c.player.x+c.xDrag)
	c.targetY = util.Clamp(c.targetY, c.player.y-c.yDrag, c.player.y+c.yDrag)

	xLimit := c.lvl.width/2 - ScreenWidth/2
	yLimit := c.lvl.height/2 - ScreenHeight/2
	c.targetX = util.Clamp(c.targetX, -xLimit, xLimit)
	c.targetY = util.Clamp(c.targetY, -yLimit, yLimit)

	c.x += (float64(c.targetX) - c.x) * (1 - c.smoothing)
	c.y += (float64(c.targetY) - c.y) * (1 - c.smoothing)

	if math.Abs(float64(c.targetX)-c.x) < 0.2 {
		c.x = float64(c.targetX)
	}
	if math.Abs(float64(c.targetY)-c.y) < 0.2 {
		c.y = float64(c.targetY)
	}
}
