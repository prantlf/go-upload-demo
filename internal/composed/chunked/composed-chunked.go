package chunked

import (
	"io"

	composer "github.com/prantlf/go-multipart-composer"
)

func NewRequestBody(comment, filePath string) (io.ReadCloser, string, error) {
	comp := composer.NewComposer()
	comp.AddField("comment", comment)
	if err := comp.AddFile("file", filePath); err != nil {
		comp.Close()
		return nil, "", err
	}
	return comp.DetachReader(), comp.FormDataContentType(), nil
}
