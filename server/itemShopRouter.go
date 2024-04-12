package server

import(
	_itemRepository "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/repository"
	_itemShopService "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/service"
	_itemShopController "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/controller"
)

func (s *echoServer) initItemShopRouter(){

	router := s.app.Group("/v1/item-shop")

	itemShopRepository := _itemRepository.NewItemRepositoryImpl(s.db, s.app.Logger)
	itemShopService := _itemShopService.NewItemShopServiceImpl(itemShopRepository)
	itemShopController := _itemShopController.NewItemShopControllerImpl(itemShopService)

	router.GET("", itemShopController.Listing)

}