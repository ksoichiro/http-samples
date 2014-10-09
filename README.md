# http-samples

Golang HTTP server samples for my practice :smile:

* Very easy!
* Maybe you can use these samples
  to write your own mock server
  for mobile apps' CI and automated tests.

## Usage

### Hello world

server:

```sh
$ go run hello.go -port 9000
Serves HTTP in port 9000
```

Quit: Ctrl + C.

client:

```sh
$ curl -v http://localhost:9000/hello
* Adding handle: conn: 0x7fb7da003a00
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x7fb7da003a00) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 9000 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 9000 (#0)
> GET /hello HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:9000
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 08 Oct 2014 17:06:55 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
< 
Hello, world!
* Connection #0 to host localhost left intact
```

### Basic authentication

This sample uses [goji/httpauth](https://github.com/goji/httpauth).

server:

```sh
$ go run auth.go -port 9000
Serves HTTP in port 9000
```

Quit: Ctrl + C.

client:

```sh
$ curl -v --basic -u guest:foo http://localhost:9000/                                                                                                                               * Adding handle: conn: 0x7fba5c80aa00
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x7fba5c80aa00) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 9000 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 9000 (#0)
* Server auth using Basic with user 'guest'
> GET / HTTP/1.1
> Authorization: Basic Z3Vlc3Q6Zm9v
> User-Agent: curl/7.30.0
> Host: localhost:9000
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Thu, 09 Oct 2014 16:31:40 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
< 
Hello, world!
* Connection #0 to host localhost left intact
```

```sh
$ curl -v http://localhost:9000/
* Adding handle: conn: 0x7ff572803a00
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x7ff572803a00) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 9000 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 9000 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:9000
> Accept: */*
> 
< HTTP/1.1 401 Unauthorized
< Content-Type: text/plain; charset=utf-8
< Www-Authenticate: Basic realm=""Restricted""
< Date: Thu, 09 Oct 2014 16:35:11 GMT
< Content-Length: 13
< 
Unauthorized
* Connection #0 to host localhost left intact
```

## License

Copyright (c) 2014 Soichiro Kashima  
Licensed under MIT license.  
See the bundled [LICENSE](LICENSE) file for details.
