package service
import(
	
	_itemShopModel "github.com/TanaratSudjai/project-golang-api-shop/pkg/itemShop/model"
)

type ItemShopService interface{
	Listing() ([]*_itemShopModel.Item, error)
}