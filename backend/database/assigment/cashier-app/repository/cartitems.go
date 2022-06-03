package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type CartItemRepository struct {
	db *sql.DB
}

func NewCartItemRepository(db *sql.DB) *CartItemRepository {
	return &CartItemRepository{db: db}
}

func (c *CartItemRepository) FetchCartItems() ([]CartItem, error) {
	var sqlStatement string
	var cartItems []CartItem

	//TODO: add sql statement here
	//HINT: join table cart_items and products
	sqlStatement = `SELECT 
	c.id,
	p.category,
	c.product_id,
	p.product_name,
	p.price,
	c.quantity
	FROM cart_items C
	INNER JOIN products p ON c.product_id = p.id`

	rows, err := c.db.Query(sqlStatement)
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	for rows.Next() {
		var cartItem CartItem

		err := rows.Scan(
			&cartItem.ID,
			&cartItem.Category,
			&cartItem.ProductID,
			&cartItem.ProductName,
			&cartItem.Price,
			&cartItem.Quantity,
		)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil // TODO: replace this
}

func (c *CartItemRepository) FetchCartByProductID(productID int64) (CartItem, error) {
	var cartItem CartItem
	var sqlStatement string
	//TODO : you must fetch the cart by product id
	//HINT : you can use the where statement

	sqlStatement = `SELECT 
	c.id,
	p.category,
	c.product_id,
	p.product_name,
	p.price,
	c.quantity
	FROM cart_items C
	INNER JOIN products p ON c.product_id = p.id
	WHERE c.product_id = ?`

	row := c.db.QueryRow(sqlStatement, productID)
	
	err := row.Scan(
		&cartItem.ID,
		&cartItem.Category,
		&cartItem.ProductID,
		&cartItem.ProductName,
		&cartItem.Price,
		&cartItem.Quantity,
	)

	if err != nil {
		return CartItem{}, err
	}

	return cartItem, nil // TODO: replace this
}

func (c *CartItemRepository) InsertCartItem(cartItem CartItem) error {
	// TODO: you must insert the cart item
	_, err := c.FetchCartByProductID(cartItem.ProductID)
	
	if err != nil {
		sqlStmt := `INSERT INTO cart_items (product_id, quantity)
		VALUES (?, ?)`

		_, err := c.db.Exec(sqlStmt, cartItem.ProductID, cartItem.Quantity)
		if err != nil {
			return err
		}
		return nil
	}

	err = c.IncrementCartItemQuantity(cartItem)
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (c *CartItemRepository) IncrementCartItemQuantity(cartItem CartItem) error {
	//TODO : you must update the quantity of the cart item
	sqlStmt := `UPDATE cart_items SET quantity = quantity + ? WHERE product_id = ?`

	_, err := c.db.Exec(sqlStmt, cartItem.Quantity, cartItem.ProductID)

	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (c *CartItemRepository) ResetCartItems() error {
	//TODO : you must reset the cart items
	//HINT : you can use the delete statement
	sqlStmt := `DELETE FROM cart_items`

	_, err := c.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (c *CartItemRepository) TotalPrice() (int, error) {
	var sqlStatement string
	
	//TODO : you must calculate the total price of the cart items
	//HINT : you can use the sum statement
	sqlStatement = `SELECT SUM(c.quantity * p.price)
	FROM cart_items c
	INNER JOIN products p ON c.product_id = p.id`

	row := c.db.QueryRow(sqlStatement)

	var hasil int

	err := row.Scan(
		&hasil,
	)
	if err != nil {
	 	return 0, err
	}

	return hasil, nil // TODO: replace this
}
