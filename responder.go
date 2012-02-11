package main

import (
	"encoding/json"
	"github.com/trustmaster/goflow"
	"net/http"
)

// Simple JSON response generator
type Responder struct {
	flow.Component
	In <-chan *RequestPacket
}

// Processes a request packet and sends the response JSON
func (r *Responder) OnIn(p *RequestPacket) {
	js, err := json.Marshal(p.Data)
	if err != nil {
		p.Error(http.StatusInternalServerError, "Could not marshal JSON")
		return
	}
	p.Res.Write(js)
	p.Done <- true
}
