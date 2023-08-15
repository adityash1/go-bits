package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("server started")
	s.loop()
}

func (s *Server) quit() {
	// close(s.quitch)
	s.quitch <- struct{}{}
}

func (s *Server) loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			break mainloop
		case msg := <-s.msgch:
			s.handleMessage(msg)
		default:
		}
	}
	fmt.Println("server shutting down gracefully")
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("message: ", msg)
}

func main() {
	server := newServer()

	go func() {
		time.Sleep(time.Second * 5)
		server.quit()
	}()

	server.start()

	// go server.start()
	// for i := 0; i < 100; i++ {
	// 	server.sendMessage(fmt.Sprintf("number %d", i))
	// }
}
