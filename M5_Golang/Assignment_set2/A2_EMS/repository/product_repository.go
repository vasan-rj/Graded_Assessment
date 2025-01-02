package repository

import (
	"database/sql"
	"ecommerce/model"
	"errors"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (repo *ItemRepository) AddItem(item *model.Item) error {
	_, err := repo.db.Exec(`INSERT INTO products (name, description, price, stock, category_id) 
		VALUES (?, ?, ?, ?, ?)`, item.Title, item.Details, item.Cost, item.Quantity, item.CategoryRef)
	return err
}

func (repo *ItemRepository) GetItemByID(id int) (*model.Item, error) {
	row := repo.db.QueryRow(`SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`, id)
	item := &model.Item{}
	err := row.Scan(&item.ID, &item.Title, &item.Details, &item.Cost, &item.Quantity, &item.CategoryRef)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return item, nil
}

func (repo *ItemRepository) UpdateItem(item *model.Item) error {
	_, err := repo.db.Exec(`UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? 
		WHERE id = ?`, item.Title, item.Details, item.Cost, item.Quantity, item.CategoryRef, item.ID)
	return err
}

func (repo *ItemRepository) DeleteItem(id int) error {
	_, err := repo.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	return err
}

func (repo *ItemRepository) GetAllItems(page, limit int) ([]model.Item, error) {
	rows, err := repo.db.Query(`SELECT id, name, description, price, stock, category_id FROM products LIMIT ? OFFSET ?`,
		limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Details, &item.Cost, &item.Quantity, &item.CategoryRef); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
