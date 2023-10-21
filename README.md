# Go Minimax
The [Minimax](https://api.minimax.chat) unofficial Go library.

## Installation
```bash
go get github.com/Twacqwq/go-minimax
```
go-minimax requires Go version 1.18 or greater.

## Features
- [x] Chatcompletion pro
- [x] Chatcompletion
- [ ] T2A
- [ ] T2A pro
- [ ] Embeddings

## Usage

### Minimax ChatCompletion Example:

```go
package main

import (
	"context"
	"fmt"
	
	minimax "github.com/Twacqwq/go-minimax"
)

func main() {
    client := minimax.NewClient("your token", "your group id")
	resp, err := client.CreateCompletion(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "Please introduce yourself.",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", resp.Reply)
}

```

### Minimax ChatCompletionStream Example:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	minimax "github.com/Twacqwq/go-minimax"
)

func main() {
    client := minimax.NewClient("your token", "your group id")
	stream, err := client.CreateCompletionStream(context.Background(), &minimax.ChatCompletionRequest{
		Model:            minimax.Abab5Dot5,
		TokensToGenerate: 1024,
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
				SenderName: "Twac",
				Text:       "Please introduce yourself.",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%#v\n", resp)
	}
}

```