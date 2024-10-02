package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/viniciusidacruz/arquitetura-hexagonal/application"
)

type ProductDatabase struct {
	database *sql.DB
}

func NewProductDatabase(db *sql.DB) *ProductDatabase {
	return &ProductDatabase{database: db}
}

func (p *ProductDatabase) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.database.Prepare("select id, name, price, status from products where id=?")

	if err!= nil {
        return nil, err
    }

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}