/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GetPromoBestsellerInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPromoBestsellerInfoDTO{}

// GetPromoBestsellerInfoDTO Информация об акции «Бестселлеры Маркета».
type GetPromoBestsellerInfoDTO struct {
	// Является ли акция «Бестселлером Маркета». Подробнее об этой акции читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/marketing/promos/market/bestsellers).
	Bestseller bool `json:"bestseller"`
	// До какой даты можно добавить товар в акцию «Бестселлеры Маркета».  Параметр возвращается только для текущих и будущих акций «Бестселлеры Маркета». 
	EntryDeadline *time.Time `json:"entryDeadline,omitempty"`
}

type _GetPromoBestsellerInfoDTO GetPromoBestsellerInfoDTO

// NewGetPromoBestsellerInfoDTO instantiates a new GetPromoBestsellerInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPromoBestsellerInfoDTO(bestseller bool) *GetPromoBestsellerInfoDTO {
	this := GetPromoBestsellerInfoDTO{}
	this.Bestseller = bestseller
	return &this
}

// NewGetPromoBestsellerInfoDTOWithDefaults instantiates a new GetPromoBestsellerInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPromoBestsellerInfoDTOWithDefaults() *GetPromoBestsellerInfoDTO {
	this := GetPromoBestsellerInfoDTO{}
	return &this
}

// GetBestseller returns the Bestseller field value
func (o *GetPromoBestsellerInfoDTO) GetBestseller() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bestseller
}

// GetBestsellerOk returns a tuple with the Bestseller field value
// and a boolean to check if the value has been set.
func (o *GetPromoBestsellerInfoDTO) GetBestsellerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bestseller, true
}

// SetBestseller sets field value
func (o *GetPromoBestsellerInfoDTO) SetBestseller(v bool) {
	o.Bestseller = v
}

// GetEntryDeadline returns the EntryDeadline field value if set, zero value otherwise.
func (o *GetPromoBestsellerInfoDTO) GetEntryDeadline() time.Time {
	if o == nil || IsNil(o.EntryDeadline) {
		var ret time.Time
		return ret
	}
	return *o.EntryDeadline
}

// GetEntryDeadlineOk returns a tuple with the EntryDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromoBestsellerInfoDTO) GetEntryDeadlineOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EntryDeadline) {
		return nil, false
	}
	return o.EntryDeadline, true
}

// HasEntryDeadline returns a boolean if a field has been set.
func (o *GetPromoBestsellerInfoDTO) HasEntryDeadline() bool {
	if o != nil && !IsNil(o.EntryDeadline) {
		return true
	}

	return false
}

// SetEntryDeadline gets a reference to the given time.Time and assigns it to the EntryDeadline field.
func (o *GetPromoBestsellerInfoDTO) SetEntryDeadline(v time.Time) {
	o.EntryDeadline = &v
}

func (o GetPromoBestsellerInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPromoBestsellerInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bestseller"] = o.Bestseller
	if !IsNil(o.EntryDeadline) {
		toSerialize["entryDeadline"] = o.EntryDeadline
	}
	return toSerialize, nil
}

func (o *GetPromoBestsellerInfoDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bestseller",
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

	varGetPromoBestsellerInfoDTO := _GetPromoBestsellerInfoDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPromoBestsellerInfoDTO)

	if err != nil {
		return err
	}

	*o = GetPromoBestsellerInfoDTO(varGetPromoBestsellerInfoDTO)

	return err
}

type NullableGetPromoBestsellerInfoDTO struct {
	value *GetPromoBestsellerInfoDTO
	isSet bool
}

func (v NullableGetPromoBestsellerInfoDTO) Get() *GetPromoBestsellerInfoDTO {
	return v.value
}

func (v *NullableGetPromoBestsellerInfoDTO) Set(val *GetPromoBestsellerInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPromoBestsellerInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPromoBestsellerInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPromoBestsellerInfoDTO(val *GetPromoBestsellerInfoDTO) *NullableGetPromoBestsellerInfoDTO {
	return &NullableGetPromoBestsellerInfoDTO{value: val, isSet: true}
}

func (v NullableGetPromoBestsellerInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPromoBestsellerInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


