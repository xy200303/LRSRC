package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 获取去掉年份部分的时间戳（以毫秒为单位）
func getTimestampWithoutYear(t time.Time) int64 {
	baseTime := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	// 计算目标时间与基准时间之间的毫秒差
	duration := t.Sub(baseTime)
	return int64(duration.Seconds())
}
func GenerateLRID(pre string) string {
	t := time.Now()
	second := getTimestampWithoutYear(t)
	return fmt.Sprintf("%s-%d-%02d%d%04d", pre, t.Year(), t.Month(), second, rand.Intn(10000))
}
