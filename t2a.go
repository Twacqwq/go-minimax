package minimax

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

type CreateT2ARequest struct {
	Model           string         `json:"model"`
	Text            string         `json:"text"`
	TimberWeights   []TimberWeight `json:"timber_weights,omitempty"`
	VoiceId         string         `json:"voice_id,omitempty"`
	Speed           float32        `json:"speed,omitempty"`
	Vol             float32        `json:"vol,omitempty"`
	Pitch           int            `json:"pitch,omitempty"`
	AudioSampleRate int            `json:"audio_sample_rate,omitempty"`
	Bitrate         int            `json:"bitrate,omitempty"`

	Path string `json:"-"`
	Name string `json:"-"`
}

type CreateT2AResponse struct {
	TraceId      string    `json:"trace_id,omitempty"`
	BaseResp     BaseResp  `json:"base_resp,omitempty"`
	AudioFile    string    `json:"audio_file,omitempty"`
	SubtitleFile string    `json:"subtitle_file,omitempty"`
	ExtraInfo    ExtraInfo `json:"extra_info,omitempty"`
}

func (c *Client) CreateTextToSpeech(ctx context.Context, request *CreateT2ARequest) (*CreateT2AResponse, error) {
	request.Model = Speech01
	if request.AudioSampleRate != 0 || request.Bitrate != 0 {
		return nil, ErrCreateTextToSpeechProNotSupported
	}

	outputFile, err := createFile(request.Path, request.Name)
	if err != nil {
		return nil, err
	}
	defer outputFile.Close()

	req, err := c.newRequest(ctx, c.buildFullURL(speech, request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &CreateT2AResponse{}
	err = c.send(req, resp, outputFile)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateTextToSpeechPro(ctx context.Context, request *CreateT2ARequest) (*CreateT2AResponse, error) {
	request.Model = Speech01
	req, err := c.newRequest(ctx, c.buildFullURL(speechPro, Speech01Pro), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &CreateT2AResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func createFile(path, name string) (*os.File, error) {
	if name == "" {
		name = uuid.Must(uuid.NewV4(), nil).String() + ".mp3"
	}
	return os.Create(filepath.Join(path, name))
}
