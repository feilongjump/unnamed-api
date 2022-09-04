package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cast"
)

type Lang struct {
	En string `json:"en"`
	Zh string `json:"zh"`
}

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (base BaseModel) GetStringID() string {
	return cast.ToString(base.ID)
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (l *Lang) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, &l)
}

// Value 实现 driver.Valuer 接口，Value 返回 json value
func (l Lang) Value() (driver.Value, error) {
	return json.Marshal(l)
}
