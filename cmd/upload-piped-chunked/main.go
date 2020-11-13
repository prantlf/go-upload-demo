package main

import (
	"github.com/prantlf/go-upload-demo/internal/params"
	"github.com/prantlf/go-upload-demo/internal/piped/chunked"
	"github.com/prantlf/go-upload-demo/internal/requestor"
)

func main() {
	reqBody, contentType := chunked.NewRequestBody(params.Comment, params.File)
	requestor.Post(params.URL, contentType, reqBody)
}
