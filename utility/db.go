package utility

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sampleAppDemo/entity"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	con, _ := db.DB()
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(time.Minute * 10)
	DB = db
}

func MigrateDB() {
	err := DB.AutoMigrate(&entity.Item{}, &entity.User{})
	if err != nil {
		panic(err)
	}
}
