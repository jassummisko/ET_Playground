package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	pos         rl.Vector2
	outline     rl.Rectangle
	world       *Playground
	toBeRemoved bool
	isHeld      bool
	spacing     float32
	z_level     int
}

func NewEntity(pos rl.Vector2, z int) *Entity {
	var (
		spacing       float32 = 4
		outlineWidth  float32 = 24
		outlineHeight float32 = 36
	)

	return &Entity{
		pos:     pos,
		spacing: 4,
		world:   g_playground,
		z_level: z,
		outline: rl.Rectangle{
			X:      pos.X - spacing,
			Y:      pos.Y - spacing/2,
			Width:  outlineWidth,
			Height: outlineHeight,
		}}
}

func (e *Entity) Update() {
	e.outline.X = e.pos.X - e.spacing
	e.outline.Y = e.pos.Y - e.spacing
}

func (e Entity) GetColBox() rl.Rectangle {
	return e.outline
}

func (e Entity) DrawMouseBox() {

	rl.DrawRectangleLines(
		e.outline.ToInt32().X,
		e.outline.ToInt32().Y,
		e.outline.ToInt32().Width,
		e.outline.ToInt32().Height,
		g_palette[0],
	)

}

func (e *Entity) IsMousedOver() bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), e.GetColBox())
}

func (e *Entity) GetPos() rl.Vector2 {
	return e.pos
}

func (e *Entity) SetPos(x float32, y float32) {
	e.pos.X = x
	e.pos.Y = y
}

func (e *Entity) GetZLevel() int {
	return e.z_level
}

func (e *Entity) GetIsHeld() bool {
	return e.isHeld
}

func (e *Entity) SetIsHeld(isHeld bool) {
	e.isHeld = isHeld
}

func (e *Entity) IsToDelete() bool {
	return e.toBeRemoved
}

func (e *Entity) MarkToDelete() {
	e.toBeRemoved = true
}
