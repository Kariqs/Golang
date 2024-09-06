package product

import "errors"

type Product struct {
	title string
	id    string
	price float64
}

func New(title string, id string, price float64) (*Product, error) {
	if title == "" || id == "" || price < 10 {
		return nil, errors.New("fill in all the fields")
	}
	return &Product{
		title: title,
		id:    id,
		price: price,
	}, nil
}
