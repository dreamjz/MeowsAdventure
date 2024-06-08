package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"meows-adventure/internal"
)

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")

	game := &internal.Game{}
	err := game.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
