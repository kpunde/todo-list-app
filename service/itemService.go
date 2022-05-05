package service

import (
	"sampleAppDemo/entity"
	"sampleAppDemo/repository"
)

type ItemService interface {
	Save(item entity.Item) entity.Item
	Update(item entity.Item) entity.Item
	Delete(item entity.Item)
	FindAll() []entity.Item
	FindById(id int64) entity.Item
	FindByUserId(userId int64) []entity.Item
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func (i *itemService) FindById(id int64) entity.Item {
	return i.itemRepository.FindById(id)
}

func (i *itemService) FindByUserId(userId int64) []entity.Item {
	return i.itemRepository.FindByUserId(userId)
}

func (i *itemService) Update(item entity.Item) entity.Item {
	return i.itemRepository.Update(item)
}

func (i *itemService) Delete(item entity.Item) {
	i.itemRepository.Delete(item)
}

func (i *itemService) Save(item entity.Item) entity.Item {
	return i.itemRepository.Save(item)
}

func (i *itemService) FindAll() []entity.Item {
	return i.itemRepository.FindAll()
}

func NewItemService() ItemService {
	return &itemService{itemRepository: repository.NewItemRepository()}
}
