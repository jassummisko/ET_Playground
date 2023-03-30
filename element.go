package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Element struct {
	Entity
	label  string
	headed bool
}

func NewElement(pos rl.Vector2, label string, headed bool) *Element {
	return &Element{
		*NewEntity(pos, 2),
		label,
		headed,
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

	if e.headed {
		rl.DrawRectangle(
			e.outline.ToInt32().X,
			e.outline.ToInt32().Y+e.outline.ToInt32().Height,
			e.outline.ToInt32().Width,
			3,
			rl.GetColor(0x111111ff),
		)
	}
}

func (e *Element) DropInto(o Object) {
	switch v := o.(type) {
	case *Segment:
		if !v.HasElement(e.label) {
			v.elements = append(v.elements, e)
		}
	case *Element:
		if v.label == e.label {
			return
		}
		t := o.(*Element)
		elements := []*Element{t, e}
		seg := NewSegment(v.GetPosOfNewSegment(), elements)
		e.world.AddObject(seg)
	}
}

func (e *Element) AltAction() {
	e.world.AddObject(
		NewSegment(
			e.GetPosOfNewSegment(),
			[]*Element{e},
		),
	)
}

func (e *Element) GetPosOfNewSegment() rl.Vector2 {
	return rl.NewVector2(
		float32(e.pos.X-6),
		float32(e.pos.Y),
	)
}