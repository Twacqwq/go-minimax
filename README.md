# Go MiniMax
[![Go Reference](https://pkg.go.dev/badge/github.com/Twacqwq/go-minimax.svg)](https://pkg.go.dev/github.com/Twacqwq/go-minimax)
[![Go Report Card](https://goreportcard.com/badge/github.com/Twacqwq/go-minimax)](https://goreportcard.com/report/github.com/Twacqwq/go-minimax)

The [MiniMax](https://api.minimax.chat) unofficial Go library.

## Installation
```bash
go get github.com/Twacqwq/go-minimax
```
go-minimax requires Go version 1.18 or greater.

## Features
- [x] Chatcompletion pro
- [x] Chatcompletion
- [x] Embeddings
- [x] T2A
- [x] T2A pro
- [ ] T2A large
- [ ] T2A Stream
- [ ] Voice Cloning
- [x] Assistants API
- [x] Thread
- [x] Run
- [x] Message
- [x] Files
- [ ] Retrieval
- [ ] Finetune API
- [ ] Role Classification
- [ ] Role Audio Generation

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
		RoleMeta: &minimax.RoleMeta{
			UserName: "我",
			BotName:  "专家",
		},
		Prompt: "你是一个擅长发现故事中蕴含道理的专家，你很善于基于我给定的故事发现其中蕴含的道理。",
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
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
		RoleMeta: &minimax.RoleMeta{
			UserName: "我",
			BotName:  "专家",
		},
		Prompt: "你是一个擅长发现故事中蕴含道理的专家，你很善于基于我给定的故事发现其中蕴含的道理。",
		Messages: []minimax.Message{
			{
				SenderType: minimax.ChatMessageRoleUser,
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

### Minimax ChatCompletionPro Example:

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

```

### Minimax ChatCompletionProStream Example:

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

### Minimax CreateEmbeddings Example:

```go
package main

import (
	"context"
	"fmt"
	
	minimax "github.com/Twacqwq/go-minimax"
)

func main() {
    client := minimax.NewClient("your token", "your group id")
	resp, err := client.CreateEmbeddings(context.Background(), &minimax.CreateEmbeddingsRequest{
		Texts: []string{"hello"},
		// Type: minimax.EmbeddingsDbType,
		Type:  minimax.EmbeddingsQueryType,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", resp)
}

```

### Minimax CreateTextToSpeech Example:

```go
package main

import (
	"context"
	"fmt"

	minimax "github.com/Twacqwq/go-minimax"
)

func main() {
	client := minimax.NewClient("your token", "your group id")
	resp, err := client.CreateTextToSpeech(context.Background(), &minimax.CreateT2ARequest{
		Text:    "hello",
		VoiceId: "female-yujie",
		Path:    "./",
		Name:    "hello.mp3",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)
}

```

### Minimax CreateTextToSpeechPro Example:

```go
package main

import (
	"context"
	"fmt"

	minimax "github.com/Twacqwq/go-minimax"
)

func main() {
	client := minimax.NewClient("your token", "your group id")
	resp, err := client.CreateTextToSpeechPro(context.Background(), &minimax.CreateT2ARequest{
		Text:    "hello",
		VoiceId: "female-yujie",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)
}

```
