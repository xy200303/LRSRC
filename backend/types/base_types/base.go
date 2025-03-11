package base_types

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"time"
	"xiaoyun/backend/validate"
)

func init() {
	validate.SetupValidate()
}

// BaseType 基础表
type BaseType struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BasePage 页表
type BasePage struct {
	Page     int `form:"page" binding:"required,min=1" json:"page"`
	PageSize int `form:"page_size" binding:"required,min=1" json:"page_size"`
}

// DeleteStringIds 删除字符串
type DeleteStringIds struct {
	Ids []string `json:"ids" binding:"required"`
}
type DeleteIntIds struct {
	Ids []int64 `json:"ids" binding:"required"`
}

// TreeData 级联选择器数据
type TreeData struct {
	Label    string      `json:"label"`
	Value    uint64      `json:"value"`
	Children []*TreeData `json:"children"`
}

// StringList 自定义数据类型
type StringList []string

// Scan 实现了 sql.Scanner 接口，用于从数据库读取数据时将字符串转换为 StringArray
func (t *StringList) Scan(value interface{}) error {
	if value == nil {
		*t = nil
		return nil
	}
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

// Value 实现了 driver.Valuer 接口，用于将 StringArray 转换为数据库可接受的格式
func (t StringList) Value() (driver.Value, error) {
	if len(t) == 0 {
		return nil, nil
	}
	marshal, err := json.Marshal(t)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return marshal, nil
}
