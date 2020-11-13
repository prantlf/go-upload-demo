package buffered

import (
	"bytes"
	"mime/multipart"

	helpers "github.com/prantlf/go-multipart-helpers"
)

func NewRequestBody(comment, filePath string) (*bytes.Buffer, string, error) {
	reqBody := &bytes.Buffer{}
	formWriter := multipart.NewWriter(reqBody)
	if err := writeRequestBody(formWriter, comment, filePath); err != nil {
		return nil, "", err
	}
	return reqBody, formWriter.FormDataContentType(), nil
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
