package repository

import "github.com/TanaratSudjai/project-golang-api-shop/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}