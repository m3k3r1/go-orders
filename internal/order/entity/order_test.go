package entity_test

import (
	"github.com/m3k3r1/go-orders/internal/order/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.Validate(), "order ID should not be empty")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := entity.Order{ID: "123"}
	assert.Error(t, order.Validate(), "order price should not be empty")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 1.23}
	assert.Error(t, order.Validate(), "order tax should not be empty")
}

func TestGivenValidParams_WhenCreateANewOrder_ThenShouldReceiveOrder(t *testing.T) {
	order, err := entity.NewOrder("123", 1.23, 0.12)

	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 1.23, order.Price)
	assert.Equal(t, 0.12, order.Tax)
	assert.Equal(t, order.Price+order.Tax, order.FinalPrice)
}

func TestGivenValidParams_WhenCalculateFinalPrice_ThenShouldSetValue(t *testing.T) {
	order, err := entity.NewOrder("123", 1.23, 0.12)
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 1.23, order.Price)
	assert.Equal(t, 0.12, order.Tax)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, order.Price+order.Tax, order.FinalPrice)
}
