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
}

func NewMyListBox(parent walk.Container) (*MyListBox, error) {
	lb, err := walk.NewListBox(parent)
	if err != nil {
		return nil, err
	}

	mlb := &MyListBox{lb}

	if err := walk.InitWrapperWindow(mlb); err != nil {
		return nil, err
	}

	return mlb, nil
}

func (lb *MyListBox) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return lb.ListBox.WndProc(hwnd, msg, wParam, lParam)
}
