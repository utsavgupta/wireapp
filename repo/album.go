package repo

import (
	"errors"
	"sync"

	"github.com/utsavgupta/wireapp/ent"
)

type AlbumBuffer int

type Album interface {
	LoadItems() []ent.Album
	SaveItem(*ent.Album) error
}

type album struct {
	lock  *sync.Mutex
	items []ent.Album
}

func NewAlbum(size AlbumBuffer) Album {
	i := make([]ent.Album, 0, size)
	l := &sync.Mutex{}

	return &album{items: i, lock: l}
}

func (a *album) LoadItems() []ent.Album {
	return a.items
}

func (a *album) SaveItem(newAlbum *ent.Album) error {

	if len(a.items) >= cap(a.items) {
		return errors.New("Memory overflow")
	}

	a.items = append(a.items, *newAlbum)

	return nil
}
