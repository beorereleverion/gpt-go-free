package main

import (
	"fmt"

	"github.com/beorereleverion/gpt-go-free/providers"
	"github.com/sirupsen/logrus"
)

func main() {
	dai := providers.NewProvider(providers.ProviderOpchatgpts)
	stream, err := dai.NewCompletion([]providers.Message{
		{
			Role:    "user",
			Content: "Hello world",
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(*stream)
}
