# http-samples

Golang HTTP server samples for my practice :smile:

## Usage

server:

```sh
$ go run server.go
```

client:

```sh
$ curl -v http://localhost:8080/foo
* Adding handle: conn: 0x7fcab9003a00
* Adding handle: send: 0
* Adding handle: recv: 0
* Curl_addHandleToPipeline: length: 1
* - Conn 0 (0x7fcab9003a00) send_pipe: 1, recv_pipe: 0
* About to connect() to localhost port 8080 (#0)
*   Trying ::1...
* Connected to localhost (::1) port 8080 (#0)
> GET /foo HTTP/1.1
> User-Agent: curl/7.30.0
> Host: localhost:8080
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 08 Oct 2014 16:37:48 GMT
< Content-Length: 6
< Content-Type: text/plain; charset=utf-8
< 
Hello
* Connection #0 to host localhost left intact
```

## License

Copyright (c) 2014 Soichiro Kashima  
Licensed under MIT license.  
See the bundled [LICENSE](LICENSE) file for details.
