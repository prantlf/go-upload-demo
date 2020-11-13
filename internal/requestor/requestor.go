package requestor

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Post(URL, contentType string, reqBody io.Reader) {
	resp, err := http.DefaultClient.Post(URL, contentType, reqBody)
	handleResponse(resp, err)
}

func PostWithLength(URL, contentType string, contentLength int64, reqBody io.Reader) {
	req, err := http.NewRequest("POST", URL, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", contentType)
	req.ContentLength = contentLength
	resp, err := http.DefaultClient.Do(req)
	handleResponse(resp, err)
}

func handleResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
