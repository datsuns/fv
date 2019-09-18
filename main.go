package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type MyMainWindow struct {
	*walk.MainWindow
}

func NewMyMainWindow() *MyMainWindow {
	return &MyMainWindow{}
}

func (mw *MyMainWindow) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_KEYDOWN:
		fmt.Println("WM_KEYDOWN")
		return mw.MainWindow.WndProc(hwnd, msg, wParam, lParam)
	}
	return mw.MainWindow.WndProc(hwnd, msg, wParam, lParam)
}

func main() {
	mw := NewMyMainWindow()

	MainWindow{
		AssignTo: &mw.MainWindow,
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

	NewMyListBox(mw.MainWindow)
	NewMyListBox(mw.MainWindow)

	//mw.MustRegisterProperty("ItemCount", walk.NewReadOnlyProperty(
	//	func() interface{} {
	//		return model.ItemCount()
	//	},
	//	itemCountChangedPublisher.Event(),
	//))
	mw.MainWindow.Run()
}
