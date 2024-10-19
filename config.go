package main

import (
	"flag"
	"os"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridWidth    = 32
	gridHeight   = 24
	stepDelayMs  = 300
	numAgents    = 30
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	GridWidth    int
	GridHeight   int
	StepDelayMs  int
	NumAgents    int
	NoGrid       bool
	NoLogs       bool
	HeadlessMode bool
}

func LoadConfig() Config {
	var cfg Config

	flag.IntVar(&cfg.ScreenWidth, "screen-width", screenWidth, "Screen width of the simulation window")
	flag.IntVar(&cfg.ScreenHeight, "screen-height", screenHeight, "Screen height of the simulation window")
	flag.IntVar(&cfg.GridWidth, "grid-width", gridWidth, "Grid width of the simulation environment")
	flag.IntVar(&cfg.GridHeight, "grid-height", gridHeight, "Grid height of the simulation environment")
	flag.IntVar(&cfg.StepDelayMs, "step-delay", stepDelayMs, "Step delay between simulation steps in ms")
	flag.IntVar(&cfg.NumAgents, "agents", numAgents, "Number of agents in the simulation")
	flag.BoolVar(&cfg.NoGrid, "no-grid", false, "Should the drawing of a grid be omitted?")
	flag.BoolVar(&cfg.NoLogs, "no-logs", false, "Should the output of the log be suppressed?")
	flag.BoolVar(&cfg.HeadlessMode, "headless", false, "Should the simulation run in headless mode?")

	showHelp := flag.Bool("h", false, "Display help")
	flag.BoolVar(showHelp, "help", false, "Display help")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	return cfg
}
