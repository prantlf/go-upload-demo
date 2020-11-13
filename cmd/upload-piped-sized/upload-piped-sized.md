# upload-piped-sized

The [code] is optimal, but complicated. The [actual implementation] uses internal methods with disabled file content to be able to measure the size of the rest of the request body withtout duplicating the code as it is in this sketch here.

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
contentLength, err := measureRequestBody()

// post the request
req, err := http.NewRequest("POST", "http://localhost:8182/", reqBody)
req.Header.Add("Content-Type", contentType)
req.ContentLength = contentLength
resp, err := http.DefaultClient.Do(req)

// compute the content length
func measureRequestBody() (int64, error) {
  message := &bytes.Buffer{}
  writer := multipart.NewWriter(message)
  err := writer.WriteField("comment", "a comment")
  err := helpers.CreateFilePart(writer, "file", "test.txt")
  writer.Close()
  stat, err := os.Stat("test.txt")
  return int64(message.Len()) + stat.Size(), err
}
```

[code]: main.go
[actual implementation]: ../../internal/piped/sized/piped-sized.go
