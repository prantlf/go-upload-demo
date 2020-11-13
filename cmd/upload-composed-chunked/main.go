package main

import (
	"log"

	"github.com/prantlf/go-upload-demo/internal/composed/chunked"
	"github.com/prantlf/go-upload-demo/internal/params"
	"github.com/prantlf/go-upload-demo/internal/requestor"
)

func main() {
	reqBody, contentType, err := chunked.NewRequestBody(params.Comment, params.File)
	if err != nil {
		log.Fatal(err)
	}
	requestor.Post(params.URL, contentType, reqBody)
}
