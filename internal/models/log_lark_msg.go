package models

import "gorm.io/gorm"

const (
	LarkMsgTypeText = iota
	LarkMsgTypePost
)

type LogLarkMsg struct {
	gorm.Model
	Type     int    `json:"type" gorm:"column:type;comment:消息类型 0 文本"`
	Openids  string `json:"openids" gorm:"column:openids;comment:账号"`
	UnionIds string `json:"unionIds" gorm:"column:union_ids;comment:用户 union ids"`
	Content  string `json:"content" gorm:"column:content;comment:发送内容"`
}
