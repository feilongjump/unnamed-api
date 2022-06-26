package models

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// // CommonTimestampsField 时间戳
// type CommonTimestampsField struct {
//     CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
//     UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
// }

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (base BaseModel) GetStringID() string {
	return cast.ToString(base.ID)
}
