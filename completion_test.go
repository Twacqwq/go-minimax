package minimax_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestCreateCompletion(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	resp, err := client.CreateCompletion(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "你是谁?",
			},
		},
	})
	if err != nil {
		t.Log(err)
		return
	}

	fmt.Printf("%#v\n", resp)
}

func TestCreateCompletionStream(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	stream, err := client.CreateCompletionStream(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "请介绍一下你自己",
			},
		},
	})
	if err != nil {
		t.Log(err)
		return
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			t.Log(err.Error())
			return
		}
		fmt.Printf("%#v\n", resp)
	}
}
