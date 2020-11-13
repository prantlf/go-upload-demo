package sized

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	helpers "github.com/prantlf/go-multipart-helpers"
)

func NewRequestBody(comment, filePath string) (io.Reader, string, int64, error) {
	pipeReader, pipeWriter := io.Pipe()
	formWriter := multipart.NewWriter(pipeWriter)
	go func() {
		defer pipeWriter.Close()
		if err := writeRequestBody(formWriter, comment, filePath); err != nil {
			pipeWriter.CloseWithError(err)
		}
	}()
	contentLength, err := measureRequestBody(comment, filePath)
	if err != nil {
		return nil, "", 0, err
	}
	return pipeReader, formWriter.FormDataContentType(), contentLength, nil
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

func measureRequestBody(comment, filePath string) (int64, error) {
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)
	if err := writer.WriteField("comment", comment); err != nil {
		return 0, err
	}
	if _, err := helpers.CreateFilePart(writer, "file", filepath.Base(filePath)); err != nil {
		return 0, err
	}
	if err := writer.Close(); err != nil {
		return 0, err
	}
	stat, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return int64(message.Len()) + stat.Size(), nil
}
