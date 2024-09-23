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

// checks if the DeleteOffersFromArchiveDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteOffersFromArchiveDTO{}

// DeleteOffersFromArchiveDTO Товары, которые не удалось восстановить из архива.
type DeleteOffersFromArchiveDTO struct {
	// Список товаров, которые не удалось восстановить из архива.
	NotUnarchivedOfferIds []string `json:"notUnarchivedOfferIds,omitempty"`
}

// NewDeleteOffersFromArchiveDTO instantiates a new DeleteOffersFromArchiveDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteOffersFromArchiveDTO() *DeleteOffersFromArchiveDTO {
	this := DeleteOffersFromArchiveDTO{}
	return &this
}

// NewDeleteOffersFromArchiveDTOWithDefaults instantiates a new DeleteOffersFromArchiveDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteOffersFromArchiveDTOWithDefaults() *DeleteOffersFromArchiveDTO {
	this := DeleteOffersFromArchiveDTO{}
	return &this
}

// GetNotUnarchivedOfferIds returns the NotUnarchivedOfferIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteOffersFromArchiveDTO) GetNotUnarchivedOfferIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotUnarchivedOfferIds
}

// GetNotUnarchivedOfferIdsOk returns a tuple with the NotUnarchivedOfferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteOffersFromArchiveDTO) GetNotUnarchivedOfferIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotUnarchivedOfferIds) {
		return nil, false
	}
	return o.NotUnarchivedOfferIds, true
}

// HasNotUnarchivedOfferIds returns a boolean if a field has been set.
func (o *DeleteOffersFromArchiveDTO) HasNotUnarchivedOfferIds() bool {
	if o != nil && !IsNil(o.NotUnarchivedOfferIds) {
		return true
	}

	return false
}

// SetNotUnarchivedOfferIds gets a reference to the given []string and assigns it to the NotUnarchivedOfferIds field.
func (o *DeleteOffersFromArchiveDTO) SetNotUnarchivedOfferIds(v []string) {
	o.NotUnarchivedOfferIds = v
}

func (o DeleteOffersFromArchiveDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOffersFromArchiveDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NotUnarchivedOfferIds != nil {
		toSerialize["notUnarchivedOfferIds"] = o.NotUnarchivedOfferIds
	}
	return toSerialize, nil
}

type NullableDeleteOffersFromArchiveDTO struct {
	value *DeleteOffersFromArchiveDTO
	isSet bool
}

func (v NullableDeleteOffersFromArchiveDTO) Get() *DeleteOffersFromArchiveDTO {
	return v.value
}

func (v *NullableDeleteOffersFromArchiveDTO) Set(val *DeleteOffersFromArchiveDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOffersFromArchiveDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOffersFromArchiveDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOffersFromArchiveDTO(val *DeleteOffersFromArchiveDTO) *NullableDeleteOffersFromArchiveDTO {
	return &NullableDeleteOffersFromArchiveDTO{value: val, isSet: true}
}

func (v NullableDeleteOffersFromArchiveDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOffersFromArchiveDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


