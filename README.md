# Simple chat in Go using GoFlow library #

This application was originally built to demonstrate Flow-based programming in Go programming language using [GoFlow](https://github.com/trustmaster/goflow) library.

## Installation ##

You need at least Go weekly.2012-01-27 to compile this example and some proxy webserver to behave as front end (e.g. nginx).

Install it using go tool:

```
go get https://github.com/trustmaster/gochat
```

Then you can run it in command line using

```
$GOPATH/bin/gochat
```

or

```
$GOBIN/gochat
```

depending on your Go configuration. The chat server will listen on localhost port 9090 by default.

Copy index.html and jquery-1.x.x.js to your webserver folder, e.g. gochat. You need to modify your webserver configuration to proxy requests for the chat server, otherwise AJAX won't work. Here is an example from my nginx.conf:

```
		location /gochat/chat {
           		proxy_redirect     off;

           		proxy_set_header   Host             "locahost:9090";
          		proxy_set_header   X-Real-IP        $remote_addr;
           		proxy_set_header  X-Forwarded-For  $proxy_add_x_forwarded_for;
			proxy_pass         http://localhost:9090/;
		}
```

Then you can open the client part in your browser and try it, e.g. http://localhost/gochat/

## Description ##

The client side of this app is a simple HTML page powered by JavaScript retrieving JSON data from server and posting new messages via AJAX.

Architecture of the server side is shown on the following diagram:

![GoChat structure diagram](http://flowbased.wdfiles.com/local--files/goflow/gochat.png)

The part inside the "App" frame is a flow-based network. Outside it there is just an adapter for standard HTTP server interface of "net/http" package. Here are descriptions for each component:

* Router checks the path for incoming HTTP response packet and decides what controller should handle it. Note: currently it checks request method rather than path because I haven't configured my Nginx proxy to pass paths corrently along with POST method.
* Controller checks arguments for a GET request and forms a query to a Storage to retrieve new message since last query.
* SendController prepares a message for submission into the Storage.
* Storage manages access to a message Queue in memory. It can add new messages or fetch existing ones and pass them as a packet to Responder.
* Responder encloses the results in JSON response.

The data travels between processes as packets (structure) described in common.go. Each packet contains an instance of http.Request and http.ResponseWriter so it can be returned to the client at any moment. Packets are passed through the channels by pointers, so pointers travel instead of structures themselves.