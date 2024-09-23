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

// checks if the GetBidsRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBidsRecommendationsRequest{}

// GetBidsRecommendationsRequest description.
type GetBidsRecommendationsRequest struct {
	// Список товаров, для которых нужно получить рекомендации по ставкам. 
	Skus []string `json:"skus"`
}

type _GetBidsRecommendationsRequest GetBidsRecommendationsRequest

// NewGetBidsRecommendationsRequest instantiates a new GetBidsRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBidsRecommendationsRequest(skus []string) *GetBidsRecommendationsRequest {
	this := GetBidsRecommendationsRequest{}
	this.Skus = skus
	return &this
}

// NewGetBidsRecommendationsRequestWithDefaults instantiates a new GetBidsRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBidsRecommendationsRequestWithDefaults() *GetBidsRecommendationsRequest {
	this := GetBidsRecommendationsRequest{}
	return &this
}

// GetSkus returns the Skus field value
func (o *GetBidsRecommendationsRequest) GetSkus() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value
// and a boolean to check if the value has been set.
func (o *GetBidsRecommendationsRequest) GetSkusOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skus, true
}

// SetSkus sets field value
func (o *GetBidsRecommendationsRequest) SetSkus(v []string) {
	o.Skus = v
}

func (o GetBidsRecommendationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBidsRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["skus"] = o.Skus
	return toSerialize, nil
}

func (o *GetBidsRecommendationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"skus",
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

	varGetBidsRecommendationsRequest := _GetBidsRecommendationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetBidsRecommendationsRequest)

	if err != nil {
		return err
	}

	*o = GetBidsRecommendationsRequest(varGetBidsRecommendationsRequest)

	return err
}

type NullableGetBidsRecommendationsRequest struct {
	value *GetBidsRecommendationsRequest
	isSet bool
}

func (v NullableGetBidsRecommendationsRequest) Get() *GetBidsRecommendationsRequest {
	return v.value
}

func (v *NullableGetBidsRecommendationsRequest) Set(val *GetBidsRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBidsRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBidsRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBidsRecommendationsRequest(val *GetBidsRecommendationsRequest) *NullableGetBidsRecommendationsRequest {
	return &NullableGetBidsRecommendationsRequest{value: val, isSet: true}
}

func (v NullableGetBidsRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBidsRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


