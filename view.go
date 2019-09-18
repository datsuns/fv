// +build windows

package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyView struct {
	*walk.TableView
	model   *MyModel
	name    string
	key     walk.Key
	neigbor *MyView
}

var (
	keyMap map[walk.Key]walk.Key = map[walk.Key]walk.Key{
		walk.KeyJ: walk.KeyDown,
		walk.KeyK: walk.KeyUp,
	}
)

func NewMyView(parent walk.Container, name string, root string) (*MyView, error) {
	t, err := walk.NewTableView(parent)
	if err != nil {
		return nil, err
	}
	col := walk.NewTableViewColumn()
	col.SetName("path")
	t.Columns().Clear()
	t.Columns().Add(col)
	mv := &MyView{t, NewMyModel(root), name, 0, nil}
	t.SetModel(mv.model)

	if err := walk.InitWrapperWindow(mv); err != nil {
		return nil, err
	}

	return mv, nil
}

func (mv *MyView) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	fmt.Printf("[%v] [%v]\n", mv.name, Msg2Str(msg))
	//switch msg {
	//case win.WM_KEYDOWN:
	//	if walk.Key(wParam) == mv.key {
	//		fmt.Printf("[%s] change forcus\n", mv.name)
	//		mv.neigbor.SetFocus()
	//		return 0
	//	} else if walk.Key(wParam) == walk.KeyReturn {
	//		idx := mv.ListBox.CurrentIndex()
	//		v, _ := mv.model.Value(idx).(string)
	//		fmt.Printf("[%s] return w/ [%v:%v]\n", mv.name, idx, v)
	//		mv.ListBox.WndProc(hwnd, msg, wParam, lParam)
	//		mv.model.Enter(v)
	//		mv.ListBox.SetModel(mv.model)
	//	} else if mapped, exists := keyMap[walk.Key(wParam)]; exists {
	//		return mv.ListBox.WidgetBase.WndProc(hwnd, msg, uintptr(mapped), lParam)
	//	}
	//case win.WM_CHAR:
	//	return 0
	//}
	return mv.TableView.WndProc(hwnd, msg, wParam, lParam)
}

func (mv *MyView) RegisterNeigbor(neigbor *MyView, key walk.Key) {
	mv.key = key
	mv.neigbor = neigbor
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
