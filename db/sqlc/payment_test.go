package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPayment(t *testing.T, pizza Pizza, customer Customer) Payment {
	pizzaPrice, err := testQueries.GetPizza(context.Background(), pizza.ID)
	require.NoError(t, err)
	require.NotEmpty(t, pizzaPrice)

	status, err := testQueries.GetOrder(context.Background(), pizzaPrice.OrderID)
	require.NoError(t, err)
	require.NotEmpty(t, status)
	isHold := ""
	if status.Status == "Hold" {
		isHold = "NotPaid"
	} else {
		isHold = "Paid"
	}

	arg := CreatePaymentParams{
		PizzaID:       pizza.ID,
		CustomerID:    customer.ID,
		PaymentStatus: isHold,
		Bill:          pizzaPrice.Price,
	}

	payment, err := testQueries.CreatePayment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment)

	require.Equal(t, arg.PizzaID, payment.PizzaID)
	require.Equal(t, arg.CustomerID, payment.CustomerID)
	require.Equal(t, arg.PaymentStatus, payment.PaymentStatus)
	require.Equal(t, arg.Bill, payment.Bill)

	require.NotZero(t, payment.ID)

	return payment
}

func TestCreatePayment(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	pizza := createRandomPizza(t, order)

	createRandomPayment(t, pizza, customer)
}

func TestGetPayment(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	pizza := createRandomPizza(t, order)

	payment1 := createRandomPayment(t, pizza, customer)

	payment2, err := testQueries.GetPayment(context.Background(), payment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, payment2)

	require.Equal(t, payment1.ID, payment2.ID)
	require.Equal(t, payment1.PizzaID, payment2.PizzaID)
	require.Equal(t, payment1.CustomerID, payment2.CustomerID)
	require.Equal(t, payment1.PaymentStatus, payment2.PaymentStatus)
	require.Equal(t, payment1.Bill, payment2.Bill)
}

func TestListPayments(t *testing.T) {
	customer1 := createRandomCustomer(t)
	order1 := createRandomOrder(t, customer1)
	pizza1 := createRandomPizza(t, order1)

	for i := 0; i < 10; i++ {
		createRandomPayment(t, pizza1, customer1)
	}

	arg := ListPaymentsParams{
		Limit:  5,
		Offset: 5,
	}

	payments, err := testQueries.ListPayments(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, payments, 5)

	for _, payment := range payments {
		require.NotEmpty(t, payment)

	}
}

func TestUpdatePayment(t *testing.T) {
	customer := createRandomCustomer(t)
	order := createRandomOrder(t, customer)
	pizza := createRandomPizza(t, order)

	payment1 := createRandomPayment(t, pizza, customer)

	arg := UpdatePaymentStatusParams{
		ID:            payment1.ID,
		PaymentStatus: payment1.PaymentStatus,
	}

	payment2, err := testQueries.UpdatePaymentStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, payment2)

	require.Equal(t, payment1.ID, payment2.ID)
	require.Equal(t, payment1.PizzaID, payment2.PizzaID)
	require.Equal(t, payment1.CustomerID, payment2.CustomerID)
	require.Equal(t, arg.PaymentStatus, payment2.PaymentStatus)
	require.Equal(t, payment1.Bill, payment2.Bill)
}
