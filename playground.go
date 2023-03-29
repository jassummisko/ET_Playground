package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxZLevels = 3
)

type Playground struct {
	entities []Object
	held     Object
	zOffset  int
}

func NewPlayground() *Playground {
	return &Playground{entities: []Object{}, zOffset: 0}
}

func (p *Playground) AddObject(e Object) {
	p.entities = append(p.entities, e)
}

func (p Playground) GetTopEntityOfAny() (int, Object) {
	skip := p.zOffset
	for z := 0; z < maxZLevels; z++ {
		for i, e := range p.entities {
			if e.IsMousedOver() && p.held != e && e.GetZLevel() == z {
				if skip == 0 || z == maxZLevels-1 {
					return i, e
				} else {
					skip--
				}
			}
		}
	}
	return -1, nil
}

func (p Playground) GetTopEntityIgnoringLower() (int, Object) {
	skip := p.zOffset
	max := int(math.Max(float64(skip), float64(p.held.GetZLevel()+1)))
	for z := 0; z < max; z++ {
		for i, e := range p.entities {
			if e.IsMousedOver() && p.held != e && e.GetZLevel() == z {
				if skip == 0 {
					return i, e
				} else {
					skip--
				}
			}
		}
	}
	return -1, nil
}

func (p Playground) GetEntityAtMousePos() (int, Object) {
	if p.held == nil {
		return p.GetTopEntityOfAny()
	} else {
		return p.GetTopEntityIgnoringLower()
	}
}

func (p *Playground) Clean() {
	for i, e := range p.entities {
		if e.IsToDelete() {
			p.entities = removeObject(p.entities, i)
		}
	}
}

func (p *Playground) Update() {
	p.DoMouse()
	for _, e := range p.entities {
		e.Update()
	}
	//                               v THIS IS STUPID. FIX PROBLEM ID 1 BETTER
	if rl.IsKeyDown(rl.KeyLeftShift) && p.held == nil {
		p.zOffset = 1
	} else {
		p.zOffset = 0
	}

	p.Clean()
}

func (p *Playground) MoveObjectToTop(i int) {
	p.entities = moveObjectToTop(p.entities, i)
}

func (p *Playground) DoMouse() {
	i, mousedOverObject := p.GetEntityAtMousePos()
	if mousedOverObject != nil {
		if rl.IsMouseButtonPressed(0) {
			p.held = mousedOverObject
			p.held.SetIsHeld(true)
			p.MoveObjectToTop(i)
		}
	}

	if p.held != nil {
		pos := p.held.GetPos()
		mouseDelta := rl.GetMouseDelta()
		p.held.SetPos(pos.X+mouseDelta.X, pos.Y+mouseDelta.Y)
	}

	if rl.IsMouseButtonReleased(0) {
		_, obj := p.GetEntityAtMousePos()
		if obj != nil {
			p.held.DropInto(obj)
		}
		if p.held != nil {
			p.held.SetIsHeld(false)
			p.held = nil
		}
	}
}

func (p *Playground) Draw() {
	for _, e := range p.entities {
		e.Draw()

		_, mousedOver := p.GetEntityAtMousePos()
		if mousedOver != nil {
			mousedOver.DrawMouseBox()
		}
	}
}
