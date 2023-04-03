package model

import (
	"gin-go-bl/utils"
	"gorm.io/gorm"
)

type Category struct {
	Cid    uint   `gorm:"int;not null;primary_key;" json:"cid"`
	CaName string `gorm:"varchar(20);not null" json:"caname"`
}
type Categorys struct {
	Category           Category
	Categoryarticlenum int64 `json:"cateartnum"`
}

func CheckCategory(name string) (code int) {
	var category Category
	DB.Select("cid").Where("ca_name = ?", name).First(&category)
	if category.Cid != 0 {
		code = utils.ERROR_CATEGORY
		return
	}
	code = utils.SUCCESS
	return
}

func GetCategory(pageSize int, pageNum int) ([]Categorys, int64) {
	var categorys []Category
	var Categorylist []Categorys
	var Categorye Categorys
	var total int64
	DB.Model(&Category{}).Count(&total)
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	for _, v := range categorys {
		Categorye.Category = v
		DB.Model(&Article{}).Where("category_id = ?", v.Cid).Count(&Categorye.Categoryarticlenum)
		Categorylist = append(Categorylist, Categorye)
	}
	return Categorylist, total
}

func CreatCategory(data *Category) int {
	err := DB.Create(&data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func DeleteCategory(id int) int {
	var category Category
	err := DB.Where("cid = ?", id).Delete(&category).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func UpdateCategory(id int, data *Category) int {
	var category Category
	var up = make(map[string]interface{})
	up["ca_name"] = data.CaName
	err := DB.Model(&category).Where("cid = ?", id).Updates(up).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func (Category) TableName() string {
	return "category"
}
