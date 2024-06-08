package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"meows-adventure/internal"
)

func main() {
	ebiten.SetWindowSize(internal.ScreenWidth, internal.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := &internal.Game{}
	err := game.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
