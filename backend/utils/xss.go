package utils

import (
	"github.com/microcosm-cc/bluemonday"
)

// GetSafeContent 获取安全的文本内容
func GetSafeContent(content string) string {
	p := bluemonday.UGCPolicy()
	safeHTML := p.Sanitize(content)
	return safeHTML
}
