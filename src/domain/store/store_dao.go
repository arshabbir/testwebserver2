package store

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arshabbir/docker1/src/utils"
)

type storeDao struct {
	items map[string]Item
}

type StoreDao interface {
	Create(*Item) *utils.ApiErr
	Search(string) (*Item, *utils.ApiErr)
	Ping() string
}

func NewStoreDao() StoreDao {
	return &storeDao{items: make(map[string]Item)}
}

func (s *storeDao) Create(item *Item) *utils.ApiErr {

	s.items[item.Id] = *item

	log.Println("Iteam Created ", item.Id)
	return nil
}

func (s *storeDao) Search(id string) (*Item, *utils.ApiErr) {

	item, ok := s.items[id]

	if !ok {
		return nil, &utils.ApiErr{Status: http.StatusNotFound, Msg: fmt.Sprintf("Item not found: %s", id)}
	}

	return &item, nil
}

func (s *storeDao) Ping() string {
	return "pong"
}
