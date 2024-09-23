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

// checks if the WarningPromoOfferUpdateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningPromoOfferUpdateDTO{}

// WarningPromoOfferUpdateDTO Описание предупреждения, которое появилось при добавлении товара.
type WarningPromoOfferUpdateDTO struct {
	// Ваш SKU — идентификатор товара в вашей системе.  Разрешена любая последовательность длиной до 255 знаков.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * SKU товара нельзя менять — можно только удалить товар и добавить заново с новым SKU.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields) 
	OfferId string `json:"offerId" validate:"regexp=^[^\\\\x00-\\\\x08\\\\x0A-\\\\x1f\\\\x7f]{1,255}$"`
	// Предупреждения, которые появились при добавлении товара в акцию или изменении его цен.
	Warnings []PromoOfferUpdateWarningDTO `json:"warnings"`
}

type _WarningPromoOfferUpdateDTO WarningPromoOfferUpdateDTO

// NewWarningPromoOfferUpdateDTO instantiates a new WarningPromoOfferUpdateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningPromoOfferUpdateDTO(offerId string, warnings []PromoOfferUpdateWarningDTO) *WarningPromoOfferUpdateDTO {
	this := WarningPromoOfferUpdateDTO{}
	this.OfferId = offerId
	this.Warnings = warnings
	return &this
}

// NewWarningPromoOfferUpdateDTOWithDefaults instantiates a new WarningPromoOfferUpdateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningPromoOfferUpdateDTOWithDefaults() *WarningPromoOfferUpdateDTO {
	this := WarningPromoOfferUpdateDTO{}
	return &this
}

// GetOfferId returns the OfferId field value
func (o *WarningPromoOfferUpdateDTO) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *WarningPromoOfferUpdateDTO) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *WarningPromoOfferUpdateDTO) SetOfferId(v string) {
	o.OfferId = v
}

// GetWarnings returns the Warnings field value
func (o *WarningPromoOfferUpdateDTO) GetWarnings() []PromoOfferUpdateWarningDTO {
	if o == nil {
		var ret []PromoOfferUpdateWarningDTO
		return ret
	}

	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *WarningPromoOfferUpdateDTO) GetWarningsOk() ([]PromoOfferUpdateWarningDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warnings, true
}

// SetWarnings sets field value
func (o *WarningPromoOfferUpdateDTO) SetWarnings(v []PromoOfferUpdateWarningDTO) {
	o.Warnings = v
}

func (o WarningPromoOfferUpdateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningPromoOfferUpdateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerId"] = o.OfferId
	toSerialize["warnings"] = o.Warnings
	return toSerialize, nil
}

func (o *WarningPromoOfferUpdateDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerId",
		"warnings",
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

	varWarningPromoOfferUpdateDTO := _WarningPromoOfferUpdateDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWarningPromoOfferUpdateDTO)

	if err != nil {
		return err
	}

	*o = WarningPromoOfferUpdateDTO(varWarningPromoOfferUpdateDTO)

	return err
}

type NullableWarningPromoOfferUpdateDTO struct {
	value *WarningPromoOfferUpdateDTO
	isSet bool
}

func (v NullableWarningPromoOfferUpdateDTO) Get() *WarningPromoOfferUpdateDTO {
	return v.value
}

func (v *NullableWarningPromoOfferUpdateDTO) Set(val *WarningPromoOfferUpdateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningPromoOfferUpdateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningPromoOfferUpdateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningPromoOfferUpdateDTO(val *WarningPromoOfferUpdateDTO) *NullableWarningPromoOfferUpdateDTO {
	return &NullableWarningPromoOfferUpdateDTO{value: val, isSet: true}
}

func (v NullableWarningPromoOfferUpdateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningPromoOfferUpdateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


