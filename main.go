package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

var f declarative.Font = declarative.Font{
	PointSize: 20,
}

func main() {
	var mw *walk.MainWindow

	fmt.Println("hello")

	declarative.MainWindow{
		AssignTo: &mw,
		Title:    "SCREAMO",
		MinSize:  declarative.Size{600, 400},
		Layout:   declarative.HBox{},
		Font:     f,
	}.Create()

	left, err := NewMyListBox(mw, "LEFT", ".")
	if err != nil {
		panic(err)
	}
	right, err := NewMyListBox(mw, "RIGHT", ".")
	if err != nil {
		panic(err)
	}
	left.RegisterNeigbor(right, walk.KeyL)
	right.RegisterNeigbor(left, walk.KeyH)
	mw.Run()
}
