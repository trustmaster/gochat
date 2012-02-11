package main

import (
	"github.com/trustmaster/goflow"
	"net/http"
)

// Controller for send operations
type SendController struct {
	flow.Component
	In  <-chan *RequestPacket
	Out chan<- *PostRequestPacket
}

// Parses a POST request and passes a new message to a Storage
func (c *SendController) OnIn(p *RequestPacket) {
	if p.Req.Method == "POST" {
		c.Out <- &PostRequestPacket{RequestPacket: p, Author: p.Req.FormValue("author"), Text: p.Req.FormValue("text")}
	} else {
		p.Error(http.StatusBadRequest, "Bad request method: "+p.Req.Method)
	}
}
