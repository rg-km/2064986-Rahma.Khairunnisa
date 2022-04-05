package repository

import (
	"strconv"
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
	data, err := u.db.Load("product")
	if err != nil {
		var record = [][]string{{
			"category", "product_name", "price",
		}}
		if err := u.db.Save("product",record); err != nil {
			return nil, err
		}
	}

	var products []Product
	
	for _, dat := range data {
		price, err := strconv.Atoi(dat[2])
		if err != nil {
			return nil, err
		}
		product := Product{
			Category: dat[0],
			ProductName: dat[1],
			Price: price,
		}

		products = append(products, product)
	}

	return products, nil
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	return u.LoadOrCreate() // TODO: replace this
}
