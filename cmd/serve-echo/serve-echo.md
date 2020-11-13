# serve-echo

The [HTTP server] listens on `localhost`, port `8182`, path `/`. See also the source code of the [common parameters].

An example of a printed request method, path, header with body:

    POST /

    User-Agent: Go-http-client/1.1
    Content-Length: 492
    Content-Type: multipart/form-data; boundary=64f99ebb45795f28a863
    Accept-Encoding: gzip

    --64f99ebb45795f28a863
    Content-Disposition: form-data; name="comment"

    module descriptor
    --64f99ebb45795f28a863
    Content-Disposition: form-data; name="file"; filename="go.mod"
    Content-Type: application/octet-stream

    module github.com/prantlf/go-upload-demo

    go 1.15

    --64f99ebb45795f28a863--

[HTTP server]: main.go
[common parameters]: ../../internal/params/params.go
