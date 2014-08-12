package main

import (
	"github.com/go-gl/gl"
)

type Drawable interface {
	draw()
	lighting() bool
}

type DisplayList struct {
	id       uint
	lighting bool
}

func (list *DisplayList) draw() {
	list.setLighting()

	gl.CallList(list.id)
}

func (list *DisplayList) setLighting() {
	if list.lighting {
		gl.Enable(gl.LIGHTING)
	} else {
		gl.Disable(gl.LIGHTING)
	}
}

type DisplayListManager struct {
	lists map[uint]DisplayList
}

func NewDisplayListManager() *DisplayListManager {
	return &DisplayListManager{
		lists: make(map[uint]DisplayList),
	}
}

func (manager *DisplayListManager) draw() {
	for _, list := range manager.lists {
		list.draw()
	}
}

func (manager *DisplayListManager) add(drawable Drawable) uint {
	id := gl.GenLists(1)
	gl.NewList(id, gl.COMPILE)
	drawable.draw()
	gl.EndList()

	manager.lists[id] = DisplayList{id, drawable.lighting()}

	return id
}

func (manager *DisplayListManager) remove(id uint) {
	gl.DeleteLists(id, 1)

	delete(manager.lists, id)
}
