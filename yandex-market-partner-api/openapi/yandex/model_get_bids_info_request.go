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

// checks if the GetBidsInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBidsInfoRequest{}

// GetBidsInfoRequest description.
type GetBidsInfoRequest struct {
	// Список товаров, для которых нужно получить значения ставок.  Если список не задан, постранично возвращаются все товары со ставками.  Если список задан, результаты возвращаются одной страницей, а параметры `page_token` и `limit` игнорируются. 
	Skus []string `json:"skus,omitempty"`
}

// NewGetBidsInfoRequest instantiates a new GetBidsInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBidsInfoRequest() *GetBidsInfoRequest {
	this := GetBidsInfoRequest{}
	return &this
}

// NewGetBidsInfoRequestWithDefaults instantiates a new GetBidsInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBidsInfoRequestWithDefaults() *GetBidsInfoRequest {
	this := GetBidsInfoRequest{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetBidsInfoRequest) GetSkus() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBidsInfoRequest) GetSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GetBidsInfoRequest) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given []string and assigns it to the Skus field.
func (o *GetBidsInfoRequest) SetSkus(v []string) {
	o.Skus = v
}

func (o GetBidsInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBidsInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	return toSerialize, nil
}

type NullableGetBidsInfoRequest struct {
	value *GetBidsInfoRequest
	isSet bool
}

func (v NullableGetBidsInfoRequest) Get() *GetBidsInfoRequest {
	return v.value
}

func (v *NullableGetBidsInfoRequest) Set(val *GetBidsInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBidsInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBidsInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBidsInfoRequest(val *GetBidsInfoRequest) *NullableGetBidsInfoRequest {
	return &NullableGetBidsInfoRequest{value: val, isSet: true}
}

func (v NullableGetBidsInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBidsInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


