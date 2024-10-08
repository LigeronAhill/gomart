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

// checks if the GetPromoOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPromoOffersRequest{}

// GetPromoOffersRequest Получение списка товаров, которые участвуют или могут участвовать в акции.
type GetPromoOffersRequest struct {
	// Идентификатор акции.
	PromoId string `json:"promoId"`
	StatusType *PromoOfferParticipationStatusFilterType `json:"statusType,omitempty"`
}

type _GetPromoOffersRequest GetPromoOffersRequest

// NewGetPromoOffersRequest instantiates a new GetPromoOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPromoOffersRequest(promoId string) *GetPromoOffersRequest {
	this := GetPromoOffersRequest{}
	this.PromoId = promoId
	return &this
}

// NewGetPromoOffersRequestWithDefaults instantiates a new GetPromoOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPromoOffersRequestWithDefaults() *GetPromoOffersRequest {
	this := GetPromoOffersRequest{}
	return &this
}

// GetPromoId returns the PromoId field value
func (o *GetPromoOffersRequest) GetPromoId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromoId
}

// GetPromoIdOk returns a tuple with the PromoId field value
// and a boolean to check if the value has been set.
func (o *GetPromoOffersRequest) GetPromoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromoId, true
}

// SetPromoId sets field value
func (o *GetPromoOffersRequest) SetPromoId(v string) {
	o.PromoId = v
}

// GetStatusType returns the StatusType field value if set, zero value otherwise.
func (o *GetPromoOffersRequest) GetStatusType() PromoOfferParticipationStatusFilterType {
	if o == nil || IsNil(o.StatusType) {
		var ret PromoOfferParticipationStatusFilterType
		return ret
	}
	return *o.StatusType
}

// GetStatusTypeOk returns a tuple with the StatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromoOffersRequest) GetStatusTypeOk() (*PromoOfferParticipationStatusFilterType, bool) {
	if o == nil || IsNil(o.StatusType) {
		return nil, false
	}
	return o.StatusType, true
}

// HasStatusType returns a boolean if a field has been set.
func (o *GetPromoOffersRequest) HasStatusType() bool {
	if o != nil && !IsNil(o.StatusType) {
		return true
	}

	return false
}

// SetStatusType gets a reference to the given PromoOfferParticipationStatusFilterType and assigns it to the StatusType field.
func (o *GetPromoOffersRequest) SetStatusType(v PromoOfferParticipationStatusFilterType) {
	o.StatusType = &v
}

func (o GetPromoOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPromoOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["promoId"] = o.PromoId
	if !IsNil(o.StatusType) {
		toSerialize["statusType"] = o.StatusType
	}
	return toSerialize, nil
}

func (o *GetPromoOffersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"promoId",
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

	varGetPromoOffersRequest := _GetPromoOffersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPromoOffersRequest)

	if err != nil {
		return err
	}

	*o = GetPromoOffersRequest(varGetPromoOffersRequest)

	return err
}

type NullableGetPromoOffersRequest struct {
	value *GetPromoOffersRequest
	isSet bool
}

func (v NullableGetPromoOffersRequest) Get() *GetPromoOffersRequest {
	return v.value
}

func (v *NullableGetPromoOffersRequest) Set(val *GetPromoOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPromoOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPromoOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPromoOffersRequest(val *GetPromoOffersRequest) *NullableGetPromoOffersRequest {
	return &NullableGetPromoOffersRequest{value: val, isSet: true}
}

func (v NullableGetPromoOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPromoOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


