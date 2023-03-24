package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mouser interface {
	SetSelected(bool)
	GetSelected() bool
	IsMousedOver() bool
	DrawMouseBox()
}

type Mover interface {
	SetPos(x float32, y float32)
	GetPos() rl.Vector2
}

type Collider interface {
	GetColBox() rl.Rectangle
}

type Object interface {
	Mouser
	Mover
	Collider
	Update()
	Draw()
}
