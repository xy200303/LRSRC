package ai_controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"strings"
	"xiaoyun/backend/service/ai_service"
	"xiaoyun/backend/service/system_service"
	"xiaoyun/backend/types/ai_types"
	"xiaoyun/backend/utils/resp"
	"xiaoyun/backend/validate"
)

func SummaryContent(c *gin.Context) {
	var req ai_types.SummaryContentReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}

	content, err := ai_service.SummaryContent(req.Content)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, content)
}

// CommonChat 通用聊天
func CommonChat(c *gin.Context) {
	var req ai_types.CommonChat
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	content, err := ai_service.CommonChat(req.Text, req.SystemPrompt)
	if err != nil {
		resp.Err(c, err)
		return
	}
	resp.Ok(c, content)
}

func CommonChatStream(c *gin.Context) {
	var req openai.ChatCompletionRequest
	if err := c.ShouldBind(&req); err != nil {
		resp.Resp(c, 400, validate.GetErrMsg(&req, err), nil)
		return
	}
	req.Model = system_service.SysConfigMap.SysAiModel
	// 设置响应头以支持流式传输
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	stream, err := ai_service.CommonChatStream(req)
	if err != nil {
		log.Println(err)
		resp.Err(c, err)
		return
	}
	defer stream.Close()
	// 逐行读取响应并实时推送
	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Stream finished")
				break
			} else {
				c.String(http.StatusInternalServerError, "Error reading response: %v", err)
				return
			}
		}
		// 确保有内容可发送
		if len(response.Choices) > 0 && response.Choices[0].Delta.Content != "" {
			// 将内容封装为 JSON 格式
			// 转换为 JSON 字符串
			jsonStr, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
			} else {
				// 发送数据到客户端
				fmt.Fprintf(c.Writer, "data: %s\n\n", strings.ReplaceAll(string(jsonStr), "\n", "\\n"))
				c.Writer.Flush()
			}

		}
	}
}
