package application_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/viniciusidacruz/arquitetura-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Price = 100
	product.Name = "test product"
	product.Status = application.DISABLED

    err := product.Enable()
    require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Error(t, err)
}