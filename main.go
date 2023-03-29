package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Element Theory Playground")
	rl.SetTargetFPS(60)

	ents := []*Element{
		NewElement(rl.NewVector2(100, 100), "L"),
		NewElement(rl.NewVector2(100, 100), "A"),
		NewElement(rl.NewVector2(100, 100), "U"),
	}

	testPlayground := NewPlayground()

	for _, ent := range ents {
		testPlayground.AddObject(ent)
	}
	testSeg := NewSegment(rl.NewVector2(100, 100), ents)

	testPlayground.AddObject(testSeg)
	testPlayground.AddObject(NewElement(rl.NewVector2(300, 300), "H"))
	testPlayground.AddObject(NewElement(rl.NewVector2(400, 300), "L"))

	for !rl.WindowShouldClose() {
		testPlayground.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0xddddddff))
		testPlayground.Draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
