package mysql

import (
	"d-dashboard/connect"
	"d-dashboard/entities"
)

func CreateUser(data *entities.User) error {
	db := connect.db // เชื่อมต่อกับฐานข้อมูลด้วย GORM

	// ใช้เมธอด Create ของ GORM เพื่อเพิ่มข้อมูลลงในฐานข้อมูล
	result := db.Create(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
