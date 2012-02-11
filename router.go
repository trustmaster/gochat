package main

import (
	"github.com/trustmaster/goflow"
	"net/http"
	"strings"
)

// Simple table-based request router
type Router struct {
	flow.Component
	In   chan *RequestPacket
	Show chan<- *RequestPacket
	Send chan<- *RequestPacket
}

// Base path to chat from webserver root
var BasePath = "/"

// Request handler
func (r *Router) OnIn(p *RequestPacket) {
	path := p.Req.URL.Path
	if strings.Index(path, BasePath) == 0 {
		path = path[len(BasePath):]
	}
	var fwd chan<- *RequestPacket
	ok := true
	//	switch path { // broken in my nginx
	//	case "/send":
	//		fwd = r.Send
	//	case "/":
	//		fwd = r.Show
	//	default:
	//		ok = false
	//	}
	if p.Req.Method == "POST" {
		fwd = r.Send
	} else {
		fwd = r.Show
	}
	if !ok {
		p.Error(http.StatusBadRequest, "Unknown path: "+path)
	} else {
		fwd <- p
	}
}
