package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Environment struct {
	Simulation    string
	Width, Height int
	Step          int
	Agents        []*Agent
	Config        *Config
}

type TotalState struct {
	Simulation  string         `json:"simulation"`
	Strategies  map[string]int `json:"strategies"`
	Step        int            `json:"step"`
	Points      map[string]int `json:"points"`
	TotalPoints int            `json:"points_total"`
	Timestamp   int64          `json:"ts"`
}

func NewEnvironment(cfg *Config) *Environment {
	width := cfg.GridWidth
	height := cfg.GridHeight
	numAgents := cfg.NumAgents

	env := &Environment{
		Simulation: uuid.New().String(),
		Width:      width,
		Height:     height,
		Agents:     make([]*Agent, numAgents),
		Step:       0,
		Config:     cfg,
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
		Timestamp:   time.Now().UnixNano(),
	}

	return entry, nil
}
