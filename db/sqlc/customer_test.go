package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/M1racle-Heen/simple_order_pizza/util"
	"github.com/stretchr/testify/require"
)

func createRandomCustomer(t *testing.T) Customer {
	arg := CreateCustomerParams{
		FullName: util.RandomName(),
		Email:    util.RandomEmail(),
		Phone:    util.RandomNumber(),
		Address:  util.RandomAddress(),
	}

	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)

	require.Equal(t, arg.FullName, customer.FullName)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.Phone, customer.Phone)
	require.Equal(t, arg.Address, customer.Address)

	require.NotZero(t, customer.ID)
	return customer
}

func TestCreateCustomer(t *testing.T) {
	createRandomCustomer(t)
}

func TestDeleteCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	err := testQueries.DeleteCustomer(context.Background(), customer1.ID)
	require.NoError(t, err)

	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, customer2)
}

func TestGetCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomer(context.Background(), customer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Phone, customer2.Phone)
	require.Equal(t, customer1.Address, customer2.Address)
}

func TestGetCustomerForUpdate(t *testing.T) {
	customer1 := createRandomCustomer(t)
	customer2, err := testQueries.GetCustomerForUpdate(context.Background(), customer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Phone, customer2.Phone)
	require.Equal(t, customer1.Address, customer2.Address)
}

func TestUpdateCustomer(t *testing.T) {
	customer1 := createRandomCustomer(t)

	arg := UpdateCustomerAddressParams{
		ID:      customer1.ID,
		Address: util.RandomAddress(),
	}

	customer2, err := testQueries.UpdateCustomerAddress(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer2)

	require.Equal(t, customer1.ID, customer2.ID)
	require.Equal(t, customer1.FullName, customer2.FullName)
	require.Equal(t, customer1.Email, customer2.Email)
	require.Equal(t, customer1.Phone, customer2.Phone)
	require.Equal(t, arg.Address, customer2.Address)
}

func TestListCustomers(t *testing.T) {
	var lastCustomer Customer

	for i := 0; i < 10; i++ {
		lastCustomer = createRandomCustomer(t)
	}

	arg := ListCustomersParams{
		Limit:  5,
		Offset: 0,
	}

	customers, err := testQueries.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)

	for _, customer := range customers {
		require.NotEmpty(t, customer)
		require.NotEqual(t, lastCustomer.Email, nil)
	}
}
