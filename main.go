package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridWidth    = 32
	gridHeight   = 24
)

func main() {
	width, height := 32, 24
	numAgents := 30
	stepDelay := 300 * time.Millisecond
	env := NewEnvironment(width, height, numAgents)

	game := &Game{
		Env:       env,
		StepDelay: stepDelay,
		lastStep:  time.Now(),
		gridImage: createGridImage(),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Simulation of cooperation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
