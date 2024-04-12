package controller

import (
	"net/http"

	"github.com/TanaratSudjai/project-golang-api-shop/pkg/custom"
	_itemShopService "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/service"
	"github.com/labstack/echo/v4"

	_itemShopException "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/exception"
	
) 

type itemShopControllerImpl struct{

	itemShopService  _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(
	itemShopService _itemShopService.ItemShopService,
)ItemShopController{
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error{

	itemModelList, err := c.itemShopService.Listing()
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemModelList)

	//return custom.Error(pctx, http.StatusInternalServerError, (&_itemShopException.ItemListing{}).Error())
}