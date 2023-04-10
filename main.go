package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 600
)

var g_palette []rl.Color = []rl.Color{
	rl.GetColor(0x111111ff), // Black
	rl.GetColor(0xddddddff), // White
}

var g_playground *Playground

func initPlayground() {
	g_playground = NewPlayground()
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Element Theory Playground")
	rl.SetTargetFPS(60)

	initPlayground()

	for !rl.WindowShouldClose() {
		g_playground.Update()
		rl.BeginDrawing()
		rl.ClearBackground(g_palette[1])
		g_playground.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
