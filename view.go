// +build windows

package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"unsafe"
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
	col.SetWidth(300)
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
	switch msg {
	case win.WM_NOTIFY:
		nmh := ((*win.NMHDR)(unsafe.Pointer(lParam)))
		//fmt.Printf("[%v] [%v]\n", mv.name, LvMsg2Str(nmh.Code))
		switch nmh.Code {
		case win.LVN_KEYDOWN:
			info := ((*win.NMTVKEYDOWN)(unsafe.Pointer(lParam)))
			key := walk.Key(info.WVKey)
			fmt.Printf("[%s] keydown [%v][%v]\n", mv.name, key, info)
			if key == mv.key {
				fmt.Printf("[%s] change forcus\n", mv.name)
				mv.neigbor.SetFocus()
				return 0
			} else if mapped, exists := keyMap[key]; exists {
				fmt.Printf("[%s] mapped to [%v]\n", mv.name, mapped)
				info.WVKey = uint16(mapped)
				fmt.Printf("[%s] map [%v]\n", mv.name, info)
				// TODO maybe item-changed method?
				return mv.TableView.WndProc(hwnd, msg, wParam, lParam)
			}
		}

		//case win.WM_KEYDOWN:
		//	if walk.Key(wParam) == mv.key {
		//		fmt.Printf("[%s] change forcus\n", mv.name)
		//		mv.neigbor.SetFocus()
		//		return 0
		//	} else if walk.Key(wParam) == walk.KeyReturn {
		//		//idx := mv.ListBox.CurrentIndex()
		//		//v, _ := mv.model.Value(idx).(string)
		//		//fmt.Printf("[%s] return w/ [%v:%v]\n", mv.name, idx, v)
		//		//mv.ListBox.WndProc(hwnd, msg, wParam, lParam)
		//		//mv.model.Enter(v)
		//		//mv.ListBox.SetModel(mv.model)
		//	} else if mapped, exists := keyMap[walk.Key(wParam)]; exists {
		//		return mv.TableView.WndProc(hwnd, msg, uintptr(mapped), lParam)
		//	}
		//case win.WM_CHAR:
		//	return 0
	}
	return mv.TableView.WndProc(hwnd, msg, wParam, lParam)
}

func (mv *MyView) RegisterNeigbor(neigbor *MyView, key walk.Key) {
	mv.key = key
	mv.neigbor = neigbor
}

type MyModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	item       *Path
}

func NewMyModel(root string) *MyModel {
	m := &MyModel{}
	m.item = NewPath(root)
	m.item.Load()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *MyModel) RowCount() int {
	return m.item.Len()
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *MyModel) Value(row, col int) interface{} {
	return m.item.At(row)
	//item := m.items[row]

	//switch col {
	//case 0:
	//	return item.Index

	//case 1:
	//	return item.Bar

	//case 2:
	//	return item.Baz

	//case 3:
	//	return item.Quux
	//}

	//panic("unexpected col")
}

// Called by the TableView to retrieve if a given row is checked.
func (m *MyModel) Checked(row int) bool {
	//return m.items[row].checked
	return false
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *MyModel) SetChecked(row int, checked bool) error {
	//m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *MyModel) Sort(col int, order walk.SortOrder) error {
	return nil
	//m.sortColumn, m.sortOrder = col, order

	//sort.SliceStable(m.items, func(i, j int) bool {
	//	a, b := m.items[i], m.items[j]

	//	c := func(ls bool) bool {
	//		if m.sortOrder == walk.SortAscending {
	//			return ls
	//		}

	//		return !ls
	//	}

	//	switch m.sortColumn {
	//	case 0:
	//		return c(a.Index < b.Index)

	//	case 1:
	//		return c(a.Bar < b.Bar)

	//	case 2:
	//		return c(a.Baz < b.Baz)

	//	case 3:
	//		return c(a.Quux.Before(b.Quux))
	//	}

	//	panic("unreachable")
	//})

	//return m.SorterBase.Sort(col, order)
}

func (m *MyModel) ResetRows() {
	// Create some random data.
	//m.items = make([]*Foo, rand.Intn(50000))

	//now := time.Now()

	//for i := range m.items {
	//	m.items[i] = &Foo{
	//		Index: i,
	//		Bar:   strings.Repeat("*", rand.Intn(5)+1),
	//		Baz:   rand.Float64() * 1000,
	//		Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
	//	}
	//}

	//// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
