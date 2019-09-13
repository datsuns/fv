// Copyright 2012 The Walk Authors. All rights reserved.
// Use of lb source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package main

import (
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

type MyListBox struct {
	*walk.ListBox
	model *MyModel
}

func NewMyListBox(parent walk.Container) (*MyListBox, error) {
	lb, err := walk.NewListBox(parent)
	if err != nil {
		return nil, err
	}
	mlb := &MyListBox{lb, NewMyModel()}
	lb.SetModel(mlb.model)

	if err := walk.InitWrapperWindow(mlb); err != nil {
		return nil, err
	}

	return mlb, nil
}

func (mlb *MyListBox) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_KEYDOWN:
		if walk.Key(wParam) == walk.KeyReturn {
			return mlb.ListBox.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
		}
	}
	return mlb.ListBox.WndProc(hwnd, msg, wParam, lParam)
}

type MyModel struct {
	walk.ListModelBase
	items []string
}

func NewMyModel() *MyModel {
	m := &MyModel{items: make([]string, 2)}
	m.items[0] = "hello"
	m.items[1] = "my"
	return m
}

func (m *MyModel) ItemCount() int {
	return len(m.items)
}

func (m *MyModel) Value(index int) interface{} {
	return m.items[index]
}
