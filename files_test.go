package minimax_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Twacqwq/go-minimax"
)

func TestCreateFile(t *testing.T) {
	res, err := client.CreateFile(context.Background(), &minimax.FileRequest{
		Purpose:  minimax.Retrieval,
		FilePath: "test.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestListFiles(t *testing.T) {
	res, err := client.ListFiles(context.Background(), minimax.Retrieval)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res.Files {
		fmt.Printf("%+v\n", v)
	}
}

func TestRetrieveFile(t *testing.T) {
	res, err := client.RetrieveFile(context.Background(), -1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", res.File)
}

func TestDeleteFile(t *testing.T) {
	res, err := client.DeleteFile(context.Background(), &minimax.DeleteFileRequest{
		FileId: -1,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
