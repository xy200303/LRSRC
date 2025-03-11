package utils

import "time"

// GetCurrentDate 获取当前日期并按 "YYYY/MM/DD" 格式返回
func GetCurrentDate() string {
	// 获取当前时间
	now := time.Now()
	// 按照 "2006/01/02" 格式化时间（Go 中的特殊日期格式）
	formattedDate := now.Format("2006/01/02")
	return formattedDate
}
