package connect

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// ข้อมูลสำหรับการเชื่อมต่อกับ MySQL
	dsn := "root:keep1234@tcp(127.0.0.1:3306)/d-dashboard?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
