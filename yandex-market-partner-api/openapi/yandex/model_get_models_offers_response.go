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

// checks if the GetModelsOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetModelsOffersResponse{}

// GetModelsOffersResponse Ответ на запрос списка предложений для моделей.
type GetModelsOffersResponse struct {
	// Список моделей товаров.
	Models []EnrichedModelDTO `json:"models"`
	Currency *CurrencyType `json:"currency,omitempty"`
	// Идентификатор региона, для которого выводится информация о предложениях модели (доставляемых в этот регион).  Информацию о регионе по идентификатору можно получить с помощью запроса [GET regions/{regionId}](../../reference/regions/searchRegionsById.md). 
	RegionId *int64 `json:"regionId,omitempty"`
}

type _GetModelsOffersResponse GetModelsOffersResponse

// NewGetModelsOffersResponse instantiates a new GetModelsOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetModelsOffersResponse(models []EnrichedModelDTO) *GetModelsOffersResponse {
	this := GetModelsOffersResponse{}
	this.Models = models
	return &this
}

// NewGetModelsOffersResponseWithDefaults instantiates a new GetModelsOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetModelsOffersResponseWithDefaults() *GetModelsOffersResponse {
	this := GetModelsOffersResponse{}
	return &this
}

// GetModels returns the Models field value
func (o *GetModelsOffersResponse) GetModels() []EnrichedModelDTO {
	if o == nil {
		var ret []EnrichedModelDTO
		return ret
	}

	return o.Models
}

// GetModelsOk returns a tuple with the Models field value
// and a boolean to check if the value has been set.
func (o *GetModelsOffersResponse) GetModelsOk() ([]EnrichedModelDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Models, true
}

// SetModels sets field value
func (o *GetModelsOffersResponse) SetModels(v []EnrichedModelDTO) {
	o.Models = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetModelsOffersResponse) GetCurrency() CurrencyType {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyType
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetModelsOffersResponse) GetCurrencyOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetModelsOffersResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyType and assigns it to the Currency field.
func (o *GetModelsOffersResponse) SetCurrency(v CurrencyType) {
	o.Currency = &v
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *GetModelsOffersResponse) GetRegionId() int64 {
	if o == nil || IsNil(o.RegionId) {
		var ret int64
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetModelsOffersResponse) GetRegionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *GetModelsOffersResponse) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int64 and assigns it to the RegionId field.
func (o *GetModelsOffersResponse) SetRegionId(v int64) {
	o.RegionId = &v
}

func (o GetModelsOffersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetModelsOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["models"] = o.Models
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.RegionId) {
		toSerialize["regionId"] = o.RegionId
	}
	return toSerialize, nil
}

func (o *GetModelsOffersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"models",
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

	varGetModelsOffersResponse := _GetModelsOffersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetModelsOffersResponse)

	if err != nil {
		return err
	}

	*o = GetModelsOffersResponse(varGetModelsOffersResponse)

	return err
}

type NullableGetModelsOffersResponse struct {
	value *GetModelsOffersResponse
	isSet bool
}

func (v NullableGetModelsOffersResponse) Get() *GetModelsOffersResponse {
	return v.value
}

func (v *NullableGetModelsOffersResponse) Set(val *GetModelsOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetModelsOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetModelsOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetModelsOffersResponse(val *GetModelsOffersResponse) *NullableGetModelsOffersResponse {
	return &NullableGetModelsOffersResponse{value: val, isSet: true}
}

func (v NullableGetModelsOffersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetModelsOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


