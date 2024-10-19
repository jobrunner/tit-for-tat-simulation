package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config := LoadConfig()

	env := NewEnvironment(&config)
	game := NewGame(env)

	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Simulation of cooperation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
