package adminService

import "go-gin/models"

type Admin struct {
	AdminID   int
	AdminName string
	Password  string
	State     int
}

func (a *Admin) Get() (*models.Admin, error) {

	admin, err := models.GetAdminByAdminID(a.AdminID)
	if err != nil {
		return nil, err
	}

	return admin, nil
}
