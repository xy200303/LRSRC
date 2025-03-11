package ai_service

import (
	"github.com/sashabaranov/go-openai"
	"xiaoyun/backend/utils/llm"
)

func CommonChat(text string, systemPrompt string) (string, error) {
	completion, err := llm.ChatCompletion(text, systemPrompt)
	if err != nil {
		return "", nil
	}
	return completion, err
}

func SummaryContent(text string) (string, error) {
	completion, err := llm.ChatCompletion(text, llm.PROMPT_SYSTEM)
	if err != nil {
		return "", nil
	}
	return completion, err
}

func CommonChatStream(req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	steam, err := llm.ChatCompletionSteam(req)
	if err != nil {
		return nil, err
	}
	return steam, nil
}
