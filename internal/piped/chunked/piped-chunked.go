package chunked

import (
	"io"
	"mime/multipart"

	helpers "github.com/prantlf/go-multipart-helpers"
)

func NewRequestBody(comment, filePath string) (io.Reader, string) {
	pipeReader, pipeWriter := io.Pipe()
	formWriter := multipart.NewWriter(pipeWriter)
	go func() {
		defer pipeWriter.Close()
		if err := writeRequestBody(formWriter, comment, filePath); err != nil {
			pipeWriter.CloseWithError(err)
		}
	}()
	return pipeReader, formWriter.FormDataContentType()
}

func writeRequestBody(writer *multipart.Writer, comment, filePath string) error {
	if err := writer.WriteField("comment", comment); err != nil {
		return err
	}
	if err := helpers.WriteFile(writer, "file", filePath); err != nil {
		return err
	}
	return writer.Close()
}
