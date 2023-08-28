package providers

type provider struct {
	url         string
	active      bool
	canStream   bool
	gpt4Support bool
}

func (p *provider) CanStream() bool {
	return p.canStream
}

func (p *provider) Active() bool {
	return p.active
}

func (p *provider) GPT4Support() bool {
	return p.gpt4Support
}

type Provider interface {
	Active() bool
	CanStream() bool
	NewCompletionStream(messages []Message) (chan string, error)
	NewCompletion(messages []Message) (*string, error)
}

type ProviderType uint8

const (
	ProviderTypeDeepAI    ProviderType = 1
	ProviderTypeChatGPTAI ProviderType = 2
	ProviderOpchatgpts    ProviderType = 3
)

func NewProvider(t ProviderType) Provider {
	switch t {
	case ProviderTypeDeepAI:
		return &deepAIprovider
	case ProviderTypeChatGPTAI:
		return &chatGPTAIProvider
	case ProviderOpchatgpts:
		return &opchatgptsAIprovider
	}
	return nil
}
