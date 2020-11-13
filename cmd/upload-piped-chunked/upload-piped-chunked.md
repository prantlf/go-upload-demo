# upload-piped-chunked

The [code] is fairly simple, but some servers may not support chunked encoding and fail with such requests.

Condensed implementation without error handling, see also the [actual implementation]:

```go
// prepare the multipart message
pipeReader, pipeWriter := io.Pipe()
formWriter := multipart.NewWriter(pipeWriter)
go func() {
  defer pipeWriter.Close()
  formWriter := multipart.NewWriter(reqBody)
  err := formWriter.WriteField("comment", "a comment")
  err := helpers.WriteFile(formWriter, "file", "test.txt")
  err := formWriter.Close()
  pipeWriter.CloseWithError(err)
}()
contentType := formWriter.FormDataContentType()

// post the request
resp, err := http.DefaultClient.Post(
  "http://localhost:8182/", contentType, pipeReader)
```

[code]: main.go
[actual implementation]: ../../internal/piped/chunked/piped-chunked.go
