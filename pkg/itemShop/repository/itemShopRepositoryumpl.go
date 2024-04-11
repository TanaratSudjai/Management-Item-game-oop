package repository

type ItemShopRepositoryImpl struct{}

func NewItemRepositoryImpl() ItemShopRepository{
	return & ItemShopRepositoryImpl{}
}