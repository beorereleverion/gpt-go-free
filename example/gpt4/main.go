package main

import (
	"fmt"

	"github.com/beorereleverion/gpt-go-free/providers"
	"github.com/sirupsen/logrus"
)

func main() {
	dai := providers.NewProvider(providers.ProviderTypeChatGPTAI)
	stream, err := dai.NewCompletion([]providers.Message{
		{
			Role:    "user",
			Content: "какую версию gpt ты используешь?",
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(*stream)

}
