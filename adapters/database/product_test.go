package database_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/arquitetura-hexagonal/adapters/database"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory")

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
        "price" float,
        "status" string
	)`

	stmt, err := db.Prepare(table);

	if err!= nil {
        log.Fatal(err.Error())
    }

	stmt.Exec()
}

func createProduct(db *sql.DB){
	insert := `insert into products values("123", "Arroz", 0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err!= nil {
        log.Fatal(err.Error())
    }

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := database.NewProductDatabase(Db)

	product, err := productDb.Get("123")
	require.Nil(t, err)
	require.Equal(t, "Arroz", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}