package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type chatGPTAI struct {
	provider
}

var chatGPTAIProvider = chatGPTAI{
	provider: provider{
		url:         "https://chatgpt.ai/gpt-4/",
		active:      true,
		gpt4Support: true,
	},
}

func (c *chatGPTAI) NewCompletion(messages []Message) (*string, error) {
	chat := ""
	for _, message := range messages {
		chat += fmt.Sprintf("%s: %s\n", message.Role, message.Content)
	}
	chat += "assistant: "

	resp, err := http.Get("https://chatgpt.ai/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`data-nonce="(.*)"\n     data-post-id="(.*)"\n     data-url="(.*)"\n     data-bot-id="(.*)"\n     data-width`)
	matches := re.FindStringSubmatch(string(body))
	nonce := matches[1]
	postId := matches[2]
	botId := matches[4]
	headers := map[string]string{
		"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		"accept":             "*/*",
		"Connection":         "keep-alive",
		"authority":          "chatgpt.ai",
		"accept-language":    "en,fr-FR;q=0.9,fr;q=0.8,es-ES;q=0.7,es;q=0.6,en-US;q=0.5,am;q=0.4,de;q=0.3",
		"cache-control":      "no-cache",
		"origin":             "https://chatgpt.ai",
		"pragma":             "no-cache",
		"referer":            "https://chatgpt.ai/gpt-4/",
		"sec-ch-ua":          `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
		"sec-ch-ua-mobile":   "?0",
		"sec-ch-ua-platform": `"Windows"`,
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"Content-Length":     "163",
		"Content-Type":       "application/x-www-form-urlencoded",
	}

	data := map[string]string{
		"_wpnonce": nonce,
		"post_id":  postId,
		"url":      "https://chatgpt.ai/gpt-4",
		"action":   "wpaicg_chat_shortcode_message",
		"message":  chat,
		"bot_id":   botId,
	}

	payload := url.Values{}
	for key, value := range data {
		payload.Add(key, value)
	}

	req, err := http.NewRequest("POST", "https://chatgpt.ai/wp-admin/admin-ajax.php", strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return parseChatGPRAIResponse(respBody)
}

func (c *chatGPTAI) NewCompletionStream(messages []Message) (chan string, error) {
	return nil, errNoStreamSupport
}

func parseChatGPRAIResponse(response []byte) (*string, error) {
	var responseData map[string]interface{}
	err := json.Unmarshal(response, &responseData)
	if err != nil {
		return nil, err
	}

	dataValue, ok := responseData["data"].(string)
	if !ok {
		return nil, errors.New("Failed to extract 'data' value from JSON response")
	}

	return &dataValue, nil
}
