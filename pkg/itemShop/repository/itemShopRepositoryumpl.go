package repository

import (
	"github.com/TanaratSudjai/project-golang-api-shop/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_itemShopException "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/exception"
)

type ItemShopRepositoryImpl struct{

	db *gorm.DB
	logger echo.Logger
}

func NewItemRepositoryImpl(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &ItemShopRepositoryImpl{db, logger}
}

func (r *ItemShopRepositoryImpl) Listing() ([]*entities.Item,error){
	itemList := make([]*entities.Item, 0)

	if err := r.db.Find(&itemList).Error; err != nil {
		r.logger.Errorf("Filed to list items : %s",err.Error())
		return nil, &_itemShopException.ItemListing{}
	}

	return itemList, nil
}