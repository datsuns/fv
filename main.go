package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
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

func Msg2Str(msg uint32) string {
	switch msg {
	case win.WM_APP:
		return "WM_APP"
	case win.WM_ACTIVATE:
		return "WM_ACTIVATE"
	case win.WM_ACTIVATEAPP:
		return "WM_ACTIVATEAPP"
	case win.WM_AFXFIRST:
		return "WM_AFXFIRST"
	case win.WM_AFXLAST:
		return "WM_AFXLAST"
	case win.WM_ASKCBFORMATNAME:
		return "WM_ASKCBFORMATNAME"
	case win.WM_CANCELJOURNAL:
		return "WM_CANCELJOURNAL"
	case win.WM_CANCELMODE:
		return "WM_CANCELMODE"
	case win.WM_CAPTURECHANGED:
		return "WM_CAPTURECHANGED"
	case win.WM_CHANGECBCHAIN:
		return "WM_CHANGECBCHAIN"
	case win.WM_CHAR:
		return "WM_CHAR"
	case win.WM_CHARTOITEM:
		return "WM_CHARTOITEM"
	case win.WM_CHILDACTIVATE:
		return "WM_CHILDACTIVATE"
	case win.WM_CLEAR:
		return "WM_CLEAR"
	case win.WM_CLOSE:
		return "WM_CLOSE"
	case win.WM_COMMAND:
		return "WM_COMMAND"
	case win.WM_COMMNOTIFY:
		return "WM_COMMNOTIFY"
	case win.WM_COMPACTING:
		return "WM_COMPACTING"
	case win.WM_COMPAREITEM:
		return "WM_COMPAREITEM"
	case win.WM_CONTEXTMENU:
		return "WM_CONTEXTMENU"
	case win.WM_COPY:
		return "WM_COPY"
	case win.WM_COPYDATA:
		return "WM_COPYDATA"
	case win.WM_CREATE:
		return "WM_CREATE"
	case win.WM_CTLCOLORBTN:
		return "WM_CTLCOLORBTN"
	case win.WM_CTLCOLORDLG:
		return "WM_CTLCOLORDLG"
	case win.WM_CTLCOLOREDIT:
		return "WM_CTLCOLOREDIT"
	case win.WM_CTLCOLORLISTBOX:
		return "WM_CTLCOLORLISTBOX"
	case win.WM_CTLCOLORMSGBOX:
		return "WM_CTLCOLORMSGBOX"
	case win.WM_CTLCOLORSCROLLBAR:
		return "WM_CTLCOLORSCROLLBAR"
	case win.WM_CTLCOLORSTATIC:
		return "WM_CTLCOLORSTATIC"
	case win.WM_CUT:
		return "WM_CUT"
	case win.WM_DEADCHAR:
		return "WM_DEADCHAR"
	case win.WM_DELETEITEM:
		return "WM_DELETEITEM"
	case win.WM_DESTROY:
		return "WM_DESTROY"
	case win.WM_DESTROYCLIPBOARD:
		return "WM_DESTROYCLIPBOARD"
	case win.WM_DEVICECHANGE:
		return "WM_DEVICECHANGE"
	case win.WM_DEVMODECHANGE:
		return "WM_DEVMODECHANGE"
	case win.WM_DISPLAYCHANGE:
		return "WM_DISPLAYCHANGE"
	case win.WM_DPICHANGED:
		return "WM_DPICHANGED"
	case win.WM_DRAWCLIPBOARD:
		return "WM_DRAWCLIPBOARD"
	case win.WM_DRAWITEM:
		return "WM_DRAWITEM"
	case win.WM_DROPFILES:
		return "WM_DROPFILES"
	case win.WM_ENABLE:
		return "WM_ENABLE"
	case win.WM_ENDSESSION:
		return "WM_ENDSESSION"
	case win.WM_ENTERIDLE:
		return "WM_ENTERIDLE"
	case win.WM_ENTERMENULOOP:
		return "WM_ENTERMENULOOP"
	case win.WM_ENTERSIZEMOVE:
		return "WM_ENTERSIZEMOVE"
	case win.WM_ERASEBKGND:
		return "WM_ERASEBKGND"
	case win.WM_EXITMENULOOP:
		return "WM_EXITMENULOOP"
	case win.WM_EXITSIZEMOVE:
		return "WM_EXITSIZEMOVE"
	case win.WM_FONTCHANGE:
		return "WM_FONTCHANGE"
	case win.WM_GETDLGCODE:
		return "WM_GETDLGCODE"
	case win.WM_GETFONT:
		return "WM_GETFONT"
	case win.WM_GETHOTKEY:
		return "WM_GETHOTKEY"
	case win.WM_GETICON:
		return "WM_GETICON"
	case win.WM_GETMINMAXINFO:
		return "WM_GETMINMAXINFO"
	case win.WM_GETTEXT:
		return "WM_GETTEXT"
	case win.WM_GETTEXTLENGTH:
		return "WM_GETTEXTLENGTH"
	case win.WM_HANDHELDFIRST:
		return "WM_HANDHELDFIRST"
	case win.WM_HANDHELDLAST:
		return "WM_HANDHELDLAST"
	case win.WM_HELP:
		return "WM_HELP"
	case win.WM_HOTKEY:
		return "WM_HOTKEY"
	case win.WM_HSCROLL:
		return "WM_HSCROLL"
	case win.WM_HSCROLLCLIPBOARD:
		return "WM_HSCROLLCLIPBOARD"
	case win.WM_ICONERASEBKGND:
		return "WM_ICONERASEBKGND"
	case win.WM_INITDIALOG:
		return "WM_INITDIALOG"
	case win.WM_INITMENU:
		return "WM_INITMENU"
	case win.WM_INITMENUPOPUP:
		return "WM_INITMENUPOPUP"
	case win.WM_INPUT:
		return "WM_INPUT"
	case win.WM_INPUTLANGCHANGE:
		return "WM_INPUTLANGCHANGE"
	case win.WM_INPUTLANGCHANGEREQUEST:
		return "WM_INPUTLANGCHANGEREQUEST"
	case win.WM_KEYDOWN:
		return "WM_KEYDOWN"
	case win.WM_KEYUP:
		return "WM_KEYUP"
	case win.WM_KILLFOCUS:
		return "WM_KILLFOCUS"
	case win.WM_MDIACTIVATE:
		return "WM_MDIACTIVATE"
	case win.WM_MDICASCADE:
		return "WM_MDICASCADE"
	case win.WM_MDICREATE:
		return "WM_MDICREATE"
	case win.WM_MDIDESTROY:
		return "WM_MDIDESTROY"
	case win.WM_MDIGETACTIVE:
		return "WM_MDIGETACTIVE"
	case win.WM_MDIICONARRANGE:
		return "WM_MDIICONARRANGE"
	case win.WM_MDIMAXIMIZE:
		return "WM_MDIMAXIMIZE"
	case win.WM_MDINEXT:
		return "WM_MDINEXT"
	case win.WM_MDIREFRESHMENU:
		return "WM_MDIREFRESHMENU"
	case win.WM_MDIRESTORE:
		return "WM_MDIRESTORE"
	case win.WM_MDISETMENU:
		return "WM_MDISETMENU"
	case win.WM_MDITILE:
		return "WM_MDITILE"
	case win.WM_MEASUREITEM:
		return "WM_MEASUREITEM"
	case win.WM_GETOBJECT:
		return "WM_GETOBJECT"
	case win.WM_CHANGEUISTATE:
		return "WM_CHANGEUISTATE"
	case win.WM_UPDATEUISTATE:
		return "WM_UPDATEUISTATE"
	case win.WM_QUERYUISTATE:
		return "WM_QUERYUISTATE"
	case win.WM_UNINITMENUPOPUP:
		return "WM_UNINITMENUPOPUP"
	case win.WM_MENURBUTTONUP:
		return "WM_MENURBUTTONUP"
	case win.WM_MENUCOMMAND:
		return "WM_MENUCOMMAND"
	case win.WM_MENUGETOBJECT:
		return "WM_MENUGETOBJECT"
	case win.WM_MENUDRAG:
		return "WM_MENUDRAG"
	case win.WM_APPCOMMAND:
		return "WM_APPCOMMAND"
	case win.WM_MENUCHAR:
		return "WM_MENUCHAR"
	case win.WM_MENUSELECT:
		return "WM_MENUSELECT"
	case win.WM_MOVE:
		return "WM_MOVE"
	case win.WM_MOVING:
		return "WM_MOVING"
	case win.WM_NCACTIVATE:
		return "WM_NCACTIVATE"
	case win.WM_NCCALCSIZE:
		return "WM_NCCALCSIZE"
	case win.WM_NCCREATE:
		return "WM_NCCREATE"
	case win.WM_NCDESTROY:
		return "WM_NCDESTROY"
	case win.WM_NCHITTEST:
		return "WM_NCHITTEST"
	case win.WM_NCLBUTTONDBLCLK:
		return "WM_NCLBUTTONDBLCLK"
	case win.WM_NCLBUTTONDOWN:
		return "WM_NCLBUTTONDOWN"
	case win.WM_NCLBUTTONUP:
		return "WM_NCLBUTTONUP"
	case win.WM_NCMBUTTONDBLCLK:
		return "WM_NCMBUTTONDBLCLK"
	case win.WM_NCMBUTTONDOWN:
		return "WM_NCMBUTTONDOWN"
	case win.WM_NCMBUTTONUP:
		return "WM_NCMBUTTONUP"
	case win.WM_NCXBUTTONDOWN:
		return "WM_NCXBUTTONDOWN"
	case win.WM_NCXBUTTONUP:
		return "WM_NCXBUTTONUP"
	case win.WM_NCXBUTTONDBLCLK:
		return "WM_NCXBUTTONDBLCLK"
	case win.WM_NCMOUSEHOVER:
		return "WM_NCMOUSEHOVER"
	case win.WM_NCMOUSELEAVE:
		return "WM_NCMOUSELEAVE"
	case win.WM_NCMOUSEMOVE:
		return "WM_NCMOUSEMOVE"
	case win.WM_NCPAINT:
		return "WM_NCPAINT"
	case win.WM_NCRBUTTONDBLCLK:
		return "WM_NCRBUTTONDBLCLK"
	case win.WM_NCRBUTTONDOWN:
		return "WM_NCRBUTTONDOWN"
	case win.WM_NCRBUTTONUP:
		return "WM_NCRBUTTONUP"
	case win.WM_NEXTDLGCTL:
		return "WM_NEXTDLGCTL"
	case win.WM_NEXTMENU:
		return "WM_NEXTMENU"
	case win.WM_NOTIFY:
		return "WM_NOTIFY"
	case win.WM_NOTIFYFORMAT:
		return "WM_NOTIFYFORMAT"
	case win.WM_NULL:
		return "WM_NULL"
	case win.WM_PAINT:
		return "WM_PAINT"
	case win.WM_PAINTCLIPBOARD:
		return "WM_PAINTCLIPBOARD"
	case win.WM_PAINTICON:
		return "WM_PAINTICON"
	case win.WM_PALETTECHANGED:
		return "WM_PALETTECHANGED"
	case win.WM_PALETTEISCHANGING:
		return "WM_PALETTEISCHANGING"
	case win.WM_PARENTNOTIFY:
		return "WM_PARENTNOTIFY"
	case win.WM_PASTE:
		return "WM_PASTE"
	case win.WM_PENWINFIRST:
		return "WM_PENWINFIRST"
	case win.WM_PENWINLAST:
		return "WM_PENWINLAST"
	case win.WM_POWER:
		return "WM_POWER"
	case win.WM_POWERBROADCAST:
		return "WM_POWERBROADCAST"
	case win.WM_PRINT:
		return "WM_PRINT"
	case win.WM_PRINTCLIENT:
		return "WM_PRINTCLIENT"
	case win.WM_QUERYDRAGICON:
		return "WM_QUERYDRAGICON"
	case win.WM_QUERYENDSESSION:
		return "WM_QUERYENDSESSION"
	case win.WM_QUERYNEWPALETTE:
		return "WM_QUERYNEWPALETTE"
	case win.WM_QUERYOPEN:
		return "WM_QUERYOPEN"
	case win.WM_QUEUESYNC:
		return "WM_QUEUESYNC"
	case win.WM_QUIT:
		return "WM_QUIT"
	case win.WM_RENDERALLFORMATS:
		return "WM_RENDERALLFORMATS"
	case win.WM_RENDERFORMAT:
		return "WM_RENDERFORMAT"
	case win.WM_SETCURSOR:
		return "WM_SETCURSOR"
	case win.WM_SETFOCUS:
		return "WM_SETFOCUS"
	case win.WM_SETFONT:
		return "WM_SETFONT"
	case win.WM_SETHOTKEY:
		return "WM_SETHOTKEY"
	case win.WM_SETICON:
		return "WM_SETICON"
	case win.WM_SETREDRAW:
		return "WM_SETREDRAW"
	case win.WM_SETTEXT:
		return "WM_SETTEXT"
	case win.WM_SETTINGCHANGE:
		return "WM_SETTINGCHANGE"
	case win.WM_SHOWWINDOW:
		return "WM_SHOWWINDOW"
	case win.WM_SIZE:
		return "WM_SIZE"
	case win.WM_SIZECLIPBOARD:
		return "WM_SIZECLIPBOARD"
	case win.WM_SIZING:
		return "WM_SIZING"
	case win.WM_SPOOLERSTATUS:
		return "WM_SPOOLERSTATUS"
	case win.WM_STYLECHANGED:
		return "WM_STYLECHANGED"
	case win.WM_STYLECHANGING:
		return "WM_STYLECHANGING"
	case win.WM_SYSCHAR:
		return "WM_SYSCHAR"
	case win.WM_SYSCOLORCHANGE:
		return "WM_SYSCOLORCHANGE"
	case win.WM_SYSCOMMAND:
		return "WM_SYSCOMMAND"
	case win.WM_SYSDEADCHAR:
		return "WM_SYSDEADCHAR"
	case win.WM_SYSKEYDOWN:
		return "WM_SYSKEYDOWN"
	case win.WM_SYSKEYUP:
		return "WM_SYSKEYUP"
	case win.WM_TCARD:
		return "WM_TCARD"
	case win.WM_THEMECHANGED:
		return "WM_THEMECHANGED"
	case win.WM_TIMECHANGE:
		return "WM_TIMECHANGE"
	case win.WM_TIMER:
		return "WM_TIMER"
	case win.WM_UNDO:
		return "WM_UNDO"
	case win.WM_USER:
		return "WM_USER"
	case win.WM_USERCHANGED:
		return "WM_USERCHANGED"
	case win.WM_VKEYTOITEM:
		return "WM_VKEYTOITEM"
	case win.WM_VSCROLL:
		return "WM_VSCROLL"
	case win.WM_VSCROLLCLIPBOARD:
		return "WM_VSCROLLCLIPBOARD"
	case win.WM_WINDOWPOSCHANGED:
		return "WM_WINDOWPOSCHANGED"
	case win.WM_WINDOWPOSCHANGING:
		return "WM_WINDOWPOSCHANGING"
	case win.WM_SYNCPAINT:
		return "WM_SYNCPAINT"
	case win.WM_MOUSEACTIVATE:
		return "WM_MOUSEACTIVATE"
	case win.WM_MOUSEMOVE:
		return "WM_MOUSEMOVE"
	case win.WM_LBUTTONDOWN:
		return "WM_LBUTTONDOWN"
	case win.WM_LBUTTONUP:
		return "WM_LBUTTONUP"
	case win.WM_LBUTTONDBLCLK:
		return "WM_LBUTTONDBLCLK"
	case win.WM_RBUTTONDOWN:
		return "WM_RBUTTONDOWN"
	case win.WM_RBUTTONUP:
		return "WM_RBUTTONUP"
	case win.WM_RBUTTONDBLCLK:
		return "WM_RBUTTONDBLCLK"
	case win.WM_MBUTTONDOWN:
		return "WM_MBUTTONDOWN"
	case win.WM_MBUTTONUP:
		return "WM_MBUTTONUP"
	case win.WM_MBUTTONDBLCLK:
		return "WM_MBUTTONDBLCLK"
	case win.WM_MOUSEWHEEL:
		return "WM_MOUSEWHEEL"
	case win.WM_XBUTTONDOWN:
		return "WM_XBUTTONDOWN"
	case win.WM_XBUTTONUP:
		return "WM_XBUTTONUP"
	case win.WM_XBUTTONDBLCLK:
		return "WM_XBUTTONDBLCLK"
	case win.WM_MOUSEHOVER:
		return "WM_MOUSEHOVER"
	case win.WM_MOUSELEAVE:
		return "WM_MOUSELEAVE"
	case win.WM_CLIPBOARDUPDATE:
		return "WM_CLIPBOARDUPDATE"
	default:
		return "UNKNOWN"
	}
}
