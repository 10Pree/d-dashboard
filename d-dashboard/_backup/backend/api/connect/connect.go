package connect

import (
	"d-dashboard/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DbInstance คือโครงสร้างที่ใช้เก็บตัวแปรฐานข้อมูล
type DbInstance struct {
	Db *gorm.DB
}

// Database เป็นตัวแปรที่ใช้เก็บการเชื่อมต่อฐานข้อมูล
var Database DbInstance

func ConnectDb() {

	dsn := "root:@tcp(127.0.0.1:3306)/db_dashboard?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		os.Exit(2)
	}

	// กรณีที่เชื่อมต่อฐานข้อมูลสำเร็จ
	log.Println("Connected to database successfully")

	// เซ็ต Logger เพื่อบันทึก log การทำงานของฐานข้อมูล
	db.Logger = logger.Default.LogMode(logger.Info)

	// ทำการ Migration เพื่อสร้างตารางที่กำหนดไว้ล่วงหน้า
	log.Println("Running Migrations")
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	// กำหนดค่าให้ตัวแปร Database ที่เก็บเป็น DbInstance ที่มีฐานข้อมูลที่เชื่อมต่อแล้ว
	Database = DbInstance{Db: db}
}
