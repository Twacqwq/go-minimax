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

var client = minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))

func TestCreateCompletion(t *testing.T) {
	resp, err := client.CreateCompletion(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		RoleMeta: &minimax.RoleMeta{
			UserName: "我",
			BotName:  "专家",
		},
		Prompt: "你是一个擅长发现故事中蕴含道理的专家，你很善于基于我给定的故事发现其中蕴含的道理。",
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				Text:       "我给定的故事：从前，在森林里有只叫聪聪的小猪，他既勤劳，又乐于助人，小动物们都很喜欢他。有一次，小兔子放风筝不小心将风筝挂在了树上，那是小兔子最喜欢的东西呀!",
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
	stream, err := client.CreateCompletionStream(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		Prompt:           "你是一名PDF专家",
		RoleMeta: &minimax.RoleMeta{
			UserName: "我",
			BotName:  "机器人",
		},
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
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

func TestCreateCompletionPro(t *testing.T) {
	res, err := client.CreateCompletionPro(context.Background(), &minimax.ChatCompletionProRequest{
		Model:            minimax.Abab6,
		TokensToGenerate: 1024,
		Messages: []minimax.ProMessage{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "请介绍一下你自己",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", res)
}

func TestCreateCompletionProStream(t *testing.T) {
	stream, err := client.CreateCompletionProStream(context.Background(), &minimax.ChatCompletionProRequest{
		Model:            minimax.Abab6,
		TokensToGenerate: 1024,
		Messages: []minimax.ProMessage{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "请介绍一下你自己",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
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
