package main

import (
	"encoding/json"
	"net/http"
)

type RequestPacket struct {
	Req  *http.Request
	Res  http.ResponseWriter
	Code int
	Data interface{}
	Done chan bool
}

// Immediately pops the request with error response
func (p *RequestPacket) Error(code int, msg string) {
	p.Res.WriteHeader(code)
	js, _ := json.Marshal(Error{Code: code, Msg: msg})
	p.Res.Write(js)
	p.Done <- true
}

type GetRequestPacket struct {
	*RequestPacket
	Since int64
}

type PostRequestPacket struct {
	*RequestPacket
	Author string
	Text   string
}

type Error struct {
	Code int
	Msg  string
}
