package repository

import (
	"sampleAppDemo/entity"
	"sampleAppDemo/utility"
)

type ItemRepository interface {
	Save(item entity.Item) entity.Item
	Update(item entity.Item) entity.Item
	Delete(item entity.Item)
	FindAll() []entity.Item
	FindById(id int64) entity.Item
	FindByUserId(userId int64) []entity.Item
}

type database struct {
}

func (d database) Save(item entity.Item) entity.Item {
	utility.DB.Create(&item)
	return item
}

func (d database) Update(item entity.Item) entity.Item {
	utility.DB.Save(&item)
	return item
}

func (d database) Delete(item entity.Item) {
	utility.DB.Delete(&item)
}

func (d database) FindAll() []entity.Item {
	var items []entity.Item
	utility.DB.Preload("Author").Find(&items)
	return items
}

func (d database) FindById(id int64) entity.Item {
	var item entity.Item
	utility.DB.Where("id =?", id).First(&item)
	return item
}

func (d database) FindByUserId(userId int64) []entity.Item {
	var items []entity.Item
	utility.DB.Preload("Author").Where("user_id =?", userId).Find(&items)
	return items
}

func NewItemRepository() ItemRepository {
	return &database{}
}
