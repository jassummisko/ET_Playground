package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxZLevels = 3
)

type State int

const (
	Main State = iota
	Menu
)

type Playground struct {
	entities []Object
	held     Object
	zOffset  int
}

func NewPlayground() *Playground {
	return &Playground{entities: []Object{}, zOffset: 0}
}

func (p *Playground) Update() {
	p.MouseInput()
	p.KeyboardInput()
	p.UpdateObjects()
	p.CleanObjects()
}

func (p *Playground) Draw() {
	p.DrawObjects()
	p.DrawMouseBox()
}

func (p *Playground) MouseInput() {
	p.HandleClick(p.GetEntityAtMousePos())
	p.HandleMouseDrag()
	p.HandleClickRelease()
}

func (p *Playground) KeyboardInput() {
	// TODO: THIS IS A BODGE. I PROBABLY WANT TO IMPLEMENT A MENU.
	keyPressed := rl.GetKeyPressed()
	if KeyIsElement(keyPressed) {
		p.AddObject(NewElement(
			rl.GetMousePosition(),
			string(keyPressed),
			rl.IsKeyDown(rl.KeyLeftShift),
		),
		)
	}
	//                               v TODO: THIS IS STUPID. FIX PROBLEM ID 1 BETTER
	if rl.IsKeyDown(rl.KeyLeftShift) && p.held == nil {
		p.zOffset = 1
	} else {
		p.zOffset = 0
	}
}

func (p *Playground) AddObject(o Object) {
	p.entities = append(p.entities, o)
}

func (p *Playground) UpdateObjects() {
	for _, e := range p.entities {
		e.Update()
	}
}

func (p *Playground) DrawObjects() {
	for _, e := range p.entities {
		e.Draw()
	}
}

func (p *Playground) DrawMouseBox() {
	_, mousedOver := p.GetEntityAtMousePos()
	if mousedOver != nil {
		mousedOver.DrawMouseBox()
	}
}

func (p Playground) GetTopEntity() (int, Object) {
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
		return p.GetTopEntity()
	} else {
		return p.GetTopEntityIgnoringLower()
	}
}

func (p *Playground) CleanObjects() {
	for i, e := range p.entities {
		if e.IsToDelete() {
			p.entities = removeObject(p.entities, i)
		}
	}
}

func (p *Playground) MoveObjectToTop(i int) {
	p.entities = moveObjectToTop(p.entities, i)
}

func (p *Playground) PickUpObject(o Object) {
	p.held = o
	p.held.SetIsHeld(true)
}

func (p *Playground) LetGoOfHeldObject() {
	if p.held != nil {
		p.held.SetIsHeld(false)
		p.held = nil
	}
}

func (p *Playground) HandleMouseDrag() {
	if p.held != nil {
		pos := p.held.GetPos()
		mouseDelta := rl.GetMouseDelta()
		p.held.SetPos(pos.X+mouseDelta.X, pos.Y+mouseDelta.Y)
	}
}

func (p *Playground) HandleClick(i int, mousedOverObject Object) {
	if mousedOverObject != nil {
		if rl.IsMouseButtonPressed(0) {
			if rl.IsKeyDown(rl.KeyLeftControl) {
				mousedOverObject.AltAction()
			} else {
				p.PickUpObject(mousedOverObject)
				p.MoveObjectToTop(i)
			}
		}

		if rl.IsMouseButtonPressed(1) {
			mousedOverObject.MarkToDelete()
		}
	}
}

func (p *Playground) HandleClickRelease() {
	if rl.IsMouseButtonReleased(0) {
		if p.held != nil {
			_, obj := p.GetEntityAtMousePos()
			if obj != nil {
				p.held.DropInto(obj)
			}
			p.LetGoOfHeldObject()
		}
	}
}