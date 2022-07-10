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
	"github.com/gin-gonic/gin"
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

func TestCreateCustomerAPI(t *testing.T) {
	customer := randomCustomer()

	testCases := []struct {
		name         string
		body         gin.H
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{
			name: "OK",
			body: gin.H{
				"full_name": customer.FullName,
				"email":     customer.Email,
				"phone":     customer.Phone,
				"address":   customer.Address,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateCustomerParams{
					FullName: customer.FullName,
					Email:    customer.Email,
					Phone:    customer.Phone,
					Address:  customer.Address,
				}

				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(customer, nil)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "InvalidFullName",
			body: gin.H{
				"full_name": "",
				"email":     customer.Email,
				"phone":     customer.Phone,
				"address":   customer.Address,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"full_name": customer.FullName,
				"email":     "",
				"phone":     customer.Phone,
				"address":   customer.Address,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidPhone",
			body: gin.H{
				"full_name": customer.FullName,
				"email":     customer.Email,
				"phone":     0,
				"address":   customer.Address,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidAddress",
			body: gin.H{
				"full_name": customer.FullName,
				"email":     customer.Email,
				"phone":     customer.Phone,
				"address":   "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InternalError",
			body: gin.H{
				"full_name": customer.FullName,
				"email":     customer.Email,
				"phone":     customer.Phone,
				"address":   customer.Address,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCustomer(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Customer{}, sql.ErrConnDone)
			},
			expectStatus: http.StatusInternalServerError,
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

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			utl := "/customers"
			request, err := http.NewRequest(http.MethodPost, utl, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)

		})
	}
}

func TestListCustomersAPI(t *testing.T) {
	n := 5
	customers := make([]db.Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = randomCustomer()
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListCustomersParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListCustomers(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(customers, nil)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchCustomers(t, recoder.Body, customers)
			},
		},
		{
			name: "InternalError",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListCustomers(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Customer{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   -1,
				pageSize: n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListCustomers(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			query: Query{
				pageID:   1,
				pageSize: 100000,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListCustomers(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

			url := "/customers"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
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

func requireBodyMatchCustomers(t *testing.T, body *bytes.Buffer, customers []db.Customer) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotCustomers []db.Customer
	err = json.Unmarshal(data, &gotCustomers)
	require.NoError(t, err)
	require.Equal(t, customers, gotCustomers)
}
