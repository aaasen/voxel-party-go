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
	lists []DisplayList
}

func NewDisplayListManager() *DisplayListManager {
	return &DisplayListManager{
		lists: make([]DisplayList, 0, 10),
	}
}

func (manager *DisplayListManager) draw() {
	for _, list := range manager.lists {
		list.draw()
	}
}

func (manager *DisplayListManager) add(drawable Drawable) {
	id := gl.GenLists(1)
	gl.NewList(id, gl.COMPILE)
	drawable.draw()
	gl.EndList()

	manager.lists = append(manager.lists, DisplayList{id, drawable.lighting()})
}
