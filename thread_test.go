package minimax_test

import (
	"context"
	"fmt"
	"testing"
)

func TestCreateThreads(t *testing.T) {
	resp, err := client.CreateThreads(context.Background(), map[string]string{
		"test1": "test1",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestRetrieveThreads(t *testing.T) {
	resp, err := client.RetrieveThreads(context.Background(), "thread_b19bf15e746a4415bad50c42d4226443")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)
}
