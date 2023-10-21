package minimax_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestCreateEmbeddingsByDbType(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	resp, err := client.CreateEmbeddings(context.Background(), &minimax.CreateEmbeddingsRequest{
		Texts: []string{"hello"},
		Type:  minimax.EmbeddingsDbType,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("%#v\n", resp)
}

func TestCreateEmbeddingsByQueryType(t *testing.T) {
	client := minimax.NewClient(os.Getenv("TOKEN"), os.Getenv("GROUPID"))
	resp, err := client.CreateEmbeddings(context.Background(), &minimax.CreateEmbeddingsRequest{
		Texts: []string{"hello"},
		Type:  minimax.EmbeddingsQueryType,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("%#v\n", resp)
}
