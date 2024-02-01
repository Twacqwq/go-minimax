package minimax

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
)

type PurposeType string

var (
	Retrieval       PurposeType = "retrieval"
	FineTune        PurposeType = "fine-tune"
	FineTuneResult  PurposeType = "fine-tune-result"
	VoiceClone      PurposeType = "voice_clone"
	Assistants      PurposeType = "assistants"
	RoleRecognition PurposeType = "role-recognition"
)

type FileRequest struct {
	Purpose  PurposeType `json:"purpose"`
	FilePath string      `json:"-"`
}

type FileResponse struct {
	File     *File     `json:"file"`
	BaseResp *BaseResp `json:"base_resp"`
}

type ListFileResponse struct {
	Files    []*File   `json:"files"`
	BaseResp *BaseResp `json:"base_resp"`
}

type RetrieveFileResponse struct {
	File     *File     `json:"file"`
	BaseResp *BaseResp `json:"base_resp"`
}

type DeleteFileRequest struct {
	FileId int64 `json:"file_id"`
}

type DeleteFileResponse struct {
	FileId   int64     `json:"file_id"`
	BaseResp *BaseResp `json:"base_resp"`
}

type File struct {
	FileId    int64  `json:"file_id"`
	Bytes     int64  `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func (c *Client) CreateFile(ctx context.Context, request *FileRequest) (*FileResponse, error) {
	var buf bytes.Buffer
	builder := c.formDataBuilder(&buf)

	file, err := os.Open(request.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// set form-data params
	err = builder.AddField("purpose", string(request.Purpose))
	if err != nil {
		return nil, err
	}
	err = builder.CreateFormFile("file", file)
	if err != nil {
		return nil, err
	}

	err = builder.Close()
	if err != nil {
		return nil, err
	}

	// send request
	req, err := c.newRequest(ctx, c.fullURL("/files/upload"), http.MethodPost, withBody(&buf), withContentType(builder.FormDataContentType()))
	if err != nil {
		return nil, err
	}
	resp := &FileResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListFiles(ctx context.Context, purpose PurposeType) (*ListFileResponse, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("%s&purpose=%s", c.fullURL("/files/list"), purpose), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &ListFileResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveFile(ctx context.Context, fileId int64) (*RetrieveFileResponse, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("%s&file_id=%d", c.fullURL("/files/retrieve"), fileId), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &RetrieveFileResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteFile(ctx context.Context, request *DeleteFileRequest) (*DeleteFileResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/files/delete"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &DeleteFileResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
