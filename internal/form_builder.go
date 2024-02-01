package minimax

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

var errFileNameEmpty = errors.New("file name must be not empty")

type FormDataBuilder interface {
	// AddField is add form-data field
	AddField(key, value string) error
	// CreateFormFile create form-data multipart file
	CreateFormFile(fieldName string, file *os.File) error
	// Close .
	Close() error
	// FormDataContentType
	FormDataContentType() string
}

type CreateFormDataBuilderFunc func(io.Writer) FormDataBuilder

type DefaultFormDataBuilder struct {
	writer *multipart.Writer
}

func NewDefaultFormDataBuilder(w io.Writer) FormDataBuilder {
	return &DefaultFormDataBuilder{
		writer: multipart.NewWriter(w),
	}
}

func (df *DefaultFormDataBuilder) CreateFormFile(fieldName string, file *os.File) error {
	return df.createFileAddForm(fieldName, file.Name(), file)
}

func (df *DefaultFormDataBuilder) AddField(key, value string) error {
	return df.writer.WriteField(key, value)
}

func (df *DefaultFormDataBuilder) Close() error {
	return df.writer.Close()
}

func (df *DefaultFormDataBuilder) FormDataContentType() string {
	return df.writer.FormDataContentType()
}

func (df *DefaultFormDataBuilder) createFileAddForm(fieldName, fileName string, r io.Reader) error {
	if fileName == "" {
		return errFileNameEmpty
	}

	w, err := df.writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	return err
}

var _ FormDataBuilder = (*DefaultFormDataBuilder)(nil)
