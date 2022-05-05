package entity

type User struct {
	Id           int64  `json:"id" gorm:"primary_key;auto_increment"`
	FirstName    string `json:"first_name" binding:"required" gorm:"type:varchar(100);"`
	LastName     string `json:"last_name" binding:"required" gorm:"type:varchar(100);"`
	Email        string `json:"email" binding:"required,email" gorm:"type:varchar(100);UNIQUE"`
	PasswordHash string `json:"password,omitempty" gorm:"type:varchar(1024);"`
}
