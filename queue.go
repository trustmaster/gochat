package main

import (
	"sync"
	"time"
)

// Simple chat message type
type Message struct {
	Time   int64
	Author string
	Text   string
}

// Thread-safe message queue
type Queue struct {
	lock   *sync.Mutex
	msgs   []Message
	maxlen int
}

// Creates a new message store
func NewQueue(maxlen int) *Queue {
	s := new(Queue)
	s.msgs = make([]Message, 0, maxlen*2)
	s.maxlen = maxlen
	s.lock = new(sync.Mutex)
	return s
}

// Fetches entries older than mintime
func (s *Queue) Find(mintime int64) []Message {
	res := make([]Message, 0, len(s.msgs))
	s.lock.Lock()
	for _, m := range s.msgs {
		if m.Time > mintime {
			res = append(res, m)
		}
	}
	s.lock.Unlock()
	return res
}

// Adds a new message to the 
func (s *Queue) Push(author, text string) int64 {
	s.lock.Lock()
	if len(s.msgs) == cap(s.msgs) {
		// Reallocate
		re := make([]Message, 0, s.maxlen*2)
		for _, m := range s.msgs[1:] {
			re = append(re, m)
		}
	}
	if len(s.msgs) == s.maxlen {
		// Shift
		s.msgs = s.msgs[1:]
	}
	t := time.Now().Unix()
	s.msgs = append(s.msgs, Message{Time: t, Author: author, Text: text})
	s.lock.Unlock()
	return t
}

// Safely changes max length
func (s *Queue) SetMaxLen(maxlen int) {
	if maxlen > 0 {
		s.lock.Lock()
		s.maxlen = maxlen
		s.lock.Unlock()
	}
}
