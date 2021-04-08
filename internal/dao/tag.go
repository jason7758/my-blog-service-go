package dao

import (
	"my-blog-service-go/internal/model"
	"my-blog-service-go/pkg/app"
)

func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error)  {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.Tag, error)  {
	tag := model.Tag{State: state}

}