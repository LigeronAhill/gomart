/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignsResponse{}

// GetCampaignsResponse Результаты поиска магазинов.
type GetCampaignsResponse struct {
	// Список с информацией по каждому магазину.
	Campaigns []CampaignDTO `json:"campaigns"`
	Pager *FlippingPagerDTO `json:"pager,omitempty"`
}

type _GetCampaignsResponse GetCampaignsResponse

// NewGetCampaignsResponse instantiates a new GetCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignsResponse(campaigns []CampaignDTO) *GetCampaignsResponse {
	this := GetCampaignsResponse{}
	this.Campaigns = campaigns
	return &this
}

// NewGetCampaignsResponseWithDefaults instantiates a new GetCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignsResponseWithDefaults() *GetCampaignsResponse {
	this := GetCampaignsResponse{}
	return &this
}

// GetCampaigns returns the Campaigns field value
func (o *GetCampaignsResponse) GetCampaigns() []CampaignDTO {
	if o == nil {
		var ret []CampaignDTO
		return ret
	}

	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value
// and a boolean to check if the value has been set.
func (o *GetCampaignsResponse) GetCampaignsOk() ([]CampaignDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaigns, true
}

// SetCampaigns sets field value
func (o *GetCampaignsResponse) SetCampaigns(v []CampaignDTO) {
	o.Campaigns = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *GetCampaignsResponse) GetPager() FlippingPagerDTO {
	if o == nil || IsNil(o.Pager) {
		var ret FlippingPagerDTO
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignsResponse) GetPagerOk() (*FlippingPagerDTO, bool) {
	if o == nil || IsNil(o.Pager) {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *GetCampaignsResponse) HasPager() bool {
	if o != nil && !IsNil(o.Pager) {
		return true
	}

	return false
}

// SetPager gets a reference to the given FlippingPagerDTO and assigns it to the Pager field.
func (o *GetCampaignsResponse) SetPager(v FlippingPagerDTO) {
	o.Pager = &v
}

func (o GetCampaignsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaigns"] = o.Campaigns
	if !IsNil(o.Pager) {
		toSerialize["pager"] = o.Pager
	}
	return toSerialize, nil
}

func (o *GetCampaignsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaigns",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetCampaignsResponse := _GetCampaignsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignsResponse)

	if err != nil {
		return err
	}

	*o = GetCampaignsResponse(varGetCampaignsResponse)

	return err
}

type NullableGetCampaignsResponse struct {
	value *GetCampaignsResponse
	isSet bool
}

func (v NullableGetCampaignsResponse) Get() *GetCampaignsResponse {
	return v.value
}

func (v *NullableGetCampaignsResponse) Set(val *GetCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignsResponse(val *GetCampaignsResponse) *NullableGetCampaignsResponse {
	return &NullableGetCampaignsResponse{value: val, isSet: true}
}

func (v NullableGetCampaignsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


