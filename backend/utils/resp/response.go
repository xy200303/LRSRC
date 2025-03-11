package resp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"xiaoyun/backend/types/base_types"
)

// Ok 响应成功
func Ok(c *gin.Context, data interface{}) {
	c.JSON(int(Success), gin.H{
		"code": Success,
		"msg":  Success.Message(),
		"data": data,
	})
	c.Abort()
}

// Err 响应错误
func Err(c *gin.Context, err error) {
	var msg string
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg = "未找到记录"
	} else {
		msg = err.Error()
	}
	c.JSON(int(Error), gin.H{
		"code": Error,
		"msg":  msg,
		"data": nil,
	})
	c.Abort()
}

// Resp 通用响应
func Resp(c *gin.Context, code ErrorCode, msg string, data interface{}) {
	c.JSON(int(code), gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

// Page 返回分页结果
func Page(c *gin.Context, result interface{}, total int64, page *base_types.BasePage) {
	// 计算总页数
	totalPage := int(math.Ceil(float64(total) / float64(page.PageSize)))
	c.JSON(int(Success), gin.H{
		"code":       Success,
		"msg":        Success.Message(),
		"data":       result,
		"total":      total,
		"page":       page.Page,
		"page_size":  page.PageSize,
		"total_page": totalPage,
	})
	c.Abort()
}
