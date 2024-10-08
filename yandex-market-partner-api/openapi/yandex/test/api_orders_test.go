/*
Партнерский API Маркета

Testing OrdersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_OrdersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrdersAPIService AcceptOrderCancellation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.AcceptOrderCancellation(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.GetOrder(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService GetOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.OrdersAPI.GetOrders(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService ProvideOrderDigitalCodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.ProvideOrderDigitalCodes(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService ProvideOrderItemIdentifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.ProvideOrderItemIdentifiers(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService SetOrderBoxLayout", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.SetOrderBoxLayout(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService SetOrderShipmentBoxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var shipmentId int64

		resp, httpRes, err := apiClient.OrdersAPI.SetOrderShipmentBoxes(context.Background(), campaignId, orderId, shipmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService UpdateOrderItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		httpRes, err := apiClient.OrdersAPI.UpdateOrderItems(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService UpdateOrderStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.OrdersAPI.UpdateOrderStatus(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrdersAPIService UpdateOrderStatuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.OrdersAPI.UpdateOrderStatuses(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
