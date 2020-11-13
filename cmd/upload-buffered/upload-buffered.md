# upload-buffered

The [code] is simple, but bigger files may exceed the free heap.

Condensed implementation without error handling, see also the [actual implementation]:

```go
// prepare the multipart message
reqBody := &bytes.Buffer{}
formWriter := multipart.NewWriter(reqBody)
err := formWriter.WriteField("comment", "a comment")
err := helpers.WriteFile(formWriter, "file", "test.txt")
err := formWriter.Close()
contentType := formWriter.FormDataContentType()

// post the request
resp, err := http.DefaultClient.Post(
  "http://localhost:8182/", contentType, reqBody)
```

[code]: main.go
[actual implementation]: ../../internal/buffered/buffered.go
