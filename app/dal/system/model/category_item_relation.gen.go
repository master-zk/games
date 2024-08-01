// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCategoryItemRelation = "category_item_relation"

// CategoryItemRelation 分类-内容-关联关系
type CategoryItemRelation struct {
	ID         uint64         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CategoryID int            `gorm:"column:category_id;not null;comment:分类ID: category.id" json:"category_id"` // 分类ID: category.id
	ItemType   int            `gorm:"column:item_type;not null;comment:内容类型: 1=游戏，2=文章，" json:"item_type"`      // 内容类型: 1=游戏，2=文章，
	ItemID     int            `gorm:"column:item_id;not null;comment:内容ID" json:"item_id"`                      // 内容ID
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName CategoryItemRelation's table name
func (*CategoryItemRelation) TableName() string {
	return TableNameCategoryItemRelation
}