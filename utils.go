package main

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
