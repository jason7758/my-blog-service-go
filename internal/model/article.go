package model

import (
	"github.com/jinzhu/gorm"
	"my-blog-service-go/pkg/app"
)

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Context string `json:"context"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article) TableName() string  {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

//更新某条数据	
func (a Article) Update(db *gorm.DB, values interface{}) error  {
	if err := db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

//获取某个编号的数据
func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	
	db = db.Where("id = ? AND state = ? and is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error 
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

type ArticleRow struct {
	ArticleID uint32
	TagID uint32
	TagName string
	ArticleTitle string

}

