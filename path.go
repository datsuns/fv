package main

import (
	"path/filepath"
)

type Path struct {
	Current string
	Childs  []string
}

func NewPath(root string) *Path {
	ret := &Path{Current: root}
	return ret
}

func (p *Path) Load() error {
	entries, err := filepath.Glob(filepath.Join(p.Current, "*"))
	if err != nil {
		return err
	}
	for _, e := range entries {
		p.Childs = append(p.Childs, filepath.Base(e))
	}
	return nil
}

func (p *Path) Len() int {
	return len(p.Childs)
}

func (p *Path) At(i int) interface{} {
	return p.Childs[i]
}
