package providers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type opchatgptsAI struct {
	provider
}

var opchatgptsAIprovider = opchatgptsAI{
	provider: provider{
		url:       "https://opchatgpts.net/wp-json/ai-chatbot/v1/chat",
		active:    true,
		canStream: true,
	},
}

func (o *opchatgptsAI) NewCompletion(messages []Message) (*string, error) {
	payload := map[string]interface{}{
		"env":             "chatbot",
		"session":         "N/A",
		"prompt":          "\n",
		"context":         "Converse as if you were an AI assistant. Be friendly, creative.",
		"messages":        messages,
		"newMessage":      messages[len(messages)-1].Content,
		"userName":        `<div class="mwai-name-text">User:</div>`,
		"aiName":          `<div class="mwai-name-text">AI:</div>`,
		"model":           "gpt-3.5-turbo",
		"temperature":     0.8,
		"maxTokens":       1024,
		"maxResults":      1,
		"apiKey":          "",
		"service":         "openai",
		"embeddingsIndex": "",
		"stop":            "",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(o.url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(respBody, &responseMap)
	if err != nil {
		panic(err)
	}

	reply := responseMap["reply"].(string)
	return &reply, nil
}

func (o *opchatgptsAI) NewCompletionStream(messages []Message) (chan string, error) {
	return nil, errNoStreamSupport
}
