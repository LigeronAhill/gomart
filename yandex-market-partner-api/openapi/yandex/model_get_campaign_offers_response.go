/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetCampaignOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignOffersResponse{}

// GetCampaignOffersResponse Ответ на запрос списка товаров в магазине.
type GetCampaignOffersResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *GetCampaignOffersResultDTO `json:"result,omitempty"`
}

// NewGetCampaignOffersResponse instantiates a new GetCampaignOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignOffersResponse() *GetCampaignOffersResponse {
	this := GetCampaignOffersResponse{}
	return &this
}

// NewGetCampaignOffersResponseWithDefaults instantiates a new GetCampaignOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignOffersResponseWithDefaults() *GetCampaignOffersResponse {
	this := GetCampaignOffersResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCampaignOffersResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignOffersResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCampaignOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetCampaignOffersResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetCampaignOffersResponse) GetResult() GetCampaignOffersResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret GetCampaignOffersResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignOffersResponse) GetResultOk() (*GetCampaignOffersResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetCampaignOffersResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetCampaignOffersResultDTO and assigns it to the Result field.
func (o *GetCampaignOffersResponse) SetResult(v GetCampaignOffersResultDTO) {
	o.Result = &v
}

func (o GetCampaignOffersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetCampaignOffersResponse struct {
	value *GetCampaignOffersResponse
	isSet bool
}

func (v NullableGetCampaignOffersResponse) Get() *GetCampaignOffersResponse {
	return v.value
}

func (v *NullableGetCampaignOffersResponse) Set(val *GetCampaignOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignOffersResponse(val *GetCampaignOffersResponse) *NullableGetCampaignOffersResponse {
	return &NullableGetCampaignOffersResponse{value: val, isSet: true}
}

func (v NullableGetCampaignOffersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


