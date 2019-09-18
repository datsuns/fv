// +build windows

package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyListBox struct {
	*walk.ListBox
	model   *MyModel
	name    string
	key     walk.Key
	neigbor *MyListBox
}

var (
	keyMap map[walk.Key]walk.Key = map[walk.Key]walk.Key{
		walk.KeyJ: walk.KeyDown,
		walk.KeyK: walk.KeyUp,
	}
)

func NewMyListBox(parent walk.Container, name string) (*MyListBox, error) {
	lb, err := walk.NewListBox(parent)
	if err != nil {
		return nil, err
	}
	mlb := &MyListBox{lb, NewMyModel(), name, 0, nil}
	lb.SetModel(mlb.model)

	if err := walk.InitWrapperWindow(mlb); err != nil {
		return nil, err
	}

	return mlb, nil
}

func (mlb *MyListBox) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_KEYDOWN:
		fmt.Printf("[%s] %v\n", mlb.name, walk.Key(wParam))
		if walk.Key(wParam) == mlb.key {
			fmt.Printf("[%s] change forcus\n", mlb.name)
			mlb.neigbor.SetFocus()
			return 0
		} else if mapped, exists := keyMap[walk.Key(wParam)]; exists {
			return mlb.ListBox.WidgetBase.WndProc(hwnd, msg, uintptr(mapped), lParam)
		}
	}
	return mlb.ListBox.WndProc(hwnd, msg, wParam, lParam)
}

func (mlb *MyListBox) RegisterNeigbor(neigbor *MyListBox, key walk.Key) {
	mlb.key = key
	mlb.neigbor = neigbor
}

type MyModel struct {
	walk.ListModelBase
	path *Path
}

func NewMyModel() *MyModel {
	m := &MyModel{}
	m.path = NewPath("/")
	m.path.Load()
	return m
}

func (m *MyModel) ItemCount() int {
	return m.path.Len()
}

func (m *MyModel) Value(index int) interface{} {
	return m.path.At(index)
}
