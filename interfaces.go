package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mouser interface {
	IsMousedOver() bool
	DrawMouseBox()
	GetAccessible() bool
	SetAccessible(bool)
}

type Mover interface {
	SetPos(x float32, y float32)
	GetPos() rl.Vector2
}

type Collider interface {
	GetColBox() rl.Rectangle
	DropInto(Object)
}

type Object interface {
	Mouser
	Mover
	Collider
	Update()
	Draw()
}
