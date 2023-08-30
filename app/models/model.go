package models

import (
	"mumu/pkg/types"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt int64 `gorm:"column:created_at;index"`
	UpdatedAt int64 `gorm:"column:updated_at;index"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
