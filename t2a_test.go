package minimax_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestTextToSpeech(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	resp, err := client.CreateTextToSpeech(context.Background(), &minimax.CreateT2ARequest{
		VoiceId: "female-yujie",
		Text:    "hello",
		Path:    "./",
		Name:    "hello.mp3",
		TimberWeights: []minimax.TimberWeight{
			{
				VoiceId: "female-yujie",
				Weight:  1,
			},
		},
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}

func TestTextToSpeechPro(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	resp, err := client.CreateTextToSpeechPro(context.Background(), &minimax.CreateT2ARequest{
		Text:    "hello",
		VoiceId: "female-yujie",
	})
	if err != nil {
		t.Log(err.Error())
		return
	}

	fmt.Printf("%#v\n", resp)
}
