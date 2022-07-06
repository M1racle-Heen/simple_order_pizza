package db

import (
	"context"
	"testing"
	"time"

	"github.com/M1racle-Heen/simple_order_pizza/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T, customer Customer) Order {
	arg := CreateOrderParams{
		CustomerID: customer.ID,
		Status:     util.RandomStatus(),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.CustomerID, order.CustomerID)
	require.Equal(t, arg.Status, order.Status)

	require.NotZero(t, order.ID)
	require.NotZero(t, order.OrderTime)

	return order
}

func TestCreateOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	createRandomOrder(t, customer)
}

func TestGetOrder(t *testing.T) {
	customer := createRandomCustomer(t)
	order1 := createRandomOrder(t, customer)
	order2, err := testQueries.GetOrder(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.CustomerID, order2.CustomerID)
	require.Equal(t, order1.Status, order2.Status)
	require.WithinDuration(t, order1.OrderTime, order2.OrderTime, time.Second)
}

func TestListOrders(t *testing.T) {
	customer := createRandomCustomer(t)
	for i := 0; i < 10; i++ {
		createRandomOrder(t, customer)
	}

	arg := ListOrdersParams{
		CustomerID: customer.ID,
		Limit:      5,
		Offset:     5,
	}

	orders, err := testQueries.ListOrders(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, orders, 5)

	for _, order := range orders {
		require.NotEmpty(t, order)
		require.Equal(t, arg.CustomerID, order.CustomerID)
	}
}
