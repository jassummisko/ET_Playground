package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Playground struct {
	entities []Object
	held     Object
}

func NewPlayground() *Playground {
	return &Playground{entities: []Object{}}
}

func (p *Playground) AddObject(e Object) {
	p.entities = append(p.entities, e)
}

func (p Playground) GetEntitiesAtMousePosition() []Object {
	var entities []Object
	for _, e := range p.entities {
		if e.IsMousedOver() && e.GetAccessible() {
			entities = append(entities, e)
		}
	}
	return entities
}

func (p Playground) GetTopEntityAtMousePosition() (int, Object) {
	for i, e := range p.entities {
		if e.IsMousedOver() && e.GetAccessible() && p.held != e {
			return i, e
		}
	}
	return -1, nil
}

func (p *Playground) Update() {

	p.DoMouse()
	for _, e := range p.entities {
		e.Update()
	}

}

func (p *Playground) MoveObjectToTop(i int) {
	p.entities = moveObjectToTop(p.entities, i)
}

func (p *Playground) DoMouse() {
	i, mousedOverObject := p.GetTopEntityAtMousePosition()
	if mousedOverObject != nil {
		if rl.IsMouseButtonPressed(0) {
			p.held = mousedOverObject
			p.MoveObjectToTop(i)
		}
	}

	if p.held != nil {
		pos := p.held.GetPos()
		mouseDelta := rl.GetMouseDelta()
		p.held.SetPos(pos.X+mouseDelta.X, pos.Y+mouseDelta.Y)
	}

	if rl.IsMouseButtonReleased(0) {
		_, obj := p.GetTopEntityAtMousePosition()
		p.held.DropInto(obj)
		p.held = nil
	}
}

func (p *Playground) Draw() {
	for _, e := range p.entities {
		e.Draw()

		_, mousedOver := p.GetTopEntityAtMousePosition()
		if mousedOver != nil {
			mousedOver.DrawMouseBox()
		}
	}
}
