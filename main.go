package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//"github.com/lxn/win"
	"os"
	"strings"
)

func main() {
	var mw *walk.MainWindow
	var inTE, outTE *walk.TextEdit
	var lb *walk.ListBox
	var model *EnvModel = NewEnvModel()
	itemCountChangedPublisher := walk.EventPublisher{}

	MainWindow{
		AssignTo: &mw,
		Title:    "SCREAMO",
		MinSize:  Size{600, 400},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &inTE,
						OnKeyDown: func(key walk.Key) {
							if key == walk.KeyReturn {
								fmt.Println("Return")
							}
						},
					},
					TextEdit{
						AssignTo: &outTE,
						ReadOnly: true,
					},
					ListBox{
						AssignTo: &lb,
						Model:    model,
						//OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,
						//OnItemActivated:       mw.lb_ItemActivated,
						OnKeyDown: func(key walk.Key) {
							if key == walk.KeyRight {
								fmt.Println("index:%v", lb.CurrentIndex())
							}
							if key == walk.KeyF1 {
								fmt.Println("down")
								now := lb.CurrentIndex()
								lb.SetCurrentIndex(now - 1)
								fmt.Printf("now[%v] set to [%v]\n", now, lb.CurrentIndex())
								lb.CurrentIndexChanged()
								//lb.SendMessage(win.LB_SETCURSEL, uintptr(lb.CurrentIndex()), 0)
							}
							if key == walk.KeyK {
								fmt.Println("K")
							}
						},
					},
				},
			},
		},
	}.Create()
	mw.MustRegisterProperty("ItemCount", walk.NewReadOnlyProperty(
		func() interface{} {
			return model.ItemCount()
		},
		itemCountChangedPublisher.Event(),
	))
	mw.Run()
}

//tv.MustRegisterProperty("ItemCount", NewReadOnlyProperty(
//	func() interface{} {
//		if tv.model == nil {
//			return 0
//		}
//		return tv.model.RowCount()
//	},
//	tv.itemCountChangedPublisher.Event()))

type EnvItem struct {
	name  string
	value string
}

type EnvModel struct {
	walk.ListModelBase
	items []EnvItem
}

func NewEnvModel() *EnvModel {
	env := os.Environ()

	m := &EnvModel{items: make([]EnvItem, len(env))}

	for i, e := range env {
		j := strings.Index(e, "=")
		if j == 0 {
			continue
		}

		name := e[0:j]
		value := strings.Replace(e[j+1:], ";", "\r\n", -1)

		m.items[i] = EnvItem{name, value}
	}

	return m
}

func (m *EnvModel) ItemCount() int {
	return len(m.items)
}

func (m *EnvModel) Value(index int) interface{} {
	return m.items[index].name
}
