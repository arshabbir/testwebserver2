package services

import (
	"github.com/arshabbir/docker1/src/domain/store"
	"github.com/arshabbir/docker1/src/utils"
)

type itemService struct {
	dao store.StoreDao
}

type ItemService interface {
	Create(*store.Item) *utils.ApiErr
	Search(string) (*store.Item, *utils.ApiErr)
	PingItemService() string
}

func NewItemService() ItemService {

	return &itemService{dao: store.NewStoreDao()}
}

func (is *itemService) Create(item *store.Item) *utils.ApiErr {
	return is.dao.Create(item)
}

func (is *itemService) Search(id string) (*store.Item, *utils.ApiErr) {
	return is.dao.Search(id)
}

func (is *itemService) PingItemService() string {

	return is.dao.Ping()
}
