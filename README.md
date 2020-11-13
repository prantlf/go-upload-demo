# go-upload-demo

Demonstrates how to make an optimal HTTP request with the multipart MIME message according to [RFC7578] in Go. It uploads a file including metadata to an echoing server, which just prints the request header and body on the console.

The best solutions for both chunked and sized (with content length) are [upload-composed-chunked] and [upload-composed-sized], which use a custom [multipart composer]. Other solutions use the standard [multipart writer].

Command line executables; see the descriptions pointed at by these links:

* [serve-echo] - HTTP server, which just prints the request header and body on the console.
* [upload-buffered] - reads the entire file to memory, posts the content length
* [upload-piped-chunked] - streams the file using a pipe, does not post the content length
* [upload-piped-sized] - streams the file using a pipe, posts the content length
* [upload-composed-chunked] - streams the file using a reader, does not post the content length
* [upload-composed-sized] - streams the file using a reader, posts the content length

The server and the uploading executables use [common parameters] and [common request & response handling] to make the testing easier.

## Usage

Install Go and GNU Make. For example, on Mac:

    brew install go make

Clone this repository and build the executable:

    git clone https://github.com/prantlf/go-upload-demo.git
    cd go-upload-demo
    make

Run the testing web server, which prints requests on the console, in the background:

    ./serve-echo &

Run the testing clients, which use different techniques to send a multi-part form request:

    ./upload-buffered
    ./upload-piped-chunked
    ./upload-piped-sized
    ./upload-composed-chunked
    ./upload-composed-sized

[RFC7578]: https://tools.ietf.org/html/rfc7578
[multipart composer]: https://pkg.go.dev/github.com/prantlf/go-multipart-composer
[multipart writer]: https://golang.org/pkg/mime/multipart/#Writer
[serve-echo]: cmd/serve-echo/serve-echo.md
[upload-buffered]: cmd/upload-buffered/upload-buffered.md
[upload-piped-chunked]: cmd/upload-piped-chunked/upload-piped-chunked.md
[upload-piped-sized]: cmd/upload-piped-sized/upload-piped-sized.md
[upload-composed-chunked]: cmd/upload-composed-chunked/upload-composed-chunked.md
[upload-composed-sized]: cmd/upload-composed-sized/upload-composed-sized.md
[common parameters]: internal/params/params.go
[common request & response handling]: internal/requestor/requestor.go
