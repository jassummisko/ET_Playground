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

type Collider interface {
	GetColBox() rl.Rectangle
	GetZLevel() int
}

type Actor interface {
	AltAction()
	DropInto(Object)
}

type Interacter interface {
	Mouser
	Mover
	Collider
	Actor
}

type Deleter interface {
	IsToDelete() bool
	MarkToDelete()
}

type Updater interface {
	Update()
	Draw()
}

type Object interface {
	Interacter
	Deleter
	Updater
}
