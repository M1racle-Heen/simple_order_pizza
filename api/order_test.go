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

func TestUpdateOrderStatusAPI(t *testing.T) {
	customer := randomCustomer()
	order := randomOrder(customer)

	testCases := []struct {
		name         string
		body         gin.H
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{
			name: "OK",
			body: gin.H{
				"id":     order.ID,
				"status": order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateOrderStatusParams{
					ID:     order.ID,
					Status: order.Status,
				}

				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(order, nil)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "InvalidID",
			body: gin.H{
				"id":     0,
				"status": order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidStatus",
			body: gin.H{
				"id":     order.ID,
				"status": "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InternalError",
			body: gin.H{
				"id":     order.ID,
				"status": order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Order{}, sql.ErrConnDone)
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

			utl := "/orders"
			request, err := http.NewRequest(http.MethodPut, utl, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}
func TestGetOrderAPI(t *testing.T) {
	customer := randomCustomer()
	order := randomOrder(customer)
	testCases := []struct {
		name          string
		orderID       int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			orderID: order.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetOrder(gomock.Any(), gomock.Eq(order.ID)).
					Times(1).
					Return(order, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchOrder(t, recorder.Body, order)

			},
		},
		{
			name:    "NotFound",
			orderID: order.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetOrder(gomock.Any(), gomock.Eq(order.ID)).
					Times(1).
					Return(db.Order{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "BadRequest",
			orderID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetOrder(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "InternalServerError",
			orderID: order.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetOrder(gomock.Any(), gomock.Eq(order.ID)).
					Times(1).
					Return(db.Order{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/orders/%d", tc.orderID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateOrderAPI(t *testing.T) {
	customer := randomCustomer()
	order := randomOrder(customer)

	testCases := []struct {
		name         string
		body         gin.H
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{
			name: "OK",
			body: gin.H{
				"customer_id": order.CustomerID,
				"status":      order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateOrderParams{
					CustomerID: order.CustomerID,
					Status:     order.Status,
				}

				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(order, nil)
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "InvalidCustomerID",
			body: gin.H{
				"customer_id": 0,
				"status":      order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidStatus",
			body: gin.H{
				"customer_id": order.CustomerID,
				"status":      "",
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "InternalError",
			body: gin.H{
				"customer_id": order.CustomerID,
				"status":      order.Status,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateOrder(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Order{}, sql.ErrConnDone)
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

			utl := "/orders"
			request, err := http.NewRequest(http.MethodPost, utl, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
		})
	}
}

func TestListOrdersAPI(t *testing.T) {
	n := 5
	customers := randomCustomer()
	orders := make([]db.Order, n)
	for i := 0; i < n; i++ {
		orders[i] = randomOrder(customers)
	}
	fmt.Println(orders)

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
				arg := db.ListOrdersParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListOrders(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(orders, nil)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				fmt.Println(recoder.Code)
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchOrders(t, recoder.Body, orders)
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
					ListOrders(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Order{}, sql.ErrConnDone)
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
					ListOrders(gomock.Any(), gomock.Any()).
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
					ListOrders(gomock.Any(), gomock.Any()).
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

			url := "/orders"
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

func requiredBodyMatchOrder(t *testing.T, body *bytes.Buffer, order db.Order) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotOrder db.Order
	err = json.Unmarshal(data, &gotOrder)
	require.NoError(t, err)
	require.Equal(t, order, gotOrder)
}

func randomOrder(customer db.Customer) db.Order {
	return db.Order{
		ID:         util.RandomInt(1, 1000),
		CustomerID: customer.ID,
		Status:     util.RandomStatus(),
	}
}

func requireBodyMatchOrders(t *testing.T, body *bytes.Buffer, order []db.Order) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)
	fmt.Println("-------------------------------------------------", err)
	var gotOrders []db.Order
	err = json.Unmarshal(data, &gotOrders)
	fmt.Println("-------------------------------------------------", err)
	require.NoError(t, err)
	require.Equal(t, order, gotOrders)
}
