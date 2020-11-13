# upload-composed-chunked

The [code] is very simple, but some servers may not support chunked encoding and fail with such requests.

Condensed implementation without error handling, see also the [actual implementation]:

```go
// prepare the multipart message
comp := composer.NewComposer()
comp.AddField("comment", "a comment")
err := comp.AddFile("file", "test.txt")
contentType := comp.FormDataContentType()
reqBody := comp.DetachReader()

// post the request
resp, err := http.DefaultClient.Post(
  "http://localhost:8182/", contentType, reqBody)
```

[code]: main.go
[actual implementation]: ../../internal/composed/chunked/composed-chunked.go
