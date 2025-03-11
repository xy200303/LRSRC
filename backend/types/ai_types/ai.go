package ai_types

type SummaryContentReq struct {
	Content string `json:"content"`
}

type CommonChat struct {
	Text         string `json:"text"`
	SystemPrompt string `json:"system_prompt"`
}
