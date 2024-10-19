package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Env       *Environment
	StepDelay time.Duration
	lastStep  time.Time
	gridImage *ebiten.Image
}

func (g *Game) CheckCollisions() {
	for i := 0; i < len(g.Env.Agents); i++ {
		for j := i + 1; j < len(g.Env.Agents); j++ {
			if g.Env.Agents[i].X == g.Env.Agents[j].X && g.Env.Agents[i].Y == g.Env.Agents[j].Y {
				g.Env.Agents[i].PlayTitForTat(g.Env.Agents[j])
			}
		}
	}
}

func createGridImage() *ebiten.Image {
	gridImage := ebiten.NewImage(screenWidth, screenHeight)

	cellWidth := float32(screenWidth) / float32(gridWidth)
	cellHeight := float32(screenHeight) / float32(gridHeight)

	path := vector.Path{}

	// Vertikale Linien
	for x := 0; x <= gridWidth; x++ {
		lx := float32(x) * cellWidth
		path.MoveTo(lx, 0)
		path.LineTo(lx, float32(screenHeight))
	}

	// Horizontale Linien
	for y := 0; y <= gridHeight; y++ {
		ly := float32(y) * cellHeight
		path.MoveTo(0, ly)
		path.LineTo(float32(screenWidth), ly)
	}

	op := &vector.StrokeOptions{}
	op.Width = 1
	op.LineJoin = vector.LineJoinMiter
	op.LineCap = vector.LineCapButt

	// Append vertices and indices for the strokes
	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, op)

	// Update the vertex colors (grey color for the lines)
	for i := range vs {
		vs[i].ColorR = 0.4 // Corresponding to red channel of the color.RGBA{100, 100, 100, 255}
		vs[i].ColorG = 0.4 // Corresponding to green channel
		vs[i].ColorB = 0.4 // Corresponding to blue channel
		vs[i].ColorA = 1.0 // Full opacity
	}

	// Create a 1x1 white texture (as a dummy texture)
	solidWhiteImage := ebiten.NewImage(1, 1)
	solidWhiteImage.Fill(color.White)

	// Draw the lines using the gridImage
	gridImage.DrawTriangles(vs, is, solidWhiteImage, nil)

	return gridImage
}

func (g *Game) Update() error {
	g.Env.Step++
	// Delay for the simulation step
	if time.Since(g.lastStep) < g.StepDelay {
		return nil
	}

	for _, agent := range g.Env.Agents {
		agent.Move(g.Env)
	}

	g.CheckCollisions()
	g.lastStep = time.Now()

	logState(g.Env)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.gridImage, op)

	for _, agent := range g.Env.Agents {
		agentColor := agent.GetColor()
		agentImg := ebiten.NewImage(20, 20)
		agentImg.Fill(agentColor)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(agent.X*20), float64(agent.Y*20))
		screen.DrawImage(agentImg, op)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("S: %s (P: %d)", agent.Strategy.Kind(), agent.Points), agent.X*20, agent.Y*20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Window size in px
	return screenWidth, screenHeight
}
