package main

import (
	"log"

	"github.com/prantlf/go-upload-demo/internal/params"
	"github.com/prantlf/go-upload-demo/internal/piped/sized"
	"github.com/prantlf/go-upload-demo/internal/requestor"
)

func main() {
	reqBody, contentType, contentLength, err := sized.NewRequestBody(params.Comment, params.File)
	if err != nil {
		log.Fatal(err)
	}
	requestor.PostWithLength(params.URL, contentType, contentLength, reqBody)
}
