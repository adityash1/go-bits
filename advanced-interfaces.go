package main

import "fmt"

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type FireBaseStorage struct {
	// connect with firebase storage
}

func (f FireBaseStorage) Get(id int) (any, error) {
	return nil, nil
}

func (f FireBaseStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

// dependency on storage struct (bad practice)
func updateValue(id int, val any) error {
	store := FireBaseStorage{}
	return store.Put(id, val)
}

// if there are too many methods in an interface and we need only one ? (bad practice)
func updateValue(id int, val any, store Storage) error {
	return store.Put(id, val)
}

// make individual interfaces for the methods
type Putter interface {
	Put(id int, val any) error
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}

type Storage interface {
	// instead of duplicating the method implement putter interface too
	// Put(id int, val any) error
	Putter
	Get(id int) (any, error)
}

type simplePutter struct {
}

func (s simplePutter) Put(id int, val any) error {
	return nil
}

func (s simplePutter) String() string {
	return ""
}

func main() {
	s := Server{
		store: FireBaseStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "aditya")
	updateValue(1, "vasundhra", s.store)
	fmt.Printf("%+v", s)

	sputter := simplePutter{}
	updateValue(1, "harsh", sputter)
	fmt.Printf("%#+v", sputter)
}
