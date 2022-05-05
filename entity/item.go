package entity

import "time"

type Item struct {
	Id          int64     `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"required" gorm:"type:varchar(100);UNIQUE"`
	Description string    `json:"description" gorm:"type:varchar(200)"`
	IsActive    bool      `json:"is_active" binding:"required" gorm:"type:boolean"`
	Author      User      `json:"author" binding:"required" gorm:"foreignkey:UserId"`
	UserId      int64     `json:"-"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
