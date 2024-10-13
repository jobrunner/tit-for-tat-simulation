package main

import (
	"math/rand"

	"github.com/google/uuid"
)

type Environment struct {
	Simulation    string
	Width, Height int
	Step          int
	Agents        []*Agent
}

type TotalState struct {
	Simulation  string         `json:"simulation"`
	Strategies  map[string]int `json:"strategies"`
	Step        int            `json:"step"`
	Points      map[string]int `json:"points"`
	TotalPoints int            `json:"points_total"`
}

func NewEnvironment(width, height, numAgents int) *Environment {
	env := &Environment{
		Simulation: uuid.New().String(),
		Width:      width,
		Height:     height,
		Agents:     make([]*Agent, numAgents),
		Step:       0,
	}

	for i := 0; i < numAgents; i++ {
		agent := &Agent{
			ID:         i,
			X:          rand.Intn(width),
			Y:          rand.Intn(height),
			Points:     0,
			Strategy:   InitializeStrategy(),
			LastChoice: "",
		}
		env.Agents[i] = agent
	}

	return env
}

func (env *Environment) TotalState() (TotalState, error) {
	totalPoints := 0
	points := make(map[string]int)
	strategies := make(map[string]int)

	for _, agent := range env.Agents {
		totalPoints += agent.Points
		points[agent.Strategy.Kind()] += agent.Points
		strategies[agent.Strategy.Kind()] += 1
	}
	entry := TotalState{
		Simulation:  env.Simulation,
		Step:        env.Step,
		TotalPoints: totalPoints,
		Points:      points,
		Strategies:  strategies,
	}

	return entry, nil
}
