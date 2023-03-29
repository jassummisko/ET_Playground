package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Segment struct {
	Entity
	elements []*Element
}

func NewSegment(pos rl.Vector2, elements []*Element) *Segment {
	return &Segment{
		Entity:   *NewEntity(pos),
		elements: elements,
	}
}

func (seg Segment) GetElements() []*Element {
	return seg.elements
}

func (seg *Segment) AddElement(element *Element) {
	seg.elements = append(seg.elements, element)
}

func (seg *Segment) Update() {
	for i, element := range seg.elements {
		element.SetAccessible(false)
		element.SetPos(seg.Entity.pos.X+float32(i)*30+6, seg.Entity.pos.Y)
	}

	seg.Entity.Update()
	seg.Entity.outline.Width = float32(len(seg.elements)*30 + 8)

}

func (seg *Segment) Draw() {
	for _, element := range seg.elements {
		element.Draw()
	}

	rl.DrawText(
		"|",
		int32(seg.pos.X),
		int32(seg.pos.Y),
		32, rl.GetColor(0x111111ff))

	rl.DrawText(
		"|",
		int32(seg.pos.X+float32(len(seg.elements)*30)),
		int32(seg.pos.Y),
		32, rl.GetColor(0x111111ff))
}

func (seg *Segment) DropInto(o Object) {}
