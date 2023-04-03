package model

import (
	"fmt"
	"gin-go-bl/utils"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Category   Category  `gorm:"foreignKey:CategoryId"`
	Aid        uint      `gorm:"int;not null;primary_key" json:"aid"`
	Title      string    `gorm:"varchar(20);not null" json:"title"`
	CategoryId uint      `gorm:"int;not null" json:"categoryid"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	Desc       string    `gorm:"varchar(255);not null" json:"desc"`
	Content    string    `gorm:"longtext;not null" json:"content"`
	Img        string    `json:"img"`
	CreateTime time.Time `json:"crtime"`
}

func GetAllArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var article []Article
	var total int64

	DB.Model(&Article{}).Count(&total)
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR, 0
	}
	return article, utils.SUCCESS, total
}
func GetAllArticles() ([]Article, int, int64) {
	var article []Article
	var total int64

	DB.Model(&Article{}).Count(&total)
	err := DB.Preload("Category").Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR, 0
	}
	return article, utils.SUCCESS, total
}
func GetArticle(id int) (Article, int) {
	var article Article
	err := DB.Preload("Category").Where("aid = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, utils.ERROR_ART_NULL
	}
	return article, utils.SUCCESS
}

func GetUserArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var article []Article
	var total int64
	DB.Model(&Article{}).Where("user_id = ?", id).Count(&total)
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("user_id = ?", id).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR_ART_NULL, 0
	}
	return article, utils.SUCCESS, total
}

func GetCatArticle(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var article []Article
	var total int64
	DB.Model(&Article{}).Where("category_id = ?", id).Count(&total)
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("category_id = ?", id).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, utils.ERROR_CAART_NULL, 0
	}
	fmt.Println(article)
	return article, utils.SUCCESS, total
}

func CreatArticle(data *Article) int {
	err := DB.Create(&data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func DeleteArticle(id int) int {
	var article Article
	err := DB.Where("aid = ?", id).Delete(&article).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func EditArticle(id int, data *Article) int {
	var article Article
	var art = make(map[string]interface{})
	art["title"] = data.Title
	art["content"] = data.Content
	art["img"] = data.Img
	art["desc"] = data.Desc
	art["category_id"] = data.CategoryId
	err := DB.Model(&article).Where("aid = ?", id).Updates(art).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func (Article) TableName() string {
	return "article"
}
