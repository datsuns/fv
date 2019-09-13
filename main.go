package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
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
				},
			},
		},
	}.Run()
}
