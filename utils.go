package main

import rl "github.com/gen2brain/raylib-go/raylib"

func removeObject(slice []Object, s int) []Object {
	return append(slice[:s], slice[s+1:]...)
}

func prependObject(x []Object, y Object) []Object {
	return append([]Object{y}, x...)
}

func moveObjectToTop(slice []Object, s int) []Object {
	obj := slice[s]
	slice = prependObject(slice, obj)
	slice = removeObject(slice, s+1)
	return slice
}

func KeyIsElement(key int32) bool {
	elementLabels := []int32{
		rl.KeyA,
		rl.KeyU,
		rl.KeyI,
		rl.KeyS,
		rl.KeyH,
		rl.KeyL,
	}

	for _, label := range elementLabels {
		if key == label {
			return true
		}
	}
	return false
}
