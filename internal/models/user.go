package models

import "gorm.io/gorm"

type User struct {
	conn *gorm.DB `gorm:"-"`

	ID       int32 `gorm:"primaryKey"`
	UserName string
	Password string
}

func NewUser(conn *gorm.DB) *User {
	return &User{
		conn: conn,
	}
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Connection() *gorm.DB {
	return u.conn.Table(u.TableName())
}

func (u *User) Create(data *User) (user *User, err error) {
	err = u.Connection().Create(&data).Error
	return data, err
}

func (u *User) Delete(id string) (user *User, err error) {
	err = u.Connection().Where("id = ?", id).Delete(&user).Error
	return user, err
}

func (u *User) Update(id string, data *User) (user *User, err error) {
	err = u.Connection().Where("id = ?", id).Updates(&data).Error
	return user, err
}

func (u *User) Find(id string) (user *User, err error) {
	err = u.Connection().Where("id = ?", id).First(&user).Error
	return user, err
}
