package repository

import "database/sql"

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) FetchProductByID(id int64) (Product, error) {
	//TODO: You must implement this function fot fetch product by id
	sqlStatement := `SELECT id, product_name, category, price, quantity FROM products WHERE id = ?`

	row := p.db.QueryRow(sqlStatement, id)

	var product Product

	err := row.Scan(
		&product.ID,
		&product.ProductName,
		&product.Category,
		&product.Price,
		&product.Quantity,
	)

	if err != nil {
		return Product{}, err
	}
	return product, nil // TODO: replace this
}

func (p *ProductRepository) FetchProductByName(productName string) (Product, error) {
	// TODO: You must implement this function for fetch product by name
	sqlStatement := `SELECT id, product_name, category, price, quantity FROM products WHERE product_name = ?`

	row := p.db.QueryRow(sqlStatement, productName)

	var product Product

	err := row.Scan(
		&product.ID,
		&product.ProductName,
		&product.Category,
		&product.Price,
		&product.Quantity,
	)

	if err != nil {
		return Product{}, err
	}
	return product, nil // TODO: replace this
}

func (p *ProductRepository) FetchProducts() ([]Product, error) {
	// TODO: You must implement this function for fetch all products
	var sqlStatement string
	var products []Product


	sqlStatement = `SELECT id, product_name, category, price, quantity FROM products`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Category,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil // TODO: replace this
}
