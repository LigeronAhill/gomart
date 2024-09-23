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

// checks if the ParameterValueDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterValueDTO{}

// ParameterValueDTO Значение характеристики.  Вы можете указывать несколько значений одной характеристики при условии, что:  * Тип характеристики — `ENUM`. * В ответе на запрос [POST category/{categoryId}/parameters](../../reference/content/getCategoryContentParameters.md) у данной характеристики поле `multivalue` имеет значение `true`.  Для этого в `parameterValues` передавайте каждое значение отдельно — несколько объектов с параметрами `parameterId`, `valueId` и `value`. Параметр `parameterId` должен быть одинаковым. 
type ParameterValueDTO struct {
	// Идентификатор характеристики.
	ParameterId int64 `json:"parameterId"`
	// Идентификатор единицы измерения. Если вы не передали параметр `unitId`, используется единица измерения по умолчанию.
	UnitId *int64 `json:"unitId,omitempty"`
	// Идентификатор значения.  Обязательно указывайте идентификатор, если передаете значение из перечня допустимых значений, полученного от Маркета.  Только для характеристик типа `ENUM`. 
	ValueId *int64 `json:"valueId,omitempty"`
	// Значение.
	Value *string `json:"value,omitempty"`
}

type _ParameterValueDTO ParameterValueDTO

// NewParameterValueDTO instantiates a new ParameterValueDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterValueDTO(parameterId int64) *ParameterValueDTO {
	this := ParameterValueDTO{}
	this.ParameterId = parameterId
	return &this
}

// NewParameterValueDTOWithDefaults instantiates a new ParameterValueDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterValueDTOWithDefaults() *ParameterValueDTO {
	this := ParameterValueDTO{}
	return &this
}

// GetParameterId returns the ParameterId field value
func (o *ParameterValueDTO) GetParameterId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value
// and a boolean to check if the value has been set.
func (o *ParameterValueDTO) GetParameterIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterId, true
}

// SetParameterId sets field value
func (o *ParameterValueDTO) SetParameterId(v int64) {
	o.ParameterId = v
}

// GetUnitId returns the UnitId field value if set, zero value otherwise.
func (o *ParameterValueDTO) GetUnitId() int64 {
	if o == nil || IsNil(o.UnitId) {
		var ret int64
		return ret
	}
	return *o.UnitId
}

// GetUnitIdOk returns a tuple with the UnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterValueDTO) GetUnitIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitId) {
		return nil, false
	}
	return o.UnitId, true
}

// HasUnitId returns a boolean if a field has been set.
func (o *ParameterValueDTO) HasUnitId() bool {
	if o != nil && !IsNil(o.UnitId) {
		return true
	}

	return false
}

// SetUnitId gets a reference to the given int64 and assigns it to the UnitId field.
func (o *ParameterValueDTO) SetUnitId(v int64) {
	o.UnitId = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *ParameterValueDTO) GetValueId() int64 {
	if o == nil || IsNil(o.ValueId) {
		var ret int64
		return ret
	}
	return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterValueDTO) GetValueIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ValueId) {
		return nil, false
	}
	return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *ParameterValueDTO) HasValueId() bool {
	if o != nil && !IsNil(o.ValueId) {
		return true
	}

	return false
}

// SetValueId gets a reference to the given int64 and assigns it to the ValueId field.
func (o *ParameterValueDTO) SetValueId(v int64) {
	o.ValueId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ParameterValueDTO) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterValueDTO) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ParameterValueDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ParameterValueDTO) SetValue(v string) {
	o.Value = &v
}

func (o ParameterValueDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterValueDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parameterId"] = o.ParameterId
	if !IsNil(o.UnitId) {
		toSerialize["unitId"] = o.UnitId
	}
	if !IsNil(o.ValueId) {
		toSerialize["valueId"] = o.ValueId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *ParameterValueDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parameterId",
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

	varParameterValueDTO := _ParameterValueDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParameterValueDTO)

	if err != nil {
		return err
	}

	*o = ParameterValueDTO(varParameterValueDTO)

	return err
}

type NullableParameterValueDTO struct {
	value *ParameterValueDTO
	isSet bool
}

func (v NullableParameterValueDTO) Get() *ParameterValueDTO {
	return v.value
}

func (v *NullableParameterValueDTO) Set(val *ParameterValueDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterValueDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterValueDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterValueDTO(val *ParameterValueDTO) *NullableParameterValueDTO {
	return &NullableParameterValueDTO{value: val, isSet: true}
}

func (v NullableParameterValueDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterValueDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


