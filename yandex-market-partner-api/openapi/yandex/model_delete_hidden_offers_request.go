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

// checks if the DeleteHiddenOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteHiddenOffersRequest{}

// DeleteHiddenOffersRequest Запрос на возобновление показа оферов.
type DeleteHiddenOffersRequest struct {
	// Список скрытых товаров. 
	HiddenOffers []HiddenOfferDTO `json:"hiddenOffers"`
}

type _DeleteHiddenOffersRequest DeleteHiddenOffersRequest

// NewDeleteHiddenOffersRequest instantiates a new DeleteHiddenOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteHiddenOffersRequest(hiddenOffers []HiddenOfferDTO) *DeleteHiddenOffersRequest {
	this := DeleteHiddenOffersRequest{}
	this.HiddenOffers = hiddenOffers
	return &this
}

// NewDeleteHiddenOffersRequestWithDefaults instantiates a new DeleteHiddenOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteHiddenOffersRequestWithDefaults() *DeleteHiddenOffersRequest {
	this := DeleteHiddenOffersRequest{}
	return &this
}

// GetHiddenOffers returns the HiddenOffers field value
func (o *DeleteHiddenOffersRequest) GetHiddenOffers() []HiddenOfferDTO {
	if o == nil {
		var ret []HiddenOfferDTO
		return ret
	}

	return o.HiddenOffers
}

// GetHiddenOffersOk returns a tuple with the HiddenOffers field value
// and a boolean to check if the value has been set.
func (o *DeleteHiddenOffersRequest) GetHiddenOffersOk() ([]HiddenOfferDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.HiddenOffers, true
}

// SetHiddenOffers sets field value
func (o *DeleteHiddenOffersRequest) SetHiddenOffers(v []HiddenOfferDTO) {
	o.HiddenOffers = v
}

func (o DeleteHiddenOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteHiddenOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hiddenOffers"] = o.HiddenOffers
	return toSerialize, nil
}

func (o *DeleteHiddenOffersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hiddenOffers",
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

	varDeleteHiddenOffersRequest := _DeleteHiddenOffersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteHiddenOffersRequest)

	if err != nil {
		return err
	}

	*o = DeleteHiddenOffersRequest(varDeleteHiddenOffersRequest)

	return err
}

type NullableDeleteHiddenOffersRequest struct {
	value *DeleteHiddenOffersRequest
	isSet bool
}

func (v NullableDeleteHiddenOffersRequest) Get() *DeleteHiddenOffersRequest {
	return v.value
}

func (v *NullableDeleteHiddenOffersRequest) Set(val *DeleteHiddenOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteHiddenOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteHiddenOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteHiddenOffersRequest(val *DeleteHiddenOffersRequest) *NullableDeleteHiddenOffersRequest {
	return &NullableDeleteHiddenOffersRequest{value: val, isSet: true}
}

func (v NullableDeleteHiddenOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteHiddenOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


