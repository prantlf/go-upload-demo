package sized

import (
	"io"

	composer "github.com/prantlf/go-multipart-composer"
)

func NewRequestBody(comment, filePath string) (io.ReadCloser, string, int64, error) {
	comp := composer.NewComposer()
	defer comp.Close()
	comp.AddField("comment", comment)
	if err := comp.AddFile("file", filePath); err != nil {
		return nil, "", 0, err
	}
	reqBody, contentLength, err := comp.DetachReaderWithSize()
	if err != nil {
		return nil, "", 0, err
	}
	return reqBody, comp.FormDataContentType(), contentLength, nil
}
