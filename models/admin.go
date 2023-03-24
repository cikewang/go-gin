package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	AdminID int `gorm:"primary_key" json:"admin_id"`
	Model
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
	State     int8   `json:"state"`
}

// GetAdminByAdminID 根据 admin_id 查询 admin 信息
func GetAdminByAdminID(adminId int) (*Admin, error) {
	var admin = new(Admin)
	err := db.Where("admin_id = ?", adminId).First(admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin, err
	}

	if admin.AdminID > 0 {
		return admin, nil
	}

	return admin, nil
}
