/*
Партнерский API Маркета

Testing OutletsAPIService

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

func Test_openapi_OutletsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OutletsAPIService CreateOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.OutletsAPI.CreateOutlet(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutletsAPIService DeleteOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.OutletsAPI.DeleteOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutletsAPIService GetOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.OutletsAPI.GetOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutletsAPIService GetOutlets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.OutletsAPI.GetOutlets(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutletsAPIService UpdateOutlet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var outletId int64

		resp, httpRes, err := apiClient.OutletsAPI.UpdateOutlet(context.Background(), campaignId, outletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
