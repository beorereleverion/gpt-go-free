# gpt-go-free

this is port of Repository [gpt4free](https://github.com/xtekky/gpt4free) Transformed from Python to Go for educational purposes only

By using this repository or any code related to it, you agree to the [legal notice](./LEGAL_NOTICE.md). The author is not responsible for any copies, forks, reuploads made by other users, or anything else related to gpt4free. This is the author's only account and repository. To prevent impersonation or irresponsible actions, please comply with the GNU GPL license this Repository uses.

## Table of Contents:

- [Getting Started](#getting-started)
- [Usage](#usage-examples)

## Getting Started

I've initiated the process of porting active providers and have successfully ported several providers according to the following table (with an added 'Ported' column and notes).

| Website| Provider| gpt-3.5 | gpt-4 | Streaming | Status | Auth | Ported | Notes |
| ------ | ------- | ------- | ----- | --------- | ------ | ---- | ------ | ----- |
| [ai.ls](https://ai.ls) | g4f.provider.Ails | ✔️ | ❌ | ✔️ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [chatgpt.ai](https://chatgpt.ai/gpt-4/) | g4f.provider.ChatgptAi | ❌ | ✔️ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ✔️ | This is not gpt-4 :( |
| [opchatgpts.net](https://opchatgpts.net) | g4f.provider.ChatgptLogin | ✔️ | ❌ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ✔️ | Some parameters has been hadcoded. too slowly|
| [deepai.org](https://deepai.org) | g4f.provider.DeepAi | ✔️ | ❌ | ✔️ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ✔️ | all OK |
| [chat.getgpt.world](https://chat.getgpt.world/) | g4f.provider.GetGpt | ✔️ | ❌ | ✔️ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [gpt-gm.h2o.ai](https://gpt-gm.h2o.ai) | g4f.provider.H2o | ❌ | ❌ | ✔️ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [opchatgpts.net](https://opchatgpts.net) | g4f.provider.Opchatgpts | ✔️ | ❌ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [wewordle.org](https://wewordle.org/) | g4f.provider.Wewordle | ✔️ | ❌ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [you.com](https://you.com) | g4f.provider.You | ✔️ | ❌ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |
| [chat9.yqcloud.top](https://chat9.yqcloud.top/) | g4f.provider.Yqcloud | ✔️ | ❌ | ❌ | ![Active](https://img.shields.io/badge/Active-brightgreen) | ❌ | ❌ | |

## Usage Examples

[another examples](./example)

```go
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
			Content: "Hello World",
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
```
