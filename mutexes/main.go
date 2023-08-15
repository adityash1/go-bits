package main

import (
	"sync"
)

type State struct {
	count int
	mu    sync.Mutex
}

func (s *State) SetState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
}

func main() {}
