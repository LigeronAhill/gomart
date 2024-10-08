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

// checks if the GetPromoAssortmentInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPromoAssortmentInfoDTO{}

// GetPromoAssortmentInfoDTO Информация о товарах в акции.
type GetPromoAssortmentInfoDTO struct {
	// Количество товаров, которые участвуют или участвовали в акции.  Учитываются только товары, которые были добавлены вручную.  Об автоматическом и ручном добавлении товаров в акцию читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/marketing/promos/market/index). 
	ActiveOffers int32 `json:"activeOffers"`
	// Количество доступных товаров в акции.  Параметр возвращается только для текущих и будущих акций. 
	PotentialOffers *int32 `json:"potentialOffers,omitempty"`
	// Есть ли изменения в ассортименте, которые еще не применились. Сохранение изменений занимает некоторое время.  Параметр возвращается только для текущих и будущих акций. 
	Processing *bool `json:"processing,omitempty"`
}

type _GetPromoAssortmentInfoDTO GetPromoAssortmentInfoDTO

// NewGetPromoAssortmentInfoDTO instantiates a new GetPromoAssortmentInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPromoAssortmentInfoDTO(activeOffers int32) *GetPromoAssortmentInfoDTO {
	this := GetPromoAssortmentInfoDTO{}
	this.ActiveOffers = activeOffers
	return &this
}

// NewGetPromoAssortmentInfoDTOWithDefaults instantiates a new GetPromoAssortmentInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPromoAssortmentInfoDTOWithDefaults() *GetPromoAssortmentInfoDTO {
	this := GetPromoAssortmentInfoDTO{}
	return &this
}

// GetActiveOffers returns the ActiveOffers field value
func (o *GetPromoAssortmentInfoDTO) GetActiveOffers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ActiveOffers
}

// GetActiveOffersOk returns a tuple with the ActiveOffers field value
// and a boolean to check if the value has been set.
func (o *GetPromoAssortmentInfoDTO) GetActiveOffersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveOffers, true
}

// SetActiveOffers sets field value
func (o *GetPromoAssortmentInfoDTO) SetActiveOffers(v int32) {
	o.ActiveOffers = v
}

// GetPotentialOffers returns the PotentialOffers field value if set, zero value otherwise.
func (o *GetPromoAssortmentInfoDTO) GetPotentialOffers() int32 {
	if o == nil || IsNil(o.PotentialOffers) {
		var ret int32
		return ret
	}
	return *o.PotentialOffers
}

// GetPotentialOffersOk returns a tuple with the PotentialOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromoAssortmentInfoDTO) GetPotentialOffersOk() (*int32, bool) {
	if o == nil || IsNil(o.PotentialOffers) {
		return nil, false
	}
	return o.PotentialOffers, true
}

// HasPotentialOffers returns a boolean if a field has been set.
func (o *GetPromoAssortmentInfoDTO) HasPotentialOffers() bool {
	if o != nil && !IsNil(o.PotentialOffers) {
		return true
	}

	return false
}

// SetPotentialOffers gets a reference to the given int32 and assigns it to the PotentialOffers field.
func (o *GetPromoAssortmentInfoDTO) SetPotentialOffers(v int32) {
	o.PotentialOffers = &v
}

// GetProcessing returns the Processing field value if set, zero value otherwise.
func (o *GetPromoAssortmentInfoDTO) GetProcessing() bool {
	if o == nil || IsNil(o.Processing) {
		var ret bool
		return ret
	}
	return *o.Processing
}

// GetProcessingOk returns a tuple with the Processing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromoAssortmentInfoDTO) GetProcessingOk() (*bool, bool) {
	if o == nil || IsNil(o.Processing) {
		return nil, false
	}
	return o.Processing, true
}

// HasProcessing returns a boolean if a field has been set.
func (o *GetPromoAssortmentInfoDTO) HasProcessing() bool {
	if o != nil && !IsNil(o.Processing) {
		return true
	}

	return false
}

// SetProcessing gets a reference to the given bool and assigns it to the Processing field.
func (o *GetPromoAssortmentInfoDTO) SetProcessing(v bool) {
	o.Processing = &v
}

func (o GetPromoAssortmentInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPromoAssortmentInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeOffers"] = o.ActiveOffers
	if !IsNil(o.PotentialOffers) {
		toSerialize["potentialOffers"] = o.PotentialOffers
	}
	if !IsNil(o.Processing) {
		toSerialize["processing"] = o.Processing
	}
	return toSerialize, nil
}

func (o *GetPromoAssortmentInfoDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activeOffers",
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

	varGetPromoAssortmentInfoDTO := _GetPromoAssortmentInfoDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPromoAssortmentInfoDTO)

	if err != nil {
		return err
	}

	*o = GetPromoAssortmentInfoDTO(varGetPromoAssortmentInfoDTO)

	return err
}

type NullableGetPromoAssortmentInfoDTO struct {
	value *GetPromoAssortmentInfoDTO
	isSet bool
}

func (v NullableGetPromoAssortmentInfoDTO) Get() *GetPromoAssortmentInfoDTO {
	return v.value
}

func (v *NullableGetPromoAssortmentInfoDTO) Set(val *GetPromoAssortmentInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPromoAssortmentInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPromoAssortmentInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPromoAssortmentInfoDTO(val *GetPromoAssortmentInfoDTO) *NullableGetPromoAssortmentInfoDTO {
	return &NullableGetPromoAssortmentInfoDTO{value: val, isSet: true}
}

func (v NullableGetPromoAssortmentInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPromoAssortmentInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


