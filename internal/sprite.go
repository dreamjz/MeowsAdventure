package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"path/filepath"
)

type Sprite struct {
	name string
	path string
	img  *ebiten.Image
}

func initSprite(name string) (*Sprite, error) {
	path := filepath.Join(spritePath, name)
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sprite: %e", err)
	}
	sprite := &Sprite{
		name: name,
		path: path,
		img:  img,
	}
	return sprite, nil
}
