package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var mw *walk.MainWindow

	fmt.Println("hello")

	MainWindow{
		AssignTo: &mw,
		Title:    "SCREAMO",
		MinSize:  Size{600, 400},
		Layout:   HBox{},
		//Children: []Widget{
		//	HSplitter{
		//		Children: []Widget{
		//			TextEdit{
		//				AssignTo: &inTE,
		//				OnKeyDown: func(key walk.Key) {
		//					if key == walk.KeyReturn {
		//						fmt.Println("Return")
		//					}
		//				},
		//			},
		//			TextEdit{
		//				AssignTo: &outTE,
		//				ReadOnly: true,
		//			},
		//			//ListBox{
		//			//	AssignTo: &lb,
		//			//	Model:    model,
		//			//	//OnCurrentIndexChanged: mw.lb_CurrentIndexChanged,
		//			//	//OnItemActivated:       mw.lb_ItemActivated,
		//			//	OnKeyDown: func(key walk.Key) {
		//			//		if key == walk.KeyRight {
		//			//			fmt.Println("index:%v", lb.CurrentIndex())
		//			//		}
		//			//		if key == walk.KeyF1 {
		//			//			fmt.Println("down")
		//			//			now := lb.CurrentIndex()
		//			//			lb.SetCurrentIndex(now - 1)
		//			//			fmt.Printf("now[%v] set to [%v]\n", now, lb.CurrentIndex())
		//			//			lb.CurrentIndexChanged()
		//			//			//lb.SendMessage(win.LB_SETCURSEL, uintptr(lb.CurrentIndex()), 0)
		//			//		}
		//			//		if key == walk.KeyK {
		//			//			fmt.Println("K")
		//			//		}
		//			//	},
		//			//},
		//		},
		//	},
	}.Create()

	left, err := NewMyListBox(mw, "LEFT")
	if err != nil {
		panic(err)
	}
	right, err := NewMyListBox(mw, "RIGHT")
	if err != nil {
		panic(err)
	}
	left.RegisterNeigbor(right, walk.KeyL)
	right.RegisterNeigbor(left, walk.KeyH)

	//mw.MustRegisterProperty("ItemCount", walk.NewReadOnlyProperty(
	//	func() interface{} {
	//		return model.ItemCount()
	//	},
	//	itemCountChangedPublisher.Event(),
	//))
	mw.Run()
}
