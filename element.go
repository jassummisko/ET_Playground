package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Element struct {
	Entity
	label string
}

func NewElement(pos rl.Vector2, label string) *Element {
	return &Element{
		*NewEntity(pos, 2),
		label,
	}
}

func (e *Element) Update() {
	e.Entity.Update()
}

func (e Element) Draw() {
	rl.DrawText(
		e.label,
		int32(e.pos.X),
		int32(e.pos.Y),
		32, rl.GetColor(0x111111ff))
}

func (e *Element) DropInto(o Object) {
	switch v := o.(type) {
	case *Segment:
		v.elements = append(v.elements, e)
	}
}
