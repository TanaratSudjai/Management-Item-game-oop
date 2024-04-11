package controller

import(

	_itemShopService "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/service"
) 

type itemShopControllerImpl struct{

	itemShopService  _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopService.ItemShopService,
)ItemShopController{
	return &itemShopControllerImpl{itemShopService}
}