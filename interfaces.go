package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mouser interface {
	IsMousedOver() bool
	DrawMouseBox()
	GetIsHeld() bool
	SetIsHeld(bool)
}

type Mover interface {
	SetPos(x float32, y float32)
	GetPos() rl.Vector2
}

type Interactable interface {
	AltAction()
	DropInto(Object)
}

type Collider interface {
	GetColBox() rl.Rectangle
	GetZLevel() int
}

type Object interface {
	Mouser
	Mover
	Collider
	Interactable
	IsToDelete() bool
	MarkToDelete()
	Update()
	Draw()
}
