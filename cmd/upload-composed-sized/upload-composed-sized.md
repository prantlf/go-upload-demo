# upload-composed-sized

The [code] is simple and optimal.

Condensed implementation without error handling, see also the [actual implementation]:

```go
// prepare the multipart message
comp := composer.NewComposer()
comp.AddField("comment", "a comment")
err := comp.AddFile("file", "test.txt")
contentType := comp.FormDataContentType()
reqBody, contentLength, err := comp.DetachReaderWithSize()

// post the request
req, err := http.NewRequest("POST", "http://localhost:8182/", reqBody)
req.Header.Add("Content-Type", contentType)
req.ContentLength = contentLength
resp, err := http.DefaultClient.Do(req)
```

[code]: main.go
[actual implementation]: ../../internal/composed/sized/composed-sized.go
