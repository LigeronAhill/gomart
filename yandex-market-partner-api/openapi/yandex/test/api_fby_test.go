/*
Партнерский API Маркета

Testing FbyAPIService

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

func Test_openapi_FbyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FbyAPIService AddHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.AddHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService AddOffersToArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.AddOffersToArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService CalculateTariffs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.CalculateTariffs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService ConfirmBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.ConfirmBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService ConfirmCampaignPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.ConfirmCampaignPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService CreateChat", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.CreateChat(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeleteCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.DeleteCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeleteGoodsFeedbackComment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.DeleteGoodsFeedbackComment(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeleteHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.DeleteHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeleteOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.DeleteOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeleteOffersFromArchive", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.DeleteOffersFromArchive(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService DeletePromoOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.DeletePromoOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateBoostConsolidatedReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateBoostConsolidatedReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateCompetitorsPositionReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateCompetitorsPositionReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateGoodsFeedbackReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateGoodsFeedbackReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateGoodsMovementReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateGoodsMovementReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateGoodsRealizationReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateGoodsRealizationReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateGoodsTurnoverReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateGoodsTurnoverReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GeneratePricesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GeneratePricesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateShelfsStatisticsReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateShelfsStatisticsReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateShowsSalesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateShowsSalesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateStocksOnWarehousesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateStocksOnWarehousesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateUnitedMarketplaceServicesReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateUnitedMarketplaceServicesReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateUnitedNettingReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateUnitedNettingReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GenerateUnitedOrdersReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GenerateUnitedOrdersReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetAllOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetAllOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetBidsInfoForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetBidsInfoForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetBidsRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetBidsRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetBusinessQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetBusinessQuarantineOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetBusinessSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetBusinessSettings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaign(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignLogins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignLogins(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignQuarantineOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignQuarantineOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignRegion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignRegion(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignSettings(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaigns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GetCampaigns(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCampaignsByLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var login string

		resp, httpRes, err := apiClient.FbyAPI.GetCampaignsByLogin(context.Background(), login).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCategoriesMaxSaleQuantum", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GetCategoriesMaxSaleQuantum(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCategoriesTree", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GetCategoriesTree(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetCategoryContentParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var categoryId int64

		resp, httpRes, err := apiClient.FbyAPI.GetCategoryContentParameters(context.Background(), categoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetChatHistory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetChatHistory(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetChats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetChats(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FbyAPI.GetFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetFeedIndexLogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FbyAPI.GetFeedIndexLogs(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetFeedbackAndCommentUpdates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetFeedbackAndCommentUpdates(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetFeeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetFeeds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetFulfillmentWarehouses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.GetFulfillmentWarehouses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetGoodsFeedbackComments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetGoodsFeedbackComments(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetGoodsFeedbacks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetGoodsFeedbacks(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetGoodsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetGoodsStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetHiddenOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetHiddenOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOfferCardsContentStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOfferCardsContentStatus(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOfferRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOfferRecommendations(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOrder(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOrderBusinessBuyerInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOrderBusinessBuyerInfo(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOrderBusinessDocumentsInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOrderBusinessDocumentsInfo(context.Background(), campaignId, orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOrders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOrders(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetOrdersStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetOrdersStats(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetPricesByOfferIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetPricesByOfferIds(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetPromoOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetPromoOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetPromos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetPromos(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetQualityRatings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetQualityRatings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetReportInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reportId string

		resp, httpRes, err := apiClient.FbyAPI.GetReportInfo(context.Background(), reportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetReturn", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64

		resp, httpRes, err := apiClient.FbyAPI.GetReturn(context.Background(), campaignId, orderId, returnId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetReturnPhoto", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var orderId int64
		var returnId int64
		var itemId int64
		var imageHash string

		resp, httpRes, err := apiClient.FbyAPI.GetReturnPhoto(context.Background(), campaignId, orderId, returnId, itemId, imageHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetReturns", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetReturns(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetStocks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetStocks(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetSuggestedOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetSuggestedOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetSuggestedOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.GetSuggestedOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService GetSuggestedPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.GetSuggestedPrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService PutBidsForBusiness", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.PutBidsForBusiness(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService PutBidsForCampaign", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.PutBidsForCampaign(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService RefreshFeed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FbyAPI.RefreshFeed(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SearchRegionChildren", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.FbyAPI.SearchRegionChildren(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SearchRegionsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var regionId int64

		resp, httpRes, err := apiClient.FbyAPI.SearchRegionsById(context.Background(), regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SearchRegionsByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FbyAPI.SearchRegionsByName(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SendFileToChat", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.SendFileToChat(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SendMessageToChat", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.SendMessageToChat(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SetFeedParams", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64
		var feedId int64

		resp, httpRes, err := apiClient.FbyAPI.SetFeedParams(context.Background(), campaignId, feedId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService SkipGoodsFeedbacksReaction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.SkipGoodsFeedbacksReaction(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateBusinessPrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateBusinessPrices(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateCampaignOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateCampaignOffers(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateGoodsFeedbackComment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateGoodsFeedbackComment(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateOfferContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateOfferContent(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateOfferMappingEntries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateOfferMappingEntries(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdateOfferMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdateOfferMappings(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdatePrices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var campaignId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdatePrices(context.Background(), campaignId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FbyAPIService UpdatePromoOffers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var businessId int64

		resp, httpRes, err := apiClient.FbyAPI.UpdatePromoOffers(context.Background(), businessId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
