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

// checks if the ChangeOutletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeOutletRequest{}

// ChangeOutletRequest struct for ChangeOutletRequest
type ChangeOutletRequest struct {
	// Название точки продаж. 
	Name string `json:"name"`
	Type OutletType `json:"type"`
	// Координаты точки продаж.  Формат: долгота, широта. Разделители: запятая и / или пробел. Например, `20.4522144, 54.7104264`.  Если параметр не передан, координаты будут определены по значениям параметров, вложенных в `address`. 
	Coords *string `json:"coords,omitempty"`
	// Признак основной точки продаж.  Возможные значения:  * `false` — неосновная точка продаж. * `true` — основная точка продаж. 
	IsMain *bool `json:"isMain,omitempty"`
	// Идентификатор точки продаж, присвоенный магазином.
	ShopOutletCode *string `json:"shopOutletCode,omitempty"`
	Visibility *OutletVisibilityType `json:"visibility,omitempty"`
	Address OutletAddressDTO `json:"address"`
	// Номера телефонов точки продаж. Передавайте в формате: `+7 (999) 999-99-99`. 
	Phones []string `json:"phones"`
	WorkingSchedule OutletWorkingScheduleDTO `json:"workingSchedule"`
	// Информация об условиях доставки для данной точки продаж.  Обязательный параметр, если параметр `type=DEPOT` или `type=MIXED`. 
	DeliveryRules []OutletDeliveryRuleDTO `json:"deliveryRules,omitempty"`
	// Срок хранения заказа в собственном пункте выдачи заказов. Считается в днях.
	StoragePeriod *int64 `json:"storagePeriod,omitempty"`
}

type _ChangeOutletRequest ChangeOutletRequest

// NewChangeOutletRequest instantiates a new ChangeOutletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeOutletRequest(name string, type_ OutletType, address OutletAddressDTO, phones []string, workingSchedule OutletWorkingScheduleDTO) *ChangeOutletRequest {
	this := ChangeOutletRequest{}
	this.Name = name
	this.Type = type_
	this.Address = address
	this.Phones = phones
	this.WorkingSchedule = workingSchedule
	return &this
}

// NewChangeOutletRequestWithDefaults instantiates a new ChangeOutletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeOutletRequestWithDefaults() *ChangeOutletRequest {
	this := ChangeOutletRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ChangeOutletRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChangeOutletRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ChangeOutletRequest) GetType() OutletType {
	if o == nil {
		var ret OutletType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetTypeOk() (*OutletType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChangeOutletRequest) SetType(v OutletType) {
	o.Type = v
}

// GetCoords returns the Coords field value if set, zero value otherwise.
func (o *ChangeOutletRequest) GetCoords() string {
	if o == nil || IsNil(o.Coords) {
		var ret string
		return ret
	}
	return *o.Coords
}

// GetCoordsOk returns a tuple with the Coords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetCoordsOk() (*string, bool) {
	if o == nil || IsNil(o.Coords) {
		return nil, false
	}
	return o.Coords, true
}

// HasCoords returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasCoords() bool {
	if o != nil && !IsNil(o.Coords) {
		return true
	}

	return false
}

// SetCoords gets a reference to the given string and assigns it to the Coords field.
func (o *ChangeOutletRequest) SetCoords(v string) {
	o.Coords = &v
}

// GetIsMain returns the IsMain field value if set, zero value otherwise.
func (o *ChangeOutletRequest) GetIsMain() bool {
	if o == nil || IsNil(o.IsMain) {
		var ret bool
		return ret
	}
	return *o.IsMain
}

// GetIsMainOk returns a tuple with the IsMain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetIsMainOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMain) {
		return nil, false
	}
	return o.IsMain, true
}

// HasIsMain returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasIsMain() bool {
	if o != nil && !IsNil(o.IsMain) {
		return true
	}

	return false
}

// SetIsMain gets a reference to the given bool and assigns it to the IsMain field.
func (o *ChangeOutletRequest) SetIsMain(v bool) {
	o.IsMain = &v
}

// GetShopOutletCode returns the ShopOutletCode field value if set, zero value otherwise.
func (o *ChangeOutletRequest) GetShopOutletCode() string {
	if o == nil || IsNil(o.ShopOutletCode) {
		var ret string
		return ret
	}
	return *o.ShopOutletCode
}

// GetShopOutletCodeOk returns a tuple with the ShopOutletCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetShopOutletCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShopOutletCode) {
		return nil, false
	}
	return o.ShopOutletCode, true
}

// HasShopOutletCode returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasShopOutletCode() bool {
	if o != nil && !IsNil(o.ShopOutletCode) {
		return true
	}

	return false
}

// SetShopOutletCode gets a reference to the given string and assigns it to the ShopOutletCode field.
func (o *ChangeOutletRequest) SetShopOutletCode(v string) {
	o.ShopOutletCode = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ChangeOutletRequest) GetVisibility() OutletVisibilityType {
	if o == nil || IsNil(o.Visibility) {
		var ret OutletVisibilityType
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetVisibilityOk() (*OutletVisibilityType, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given OutletVisibilityType and assigns it to the Visibility field.
func (o *ChangeOutletRequest) SetVisibility(v OutletVisibilityType) {
	o.Visibility = &v
}

// GetAddress returns the Address field value
func (o *ChangeOutletRequest) GetAddress() OutletAddressDTO {
	if o == nil {
		var ret OutletAddressDTO
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetAddressOk() (*OutletAddressDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ChangeOutletRequest) SetAddress(v OutletAddressDTO) {
	o.Address = v
}

// GetPhones returns the Phones field value
func (o *ChangeOutletRequest) GetPhones() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Phones
}

// GetPhonesOk returns a tuple with the Phones field value
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetPhonesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phones, true
}

// SetPhones sets field value
func (o *ChangeOutletRequest) SetPhones(v []string) {
	o.Phones = v
}

// GetWorkingSchedule returns the WorkingSchedule field value
func (o *ChangeOutletRequest) GetWorkingSchedule() OutletWorkingScheduleDTO {
	if o == nil {
		var ret OutletWorkingScheduleDTO
		return ret
	}

	return o.WorkingSchedule
}

// GetWorkingScheduleOk returns a tuple with the WorkingSchedule field value
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetWorkingScheduleOk() (*OutletWorkingScheduleDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkingSchedule, true
}

// SetWorkingSchedule sets field value
func (o *ChangeOutletRequest) SetWorkingSchedule(v OutletWorkingScheduleDTO) {
	o.WorkingSchedule = v
}

// GetDeliveryRules returns the DeliveryRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeOutletRequest) GetDeliveryRules() []OutletDeliveryRuleDTO {
	if o == nil {
		var ret []OutletDeliveryRuleDTO
		return ret
	}
	return o.DeliveryRules
}

// GetDeliveryRulesOk returns a tuple with the DeliveryRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeOutletRequest) GetDeliveryRulesOk() ([]OutletDeliveryRuleDTO, bool) {
	if o == nil || IsNil(o.DeliveryRules) {
		return nil, false
	}
	return o.DeliveryRules, true
}

// HasDeliveryRules returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasDeliveryRules() bool {
	if o != nil && !IsNil(o.DeliveryRules) {
		return true
	}

	return false
}

// SetDeliveryRules gets a reference to the given []OutletDeliveryRuleDTO and assigns it to the DeliveryRules field.
func (o *ChangeOutletRequest) SetDeliveryRules(v []OutletDeliveryRuleDTO) {
	o.DeliveryRules = v
}

// GetStoragePeriod returns the StoragePeriod field value if set, zero value otherwise.
func (o *ChangeOutletRequest) GetStoragePeriod() int64 {
	if o == nil || IsNil(o.StoragePeriod) {
		var ret int64
		return ret
	}
	return *o.StoragePeriod
}

// GetStoragePeriodOk returns a tuple with the StoragePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeOutletRequest) GetStoragePeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.StoragePeriod) {
		return nil, false
	}
	return o.StoragePeriod, true
}

// HasStoragePeriod returns a boolean if a field has been set.
func (o *ChangeOutletRequest) HasStoragePeriod() bool {
	if o != nil && !IsNil(o.StoragePeriod) {
		return true
	}

	return false
}

// SetStoragePeriod gets a reference to the given int64 and assigns it to the StoragePeriod field.
func (o *ChangeOutletRequest) SetStoragePeriod(v int64) {
	o.StoragePeriod = &v
}

func (o ChangeOutletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeOutletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Coords) {
		toSerialize["coords"] = o.Coords
	}
	if !IsNil(o.IsMain) {
		toSerialize["isMain"] = o.IsMain
	}
	if !IsNil(o.ShopOutletCode) {
		toSerialize["shopOutletCode"] = o.ShopOutletCode
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	toSerialize["address"] = o.Address
	toSerialize["phones"] = o.Phones
	toSerialize["workingSchedule"] = o.WorkingSchedule
	if o.DeliveryRules != nil {
		toSerialize["deliveryRules"] = o.DeliveryRules
	}
	if !IsNil(o.StoragePeriod) {
		toSerialize["storagePeriod"] = o.StoragePeriod
	}
	return toSerialize, nil
}

func (o *ChangeOutletRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"address",
		"phones",
		"workingSchedule",
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

	varChangeOutletRequest := _ChangeOutletRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChangeOutletRequest)

	if err != nil {
		return err
	}

	*o = ChangeOutletRequest(varChangeOutletRequest)

	return err
}

type NullableChangeOutletRequest struct {
	value *ChangeOutletRequest
	isSet bool
}

func (v NullableChangeOutletRequest) Get() *ChangeOutletRequest {
	return v.value
}

func (v *NullableChangeOutletRequest) Set(val *ChangeOutletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeOutletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeOutletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeOutletRequest(val *ChangeOutletRequest) *NullableChangeOutletRequest {
	return &NullableChangeOutletRequest{value: val, isSet: true}
}

func (v NullableChangeOutletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeOutletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


