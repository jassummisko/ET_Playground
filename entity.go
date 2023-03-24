package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	pos          rl.Vector2
	outline      rl.Rectangle
	isMousedOver bool
	selected     bool
	spacing      float32
}

func NewEntity(pos rl.Vector2) *Entity {
	var (
		spacing       float32 = 4
		outlineWidth  float32 = 24
		outlineHeight float32 = 32
	)

	return &Entity{
		pos:     pos,
		spacing: 4,
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
		rl.Black,
	)

}

func (e *Entity) SetSelected(selected bool) {
	e.selected = selected
}

func (e Entity) GetSelected() bool {
	return e.selected
}

func (e *Entity) SetMousedOver(isMousedOver bool) {
	e.isMousedOver = isMousedOver
}

func (e *Entity) IsMousedOver() bool {
	return rl.CheckCollisionPointRec(rl.GetMousePosition(), e.GetColBox())
}

func (e *Entity) SetPos(x float32, y float32) {
	e.pos.X = x
	e.pos.Y = y
}

func (e *Entity) GetPos() rl.Vector2 {
	return e.pos
}
