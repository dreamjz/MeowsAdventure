package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	spritePath = "assets/sprite"
	mapPath    = "assets/map"
)

type Game struct {
	count   int
	player  *Player
	sprites map[string]*Sprite
	maps    map[string]*Map
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	err := g.maps["level1"].Draw(screen)
	if err != nil {
		ebitenutil.DebugPrint(screen, err.Error())
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func (g *Game) Init() error {
	// init sprites
	g.sprites = map[string]*Sprite{}
	entries, err := os.ReadDir(spritePath)
	if err != nil {
		return fmt.Errorf("failed to load sprite : %w", err)
	}
	for _, e := range entries {
		spriteFileName := e.Name()
		sprite, err1 := initSprite(spriteFileName)
		if err1 != nil {
			return fmt.Errorf("failed to load sprite : %w", err1)
		}
		name := strings.TrimSuffix(spriteFileName, filepath.Ext(spriteFileName))
		g.sprites[name] = sprite
	}

	// init maps
	g.maps = map[string]*Map{}
	entries, err = os.ReadDir(mapPath)
	if err != nil {
		return fmt.Errorf("failed to read map dir : %w", err)
	}
	for _, e := range entries {
		mapFileName := e.Name()
		m, err1 := initMap(mapFileName, g)
		if err1 != nil {
			return fmt.Errorf("failed to load map : %w", err1)
		}
		mapName := strings.TrimSuffix(mapFileName, filepath.Ext(mapFileName))
		g.maps[mapName] = m
	}

	return nil
}
