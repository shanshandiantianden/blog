package model

import (
	"gin-go-bl/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"varchar(20);not null"  json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Phone    string `gorm:"varchar(20);not null" json:"phone" validate:"required,number=true,min=11,max=11"label:"手机号"`
	Password string `gorm:"size:255;not null" json:"password" validate:"required,min=4,max=15"label:"密码"`
	Avatar   string `gorm:"size:255;not null" json:"avatar" `
}

func CheckUser(name string) (code int) {
	var user User
	DB.Select("id").Where("user_name = ?", name).First(&user)
	if user.ID != 0 {
		code = utils.ERROR_USERNAME_USED
		return
	}
	code = utils.SUCCESS
	return
}

//func (users User)GetUsers(pageSize int, pageNum int) ([]User, int64) {
//	var user []User
//	var total int64
//	DB.Model(&User{}).Count(&total)
//	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, 0
//	}
//	return user, total
//}
func GetUser(pageSize int, pageNum int) ([]User, int64) {
	var user []User
	var total int64
	DB.Model(&User{}).Count(&total)
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return user, total
}

func CreatUser(data *User) int {
	data.Password = utils.ScryptPassword(data.Password)
	err := DB.Create(&data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func DeleteUser(id int) int {
	var user User
	err := DB.Where("id = ?", id).Delete(&user).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func UpdateUser(id int, data *User) int {
	var user User
	var up = make(map[string]interface{})
	up["user_name"] = data.UserName
	up["phone"] = data.Phone
	up["avatar"] = data.Avatar
	err := DB.Model(&user).Where("id = ?", id).Updates(up).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func GetUserId(name string) User {
	var user User
	DB.Where("user_name =?", name).First(&user)
	if err != nil && err != gorm.ErrRecordNotFound {
		return user
	}
	return user
}

func CheckLogin(name string, password string) (code int) {
	var user User
	DB.Where("user_name = ?", name).First(&user)
	if user.ID == 0 {
		code = utils.ERROR_USERNAME_NOT
		return
	}
	if utils.ScryptPassword(password) != user.Password {
		code = utils.ERROR_PASSWORD_WRONG
		return
	}
	code = utils.SUCCESS
	return
}

func (User) TableName() string {
	return "user"
}
