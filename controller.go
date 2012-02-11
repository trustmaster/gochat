package main

import (
	"github.com/trustmaster/goflow"
	"strconv"
)

// Default controller for message retreival
type Controller struct {
	flow.Component
	In  <-chan *RequestPacket
	Out chan<- *GetRequestPacket
}

// Simple controller routing
func (c *Controller) OnIn(p *RequestPacket) {
	since, err := strconv.ParseInt(p.Req.FormValue("since"), 10, 64)
	if err != nil || since < 0 {
		since = 0
	}
	c.Out <- &GetRequestPacket{RequestPacket: p, Since: since}
}
