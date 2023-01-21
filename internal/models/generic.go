package models

import "gorm.io/gorm"

func GInsert[T any](db *gorm.DB, t *T) error {
	return db.Model(t).Create(t).Error
}
