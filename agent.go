package main

import (
	"image/color"
	"math/rand"
)

var (
	colorCooperate = color.RGBA{0, 255, 0, 255} // Green for Always Cooperate
	colorDefect    = color.RGBA{255, 0, 0, 255} // Red for Always Defect
	colorTitForTat = color.RGBA{0, 0, 255, 255} // Blue for Tit-for-Tat
)

type Agent struct {
	ID         int
	X, Y       int
	Points     int
	Strategy   Decider
	LastChoice string
}

func (a *Agent) Move(env *Environment) {
	newX := a.X + rand.Intn(3) - 1 // Movement in X direction (-1, 0, 1)
	newY := a.Y + rand.Intn(3) - 1 // Movement in X direction (-1, 0, 1)

	// Ensure that the new position is in the environment
	if newX >= 0 && newX < env.Width && newY >= 0 && newY < env.Height {
		a.X = newX
		a.Y = newY
	}
}

func (a *Agent) GetColor() color.Color {
	switch a.Strategy.(type) {
	case *AlwaysCooperate:
		return colorCooperate
	case *AlwaysDefect:
		return colorDefect
	case *TitForTat:
		return colorTitForTat
	default:
		return color.White
	}
}

func InitializeStrategy() Decider {
	var strategy Decider
	switch rand.Intn(3) {
	case 0:
		strategy = &AlwaysCooperate{}
	case 1:
		strategy = &AlwaysDefect{}
	case 2:
		strategy = &TitForTat{}
	}
	return strategy
}

func (a *Agent) PlayTitForTat(opponent *Agent) {
	choice := a.Strategy.Decide(opponent.LastChoice)

	if choice == "C" && opponent.LastChoice == "C" {
		a.Points += 3
		opponent.Points += 3
	} else if choice == "C" && opponent.LastChoice == "D" {
		opponent.Points += 5
		a.Points += 0
	} else if choice == "D" && opponent.LastChoice == "C" {
		a.Points += 5
		opponent.Points += 0
	} else if choice == "D" && opponent.LastChoice == "D" {
		a.Points += 1
		opponent.Points += 1
	}

	// Saves the decision for the next round (for tit-for-tat strategies)
	a.LastChoice = choice
}
