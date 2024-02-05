package minimax_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestCreateRun(t *testing.T) {
	resp, err := client.CreateRun(context.Background(), &minimax.RunCreateRequest{
		ThreadId:    "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
		AssistantId: "asst_7b4015a894834de9b0122e0119704982",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestRetrieveRun(t *testing.T) {
	resp, err := client.RetrieveRun(context.Background(), &minimax.RunRetrieveRequest{
		ThreadId: "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
		RunId:    "run_4be5a20ef9c5428385bab58829523606",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestListRun(t *testing.T) {
	resp, err := client.ListRun(context.Background(), &minimax.ListRunOption{
		ThreadId: "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestSubmitRun(t *testing.T) {
	resp, err := client.SumbmitToolOutputsRun(context.Background(), &minimax.SubmitRequest{
		ThreadId:   "thread_43c9174103ad4b1ba7c4ff1be5ce8e30",
		RunId:      "run_4be5a20ef9c5428385bab58829523606",
		ToolOutput: nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
