package internal

import (
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Map struct {
	game *Game
	name string
	info *mapInfo
}

type mapInfo struct {
	Width       int           `json:"width"`
	Height      int           `json:"height"`
	BlockWidth  int           `json:"block_width"`
	BlockHeight int           `json:"block_height"`
	Sprites     []interface{} `json:"sprites"`
	Data        []int         `json:"data"`
}

func initMap(name string, g *Game) (*Map, error) {
	path := filepath.Join(mapPath, name)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load map %q: %w", path, err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read map file  %q: %w", path, err)
	}
	info := &mapInfo{}
	err = json.Unmarshal(data, info)
	if err != nil {
		return nil, fmt.Errorf("failed to load map  %q json info: %w", path, err)
	}
	m := &Map{
		game: g,
		name: name,
		info: info,
	}
	return m, nil
}

func (m *Map) Draw(screen *ebiten.Image) error {
	for i, n := range m.info.Data {
		if n == 0 {
			continue
		}
		si := m.info.Sprites[n].(string)
		spInfo := parseSpriteInfo(si)
		spInfo.width = m.info.BlockWidth
		spInfo.height = m.info.BlockHeight
		s, err := getSprite(spInfo, m.game.sprites)
		if err != nil {
			return fmt.Errorf("failed to get sprite: %w", err)
		}
		x0 := i % m.info.Width * m.info.BlockWidth
		y0 := (m.info.Height - i/m.info.Width - 1) * m.info.BlockHeight
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(x0), float64(y0)-128)
		screen.DrawImage(s, opt)
	}
	return nil
}

type spriteInfo struct {
	name          string
	x, y          int
	width, height int
}

func parseSpriteInfo(s string) *spriteInfo {
	sub := strings.Split(s, ":")
	name := sub[0]
	var x, y int
	if len(sub) > 1 {
		coordinate := strings.Split(sub[1], ",")
		x, _ = strconv.Atoi(coordinate[0])
		y, _ = strconv.Atoi(coordinate[1])
	}

	return &spriteInfo{
		name: name,
		x:    x,
		y:    y,
	}
}

func getSprite(info *spriteInfo, sprites map[string]*Sprite) (*ebiten.Image, error) {
	s, ok := sprites[info.name]
	if !ok {
		return nil, fmt.Errorf("sprite %s not exists", info.name)
	}
	x := info.x * info.width
	y := info.y * info.height
	return s.img.SubImage(image.Rect(x, y, x+info.width, y+info.height)).(*ebiten.Image), nil
}
