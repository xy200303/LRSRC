package system_types

import "time"

// File 文件信息
type File struct {
	FileID    string    `gorm:"primaryKey" json:"file_id"`
	Extension string    `json:"extension"`
	FileType  string    `json:"file_type"` //1本地 2阿里云 3腾讯云 4华为云
	FileName  string    `json:"file_name"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
