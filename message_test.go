package minimax_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestListMessages(t *testing.T) {
	resp, err := client.ListMessages(context.Background(), &minimax.ListMessageOption{
		ThreadId: "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range resp.Data {
		fmt.Printf("%v: %v\n", v.Role, v.Content[0].Text)
	}
}

func TestCreateMessages(t *testing.T) {
	resp, err := client.CreateMessages(context.Background(), &minimax.MessageCreateRequest{
		ThreadId: "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
		Role:     "user",
		Content:  "你好",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestRetrieveMessages(t *testing.T) {
	resp, err := client.RetrieveMessages(context.Background(), &minimax.MessageRetrieveRequest{
		MessageId: "msg_5f871f6eca564ec294f7e931a70ca259",
		ThreadId:  "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp.AsstMessage)
}
