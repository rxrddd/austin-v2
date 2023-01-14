package manager

import (
	"errors"
)

type Manager struct {
	mMap map[string]Item
}

func NewManager(items ...Item) *Manager {
	hs := &Manager{}
	hs.Init(items...)
	return hs
}

func (hs *Manager) Init(items ...Item) {
	for _, item := range items {
		hs.Register(item)
	}
}

func (hs *Manager) Register(h Item) {
	if hs.mMap == nil {
		hs.mMap = make(map[string]Item)
	}
	hs.mMap[h.Name()] = h
}

func (hs *Manager) Get(key string) (Item, error) {
	if h, ok := hs.mMap[key]; ok {
		return h, nil
	}
	return nil, errors.New("unknown item " + key)
}

type Item interface {
	Name() string
}
