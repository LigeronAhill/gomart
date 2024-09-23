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

// checks if the GenerateGoodsTurnoverRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateGoodsTurnoverRequest{}

// GenerateGoodsTurnoverRequest Данные, необходимые для генерации отчета. 
type GenerateGoodsTurnoverRequest struct {
	// Идентификатор кампании.
	CampaignId int64 `json:"campaignId"`
	// Дата, за которую нужно рассчитать оборачиваемость. Если параметр не указан, используется текущая дата.
	Date *string `json:"date,omitempty"`
}

type _GenerateGoodsTurnoverRequest GenerateGoodsTurnoverRequest

// NewGenerateGoodsTurnoverRequest instantiates a new GenerateGoodsTurnoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateGoodsTurnoverRequest(campaignId int64) *GenerateGoodsTurnoverRequest {
	this := GenerateGoodsTurnoverRequest{}
	this.CampaignId = campaignId
	return &this
}

// NewGenerateGoodsTurnoverRequestWithDefaults instantiates a new GenerateGoodsTurnoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateGoodsTurnoverRequestWithDefaults() *GenerateGoodsTurnoverRequest {
	this := GenerateGoodsTurnoverRequest{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *GenerateGoodsTurnoverRequest) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *GenerateGoodsTurnoverRequest) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *GenerateGoodsTurnoverRequest) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GenerateGoodsTurnoverRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateGoodsTurnoverRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GenerateGoodsTurnoverRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GenerateGoodsTurnoverRequest) SetDate(v string) {
	o.Date = &v
}

func (o GenerateGoodsTurnoverRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateGoodsTurnoverRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

func (o *GenerateGoodsTurnoverRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignId",
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

	varGenerateGoodsTurnoverRequest := _GenerateGoodsTurnoverRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateGoodsTurnoverRequest)

	if err != nil {
		return err
	}

	*o = GenerateGoodsTurnoverRequest(varGenerateGoodsTurnoverRequest)

	return err
}

type NullableGenerateGoodsTurnoverRequest struct {
	value *GenerateGoodsTurnoverRequest
	isSet bool
}

func (v NullableGenerateGoodsTurnoverRequest) Get() *GenerateGoodsTurnoverRequest {
	return v.value
}

func (v *NullableGenerateGoodsTurnoverRequest) Set(val *GenerateGoodsTurnoverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateGoodsTurnoverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateGoodsTurnoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateGoodsTurnoverRequest(val *GenerateGoodsTurnoverRequest) *NullableGenerateGoodsTurnoverRequest {
	return &NullableGenerateGoodsTurnoverRequest{value: val, isSet: true}
}

func (v NullableGenerateGoodsTurnoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateGoodsTurnoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


