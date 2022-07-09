package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/M1racle-Heen/simple_order_pizza/db/mock"
	db "github.com/M1racle-Heen/simple_order_pizza/db/sqlc"
	"github.com/M1racle-Heen/simple_order_pizza/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCustomerAPI(t *testing.T) {
	customer := randomCustomer()

	testCases := []struct {
		name          string
		customerID    int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			customerID: customer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomer(gomock.Any(), gomock.Eq(customer.ID)).
					Times(1).
					Return(customer, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchCustomer(t, recorder.Body, customer)

			},
		},
		{
			name:       "NotFound",
			customerID: customer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomer(gomock.Any(), gomock.Eq(customer.ID)).
					Times(1).
					Return(db.Customer{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "BadRequest",
			customerID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:       "InternalServerError",
			customerID: customer.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCustomer(gomock.Any(), gomock.Eq(customer.ID)).
					Times(1).
					Return(db.Customer{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/customers/%d", tc.customerID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomCustomer() db.Customer {
	return db.Customer{
		ID:       util.RandomInt(1, 1000),
		FullName: util.RandomName(),
		Email:    util.RandomEmail(),
		Phone:    util.RandomNumber(),
		Address:  util.RandomAddress(),
	}
}

func requiredBodyMatchCustomer(t *testing.T, body *bytes.Buffer, customer db.Customer) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotCustomer db.Customer
	err = json.Unmarshal(data, &gotCustomer)
	require.NoError(t, err)
	require.Equal(t, customer, gotCustomer)
}
