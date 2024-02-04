package minimax_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestListAssistants(t *testing.T) {
	resp, err := client.ListAssistants(context.Background(), &minimax.AssistantListOption{Limit: 1})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestCreateAssistants(t *testing.T) {
	resp, err := client.CreateAssistants(context.Background(), &minimax.AssistantCreateRequest{
		Model:        minimax.Abab5Dot5,
		Name:         "apiass",
		Instructions: "demo",
		Description:  "demo",
		Tools: []*minimax.Tool{
			{Typ: minimax.ToolCodeInterpreter},
			{Typ: minimax.ToolWebSearch},
		},
		FileIds: []string{},
		RoleMeta: &minimax.RoleMeta{
			UserName: "Twac",
			BotName:  "Bot",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestRetrieveAssistants(t *testing.T) {
	resp, err := client.RetrieveAssistants(context.Background(), "asst_f75db16b622e45ff8b028307d77e81ef")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp.Assistant)
}

func TestDeleteAssistants(t *testing.T) {
	resp, err := client.DeleteAssistants(context.Background(), "asst_f75db16b622e45ff8b028307d77e81ef")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestListAssistantFiles(t *testing.T) {
	resp, err := client.ListAssistantFiles(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
