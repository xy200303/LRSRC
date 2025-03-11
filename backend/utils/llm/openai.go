package llm

import (
	"context"
	"fmt"
	"github.com/pkoukk/tiktoken-go"
	"github.com/sashabaranov/go-openai"
	"log"
	"xiaoyun/backend/types/system_types"
)

var (
	client       *openai.Client
	SysConfigMap *system_types.SysConfigMap
)

func InitOpenAi(sysConfigMap *system_types.SysConfigMap) {
	SysConfigMap = sysConfigMap
	// 创建自定义配置
	clientConfig := openai.DefaultConfig(SysConfigMap.SysAiApiKey)
	clientConfig.BaseURL = SysConfigMap.SysAiBaseUrl // 设置自定义 API Base URL
	// 初始化客户端
	client = openai.NewClientWithConfig(clientConfig)
}

func ChatCompletion(text string, sysPrompt string) (string, error) {
	messages := make([]openai.ChatCompletionMessage, 0)
	//初始化角色
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: sysPrompt,
	})
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    SysConfigMap.SysAiModel,
			Messages: messages,
		},
	)
	if err != nil {
		fmt.Println(SysConfigMap)
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}
	content := resp.Choices[0].Message.Content
	return content, nil
}

func ChatCompletionSteam(req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	// 发送请求并处理流式响应
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Println(SysConfigMap)
		log.Printf("Failed to create chat completion stream: %v\n", err)
		return nil, err
	}
	return stream, nil
}
func CountTextTokens(text string) (int, error) {
	// 获取指定模型的编码器
	encoding, err := tiktoken.EncodingForModel(SysConfigMap.SysAiModel)
	if err != nil {
		return 0, fmt.Errorf("failed to get encoding for model %s: %v", SysConfigMap.SysAiModel, err)
	}
	// 将文本编码为 token
	tokens := encoding.Encode(text, nil, nil)
	// 返回 token 数量
	return len(tokens), nil
}

func CountMessageTokens(messages []openai.ChatCompletionMessage) (int, error) {
	// 获取指定模型的编码器
	encoding, err := tiktoken.EncodingForModel(SysConfigMap.SysAiModel)
	if err != nil {
		return 0, fmt.Errorf("failed to get encoding for model %s: %v", SysConfigMap.SysAiModel, err)
	}
	// 将每条消息拼接为字符串
	var totalTokens int
	for _, message := range messages {
		// 将角色和内容拼接为一条消息
		text := fmt.Sprintf("%s: %s", message.Role, message.Content)
		// 编码为 token 并累加
		tokens := encoding.Encode(text, nil, nil)
		totalTokens += len(tokens)
	}
	return totalTokens, nil
}
