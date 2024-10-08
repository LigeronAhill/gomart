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

// checks if the GetRegionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRegionsResponse{}

// GetRegionsResponse struct for GetRegionsResponse
type GetRegionsResponse struct {
	// Регион доставки.
	Regions []RegionDTO `json:"regions"`
	Paging *ForwardScrollingPagerDTO `json:"paging,omitempty"`
}

type _GetRegionsResponse GetRegionsResponse

// NewGetRegionsResponse instantiates a new GetRegionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionsResponse(regions []RegionDTO) *GetRegionsResponse {
	this := GetRegionsResponse{}
	this.Regions = regions
	return &this
}

// NewGetRegionsResponseWithDefaults instantiates a new GetRegionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionsResponseWithDefaults() *GetRegionsResponse {
	this := GetRegionsResponse{}
	return &this
}

// GetRegions returns the Regions field value
func (o *GetRegionsResponse) GetRegions() []RegionDTO {
	if o == nil {
		var ret []RegionDTO
		return ret
	}

	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value
// and a boolean to check if the value has been set.
func (o *GetRegionsResponse) GetRegionsOk() ([]RegionDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regions, true
}

// SetRegions sets field value
func (o *GetRegionsResponse) SetRegions(v []RegionDTO) {
	o.Regions = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *GetRegionsResponse) GetPaging() ForwardScrollingPagerDTO {
	if o == nil || IsNil(o.Paging) {
		var ret ForwardScrollingPagerDTO
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionsResponse) GetPagingOk() (*ForwardScrollingPagerDTO, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *GetRegionsResponse) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardScrollingPagerDTO and assigns it to the Paging field.
func (o *GetRegionsResponse) SetPaging(v ForwardScrollingPagerDTO) {
	o.Paging = &v
}

func (o GetRegionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regions"] = o.Regions
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

func (o *GetRegionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"regions",
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

	varGetRegionsResponse := _GetRegionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRegionsResponse)

	if err != nil {
		return err
	}

	*o = GetRegionsResponse(varGetRegionsResponse)

	return err
}

type NullableGetRegionsResponse struct {
	value *GetRegionsResponse
	isSet bool
}

func (v NullableGetRegionsResponse) Get() *GetRegionsResponse {
	return v.value
}

func (v *NullableGetRegionsResponse) Set(val *GetRegionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionsResponse(val *GetRegionsResponse) *NullableGetRegionsResponse {
	return &NullableGetRegionsResponse{value: val, isSet: true}
}

func (v NullableGetRegionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


