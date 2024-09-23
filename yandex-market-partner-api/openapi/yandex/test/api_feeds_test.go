/*
Партнерский API Маркета

Testing FeedsAPIService

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

func Test_openapi_FeedsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FeedsAPIService GetFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FeedsAPI.GetFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeedsAPIService GetFeedIndexLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FeedsAPI.GetFeedIndexLogs(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeedsAPIService GetFeeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FeedsAPI.GetFeeds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeedsAPIService RefreshFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FeedsAPI.RefreshFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FeedsAPIService SetFeedParams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FeedsAPI.SetFeedParams(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
