package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/prantlf/go-upload-demo/internal/params"
)

func main() {
	http.HandleFunc(params.Path, echo)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", params.Port), nil))
}

func echo(resp http.ResponseWriter, req *http.Request) {
	printIntro(req)
	printHeader(req.Header)
	printBody(req.Body)
}

func printIntro(req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL.Path)
}

func printHeader(header http.Header) {
	if len(header) > 0 {
		fmt.Println()
		for name, values := range header {
			for _, value := range values {
				fmt.Printf("%v: %v\n", name, value)
			}
		}
	}
}

func printBody(body io.ReadCloser) {
	if body != nil {
		defer body.Close()
		fmt.Println()
		if _, err := io.Copy(os.Stdout, body); err != nil {
			log.Println(err)
		}
	}
}
