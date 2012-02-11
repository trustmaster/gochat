== Simple chat in Go using GoFlow library ==

This application was originally built to demonstrate Flow-based programming in Go programming language using [GoFlow](https://github.com/trustmaster/goflow) library.

== Installation ==

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

== Description ==

Coming soon...