package models

import "gorm.io/gorm"

type CreatorBase struct {
	gorm.Model
	CreatorId uint `json:"creatorId" gorm:"column:creator_id;comment: 创建人 ID"`
}
