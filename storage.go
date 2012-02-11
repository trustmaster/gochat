package main

import (
	"github.com/trustmaster/goflow"
	"net/http"
)

// Default storage capacity
const DefaultQueueCapacity int = 50

// Provides data storage and retrieval services
type Storage struct {
	flow.Component
	MaxLen <-chan int
	Get    <-chan *GetRequestPacket
	Post   <-chan *PostRequestPacket
	Out    chan<- *RequestPacket

	queue *Queue
}

// Constructs a new store server
func NewStorage() *Storage {
	s := new(Storage)
	s.queue = NewQueue(DefaultQueueCapacity)
	return s
}

// Changes current max queue length
func (s *Storage) OnMaxLen(ml int) {
	s.queue.SetMaxLen(ml)
}

// Gets messages from the store and sends the to out
func (s *Storage) OnGet(p *GetRequestPacket) {
	msgs := s.queue.Find(p.Since)
	p.Data = msgs
	s.Out <- p.RequestPacket
}

// Adds a new message to the store and sends status to out
func (s *Storage) OnPost(p *PostRequestPacket) {
	t := s.queue.Push(p.Author, p.Text)
	p.Code = http.StatusCreated
	p.Data = Message{Time: t, Author: p.Author, Text: p.Text}
	s.Out <- p.RequestPacket
}
