package model

type (
	Item struct {
		ID          uint64 `json:id`
		Name        string `json:name`
		Description string `json:discription`
		Picture     string `json:picture`
		Price       uint   `json:price`
	}
)
