package service

import(

	_itemShopRepository "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/repository"
) 



type itemShopServiceImpl struct {

	itemShopRepository _itemShopRepository.ItemShopRepository

}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
)ItemShopService{
	return &itemShopServiceImpl{itemShopRepository}
}