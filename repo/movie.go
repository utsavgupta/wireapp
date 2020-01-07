package repo

import (
	"errors"
	"sync"

	"github.com/utsavgupta/wireapp/ent"
)

type Movie interface {
	LoadItems() []ent.Movie
	SaveItem(*ent.Movie) error
}

type movie struct {
	lock  *sync.Mutex
	items []ent.Movie
}

func NewMovie(size int) Movie {
	i := make([]ent.Movie, 0, size)
	l := &sync.Mutex{}

	return &movie{items: i, lock: l}
}

func (m *movie) LoadItems() []ent.Movie {
	return m.items
}

func (m *movie) SaveItem(newMovie *ent.Movie) error {

	if len(m.items) >= cap(m.items) {
		return errors.New("Memory overflow")
	}

	m.items = append(m.items, *newMovie)

	return nil
}
