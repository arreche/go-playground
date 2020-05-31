package domain

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
)

// Repository public interface
type Repository interface {
	Save(entity *Entity) (*Entity, error)
	Get(id string) (*Entity, error)
	Delete(id string) error
	List() []*Entity
	Size() int
}

type storage map[string]*Entity

type repository struct {
	sync.Mutex
	table map[string]*Entity
}

// NewRepository public factory method
func NewRepository() Repository {
	return &repository{table: make(storage)}
}

func (r *repository) Save(e *Entity) (*Entity, error) {
	r.Lock()
	defer r.Unlock()
	if len(e.ID) == 0 {
		e.ID = strconv.Itoa(rand.Int())
	}
	r.table[e.ID] = e
	return e, nil
}

func (r *repository) Get(id string) (*Entity, error) {
	r.Lock()
	defer r.Unlock()
	e, ok := r.table[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return e, nil
}

func (r *repository) Delete(id string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.table[id]; !ok {
		return errors.New("not found")
	}
	delete(r.table, id)
	return nil
}

func (r *repository) List() []*Entity {
	r.Lock()
	defer r.Unlock()
	list := make([]*Entity, 0, len(r.table))
	for _, v := range r.table {
		list = append(list, v)
	}
	return list
}

func (r *repository) Size() int {
	r.Lock()
	defer r.Unlock()
	return len(r.table)
}
