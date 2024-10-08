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

// checks if the FeedPlacementDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedPlacementDTO{}

// FeedPlacementDTO Информация о размещении предложений из прайс-листа на Маркете на момент выполнения запроса.
type FeedPlacementDTO struct {
	Status *FeedStatusType `json:"status,omitempty"`
	// Количество предложений из прайс-листа, которые размещаются на Яндекс Маркете в момент выполнения запроса.
	TotalOffersCount *int32 `json:"totalOffersCount,omitempty"`
}

// NewFeedPlacementDTO instantiates a new FeedPlacementDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedPlacementDTO() *FeedPlacementDTO {
	this := FeedPlacementDTO{}
	return &this
}

// NewFeedPlacementDTOWithDefaults instantiates a new FeedPlacementDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedPlacementDTOWithDefaults() *FeedPlacementDTO {
	this := FeedPlacementDTO{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FeedPlacementDTO) GetStatus() FeedStatusType {
	if o == nil || IsNil(o.Status) {
		var ret FeedStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedPlacementDTO) GetStatusOk() (*FeedStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FeedPlacementDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given FeedStatusType and assigns it to the Status field.
func (o *FeedPlacementDTO) SetStatus(v FeedStatusType) {
	o.Status = &v
}

// GetTotalOffersCount returns the TotalOffersCount field value if set, zero value otherwise.
func (o *FeedPlacementDTO) GetTotalOffersCount() int32 {
	if o == nil || IsNil(o.TotalOffersCount) {
		var ret int32
		return ret
	}
	return *o.TotalOffersCount
}

// GetTotalOffersCountOk returns a tuple with the TotalOffersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedPlacementDTO) GetTotalOffersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOffersCount) {
		return nil, false
	}
	return o.TotalOffersCount, true
}

// HasTotalOffersCount returns a boolean if a field has been set.
func (o *FeedPlacementDTO) HasTotalOffersCount() bool {
	if o != nil && !IsNil(o.TotalOffersCount) {
		return true
	}

	return false
}

// SetTotalOffersCount gets a reference to the given int32 and assigns it to the TotalOffersCount field.
func (o *FeedPlacementDTO) SetTotalOffersCount(v int32) {
	o.TotalOffersCount = &v
}

func (o FeedPlacementDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedPlacementDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalOffersCount) {
		toSerialize["totalOffersCount"] = o.TotalOffersCount
	}
	return toSerialize, nil
}

type NullableFeedPlacementDTO struct {
	value *FeedPlacementDTO
	isSet bool
}

func (v NullableFeedPlacementDTO) Get() *FeedPlacementDTO {
	return v.value
}

func (v *NullableFeedPlacementDTO) Set(val *FeedPlacementDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedPlacementDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedPlacementDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedPlacementDTO(val *FeedPlacementDTO) *NullableFeedPlacementDTO {
	return &NullableFeedPlacementDTO{value: val, isSet: true}
}

func (v NullableFeedPlacementDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedPlacementDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


