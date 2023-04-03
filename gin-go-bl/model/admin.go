package model

import (
	"gin-go-bl/utils"
	"gorm.io/gorm"
)

type Admin struct {
	Id       uint   `gorm:"int;not null;primary_key" json:"id"`
	Username string `gorm:"varchar(20);not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func (Admin) TableName() string {
	return "admin"
}

func CreatAdmin(data *Admin) int {
	data.Password = utils.GetMd5(data.Password)
	err := DB.Create(&data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}
func CheckAdminLogin(name string, password string) (code int) {
	var admin Admin
	DB.Where("username = ?", name).First(&admin)
	if admin.Id == 0 {
		code = 400
		return
	}
	if utils.GetMd5(password) != admin.Password {
		code = 400
		return
	}
	code = utils.SUCCESS
	return
}
func GetAdminId(name string) Admin {
	var admin Admin
	DB.Where("username =?", name).First(&admin)
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin
	}
	return admin
}
