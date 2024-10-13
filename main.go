package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Simulation of cooperation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
