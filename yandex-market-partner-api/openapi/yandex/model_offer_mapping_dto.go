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

// checks if the OfferMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferMappingDTO{}

// OfferMappingDTO Информация о текущей карточке товара на Маркете.
type OfferMappingDTO struct {
	// SKU на Маркете.
	MarketSku *int64 `json:"marketSku,omitempty"`
	// Идентификатор модели для текущей карточки товара на Маркете.  Например, две лопатки разных цветов имеют разные SKU на Маркете (параметр `marketSku`), но одинаковый идентификатор модели товара. 
	ModelId *int64 `json:"modelId,omitempty"`
	// Идентификатор категории для текущей карточки товара на Маркете.
	CategoryId *int64 `json:"categoryId,omitempty"`
}

// NewOfferMappingDTO instantiates a new OfferMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferMappingDTO() *OfferMappingDTO {
	this := OfferMappingDTO{}
	return &this
}

// NewOfferMappingDTOWithDefaults instantiates a new OfferMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferMappingDTOWithDefaults() *OfferMappingDTO {
	this := OfferMappingDTO{}
	return &this
}

// GetMarketSku returns the MarketSku field value if set, zero value otherwise.
func (o *OfferMappingDTO) GetMarketSku() int64 {
	if o == nil || IsNil(o.MarketSku) {
		var ret int64
		return ret
	}
	return *o.MarketSku
}

// GetMarketSkuOk returns a tuple with the MarketSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingDTO) GetMarketSkuOk() (*int64, bool) {
	if o == nil || IsNil(o.MarketSku) {
		return nil, false
	}
	return o.MarketSku, true
}

// HasMarketSku returns a boolean if a field has been set.
func (o *OfferMappingDTO) HasMarketSku() bool {
	if o != nil && !IsNil(o.MarketSku) {
		return true
	}

	return false
}

// SetMarketSku gets a reference to the given int64 and assigns it to the MarketSku field.
func (o *OfferMappingDTO) SetMarketSku(v int64) {
	o.MarketSku = &v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *OfferMappingDTO) GetModelId() int64 {
	if o == nil || IsNil(o.ModelId) {
		var ret int64
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingDTO) GetModelIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *OfferMappingDTO) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given int64 and assigns it to the ModelId field.
func (o *OfferMappingDTO) SetModelId(v int64) {
	o.ModelId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *OfferMappingDTO) GetCategoryId() int64 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int64
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingDTO) GetCategoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *OfferMappingDTO) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int64 and assigns it to the CategoryId field.
func (o *OfferMappingDTO) SetCategoryId(v int64) {
	o.CategoryId = &v
}

func (o OfferMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketSku) {
		toSerialize["marketSku"] = o.MarketSku
	}
	if !IsNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	if !IsNil(o.CategoryId) {
		toSerialize["categoryId"] = o.CategoryId
	}
	return toSerialize, nil
}

type NullableOfferMappingDTO struct {
	value *OfferMappingDTO
	isSet bool
}

func (v NullableOfferMappingDTO) Get() *OfferMappingDTO {
	return v.value
}

func (v *NullableOfferMappingDTO) Set(val *OfferMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferMappingDTO(val *OfferMappingDTO) *NullableOfferMappingDTO {
	return &NullableOfferMappingDTO{value: val, isSet: true}
}

func (v NullableOfferMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


