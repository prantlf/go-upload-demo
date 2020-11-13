package main

import (
	"log"

	"github.com/prantlf/go-upload-demo/internal/buffered"
	"github.com/prantlf/go-upload-demo/internal/params"
	"github.com/prantlf/go-upload-demo/internal/requestor"
)

func main() {
	reqBody, contentType, err := buffered.NewRequestBody(params.Comment, params.File)
	if err != nil {
		log.Fatal(err)
	}
	requestor.Post(params.URL, contentType, reqBody)
}
