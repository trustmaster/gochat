package main

import (
	"github.com/trustmaster/goflow"
	"net/http"
)

var (
	routerIn chan *RequestPacket
)

func handler(w http.ResponseWriter, r *http.Request) {
	rp := RequestPacket{Req: r, Res: w, Done: make(chan bool)}
	routerIn <- &rp
	<-rp.Done
}

func main() {
	// Create application net
	routerIn = make(chan *RequestPacket)
	net := NewApp()
	net.SetInPort("In", routerIn)
	// Run
	flow.RunNet(net)
	// Serve
	BasePath = "/gochat/chat"
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9090", nil)
}
