package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 600
)

var g_playground *Playground

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Element Theory Playground")
	rl.SetTargetFPS(60)

	ents := []*Element{
		NewElement(rl.NewVector2(100, 100), "L"),
		NewElement(rl.NewVector2(100, 100), "A"),
		NewElement(rl.NewVector2(100, 100), "U"),
	}

	g_playground = NewPlayground()

	for _, ent := range ents {
		g_playground.AddObject(ent)
	}
	testSeg := NewSegment(rl.NewVector2(100, 100), ents)

	g_playground.AddObject(testSeg)
	g_playground.AddObject(NewElement(rl.NewVector2(300, 300), "H"))
	g_playground.AddObject(NewElement(rl.NewVector2(400, 300), "L"))

	for !rl.WindowShouldClose() {
		g_playground.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0xddddddff))
		g_playground.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
