package db

import (
	"context"
	"testing"

	"github.com/M1racle-Heen/simple_order_pizza/util"
	"github.com/stretchr/testify/require"
)

func createRandomPizza(t *testing.T, order Order) Pizza {
	q := util.RandomInt(1, 5)
	p := util.RandomInt(1000, 2000) * q
	arg := CreatePizzaParams{
		OrderID:    order.ID,
		Price:      p,
		PizzaType:  util.RandomPizzaType(),
		PizzaQuant: q,
	}

	pizza, err := testQueries.CreatePizza(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, pizza)

	require.Equal(t, arg.OrderID, pizza.OrderID)
	require.Equal(t, arg.Price, pizza.Price)
	require.Equal(t, arg.PizzaType, pizza.PizzaType)
	require.Equal(t, arg.PizzaQuant, pizza.PizzaQuant)

	require.NotZero(t, pizza.ID)

	return pizza
}

func TestCreatePizza(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	createRandomPizza(t, order)
}

func TestGetPizza(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)

	pizza1 := createRandomPizza(t, order)

	pizza2, err := testQueries.GetPizza(context.Background(), pizza1.ID)
	require.NoError(t, err)
	require.NotZero(t, pizza2)

	require.Equal(t, pizza1.ID, pizza2.ID)
	require.Equal(t, pizza1.OrderID, pizza2.OrderID)
	require.Equal(t, pizza1.Price, pizza2.Price)
	require.Equal(t, pizza1.PizzaType, pizza2.PizzaType)
	require.Equal(t, pizza1.PizzaQuant, pizza2.PizzaQuant)
}

func TestListPizzas(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	for i := 0; i < 10; i++ {
		createRandomPizza(t, order)
	}

	arg := ListPizzasParams{
		OrderID: order.ID,
		Limit:   5,
		Offset:  5,
	}

	pizzas, err := testQueries.ListPizzas(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, pizzas, 5)

	for _, pizza := range pizzas {
		require.NotEmpty(t, pizza)
		require.Equal(t, arg.OrderID, pizza.OrderID)
	}
}
