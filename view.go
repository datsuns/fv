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

func NewMyListBox(parent walk.Container, name string, root string) (*MyListBox, error) {
	lb, err := walk.NewListBox(parent)
	if err != nil {
		return nil, err
	}
	mlb := &MyListBox{lb, NewMyModel(root), name, 0, nil}
	lb.SetModel(mlb.model)

	if err := walk.InitWrapperWindow(mlb); err != nil {
		return nil, err
	}

	return mlb, nil
}

func (mlb *MyListBox) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	//fmt.Printf("[%v] [%v]\n", mlb.name, Msg2Str(msg))
	switch msg {
	case win.WM_KEYDOWN:
		if walk.Key(wParam) == mlb.key {
			fmt.Printf("[%s] change forcus\n", mlb.name)
			mlb.neigbor.SetFocus()
			return 0
		} else if walk.Key(wParam) == walk.KeyReturn {
			idx := mlb.ListBox.CurrentIndex()
			v, _ := mlb.model.Value(idx).(string)
			fmt.Printf("[%s] return w/ [%v:%v]\n", mlb.name, idx, v)
			mlb.ListBox.WndProc(hwnd, msg, wParam, lParam)
			mlb.model.Enter(v)
			mlb.ListBox.SetModel(mlb.model)
		} else if mapped, exists := keyMap[walk.Key(wParam)]; exists {
			return mlb.ListBox.WidgetBase.WndProc(hwnd, msg, uintptr(mapped), lParam)
		}
	case win.WM_CHAR:
		return 0
	}
	return mlb.ListBox.WndProc(hwnd, msg, wParam, lParam)
}

func (mlb *MyListBox) RegisterNeigbor(neigbor *MyListBox, key walk.Key) {
	mlb.key = key
	mlb.neigbor = neigbor
}

type MyModel struct {
	walk.ListModelBase
	Path *Path
}

func NewMyModel(root string) *MyModel {
	m := &MyModel{}
	m.Path = NewPath(root)
	m.Path.Load()
	return m
}

func (m *MyModel) ItemCount() int {
	return m.Path.Len()
}

func (m *MyModel) Value(index int) interface{} {
	return m.Path.At(index)
}

func (m *MyModel) Enter(next string) {
	m.Path.Enter(next)
}
