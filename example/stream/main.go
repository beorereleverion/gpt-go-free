package main

import (
	"fmt"

	"github.com/beorereleverion/gpt-go-free/providers"
	"github.com/sirupsen/logrus"
)

func main() {
	dai := providers.NewProvider(providers.ProviderTypeDeepAI)
	stream, err := dai.NewCompletionStream([]providers.Message{
		{
			Role:    "user",
			Content: "Hello world",
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		str, ok := <-stream
		if !ok {
			break
		}
		fmt.Printf("%s", str)
	}
}
