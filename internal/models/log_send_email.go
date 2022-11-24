package models

import "gorm.io/gorm"

type LogSendEmail struct {
	CreatorBase
	Feature       string `json:"feature" gorm:"column:feature;comment: 功能模块"`
	OperationType string `json:"operationType" gorm:"column:operation_type;comment:操作类型"`
	Description   string `json:"description" gorm:"column:description;comment:日志描述"`
	Result        bool   `json:"result" gorm:"column:result;comment:操作结果"`
}

func (t *LogSendEmail) TableName() string {
	return "log_send_emails"
}

func (t *LogSendEmail) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}
