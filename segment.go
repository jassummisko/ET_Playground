package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Segment struct {
	Entity
	elements []*Element
}

func NewSegment(pos rl.Vector2, elements []*Element) *Segment {
	return &Segment{
		Entity:   *NewEntity(pos, 1),
		elements: elements,
	}
}

func (seg *Segment) Update() {
	seg.UpdateElements()
	seg.Entity.Update()
	seg.UpdateOutline()
}

func (seg *Segment) Draw() {
	for _, element := range seg.elements {
		element.Draw()
	}
	seg.DrawPipes()
}

func (seg *Segment) UpdateElements() {
	for i, element := range seg.elements {
		if element.GetIsHeld() || element.IsToDelete() {
			seg.RemoveElement(i)
		} else {
			element.SetPos(seg.Entity.pos.X+float32(i)*30+6, seg.Entity.pos.Y)
		}
	}
}

func (seg *Segment) UpdateOutline() {
	seg.Entity.outline.Width = float32(seg.GetWidth() + 10)
}

func (seg Segment) GetElements() []*Element {
	return seg.elements
}

func (seg *Segment) AddElement(element *Element) {
	seg.elements = append(seg.elements, element)
}

func (seg *Segment) RemoveElement(index int) {
	seg.elements = append(seg.elements[:index], seg.elements[index+1:]...)
}

func (seg *Segment) HasElement(label string) bool {
	for _, element := range seg.elements {
		if element.label == label {
			return true
		}
	}
	return false
}

func (seg *Segment) DrawPipes() {
	rl.DrawText(
		"|",
		int32(seg.pos.X),
		int32(seg.pos.Y),
		32, g_palette[0])

	rl.DrawText(
		"|",
		int32(seg.pos.X)+seg.GetWidth(),
		int32(seg.pos.Y),
		32, g_palette[0])
}

func (seg Segment) GetWidth() int32 {
	if len(seg.elements) > 0 {
		return int32(len(seg.elements) * 30)
	} else {
		return int32(30)
	}
}

func (seg *Segment) DropInto(o Object) {
	//TODO: CREATE PROSODIC STRUCTURE
}

func (seg *Segment) AltAction() {
	seg.MarkToDelete()
}
