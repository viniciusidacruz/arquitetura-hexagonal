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

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
    product.Price = 0
    product.Name = "test product"
    product.Status = application.ENABLED

    err := product.Disable()
    require.Nil(t, err)

    product.Price = 100
    err = product.Disable()
    require.Error(t, err)
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
    product.ID = "test_id"

    require.Equal(t, product.GetID(), "test_id")
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
    product.Name = "test product"

    require.Equal(t, product.GetName(), "test product")
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
    product.Price = 100.2

    require.Equal(t, product.GetPrice(), 100.2)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
    product.Status = application.ENABLED

    require.Equal(t, product.GetStatus(), application.ENABLED)
}