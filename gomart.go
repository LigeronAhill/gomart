package gomart

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/LigeronAhill/gomart/internal/models"
)

// <http_method> https://api.partner.market.yandex.ru/<resource>.<format>?<parameters>
// <http_method> ― DELETE, GET, POST или PUT.
// <resource> ― URL ресурса, над которым выполняется действие. Названия ресурсов приведены в описании соответствующих методов.
// <format> ― это необязательная часть запроса, которая влияет на способ представления ответа. Формат ответа может быть указан в HTTP-заголовке Accept. Данные передаются в формате JSON. В описании каждого метода приведены примеры запросов и ответов.
// <parameters> ― обязательные и необязательные параметры запроса, которые не входят в состав URL ресурса.

type Client struct {
	client    http.Client
	token     string
	baseUrl   url.URL
	campaigns []*models.CampaignDTO
	ctx       context.Context
}

func Init(token string) (*Client, error) {
	client := http.Client{}
	ctx := context.Background()
	token = "Bearer " + token
	baseUrl := url.URL{
		Scheme: "https",
		Host:   "api.partner.market.yandex.ru",
	}
	temp := Client{
		client:  client,
		token:   token,
		baseUrl: baseUrl,
		ctx:     ctx,
	}
	err := temp.getCampaigns(1, 100)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &temp, nil
}

func (c *Client) getCampaigns(page, pageSize int) error {
	query := url.Values{}
	query.Add("page", strconv.Itoa(page))
	query.Add("pageSize", strconv.Itoa(pageSize))
	body, err := c.GetRequest("campaigns", query)
	if err != nil {
		return err
	}
	var response CampaignsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	c.campaigns = response.Campaigns
	return nil
}

type CampaignsResponse struct {
	Campaigns []*models.CampaignDTO
	Pager     *models.FlippingPagerDTO
}

func (c *Client) GetRequest(resource string, query url.Values) ([]byte, error) {
	u := c.baseUrl.JoinPath(resource)
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.token)
	req.Header.Add("Accept-Encoding", "application/json")
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
