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

// checks if the FlippingPagerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlippingPagerDTO{}

// FlippingPagerDTO Модель для пагинации.
type FlippingPagerDTO struct {
	// Сколько всего найдено элементов.
	Total *int32 `json:"total,omitempty"`
	// Начальный номер найденного элемента на странице.
	From *int32 `json:"from,omitempty"`
	// Конечный номер найденного элемента на странице.
	To *int32 `json:"to,omitempty"`
	// Текущая страница.
	CurrentPage *int32 `json:"currentPage,omitempty"`
	// Общее количество страниц.
	PagesCount *int32 `json:"pagesCount,omitempty"`
	// Размер страницы.
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewFlippingPagerDTO instantiates a new FlippingPagerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlippingPagerDTO() *FlippingPagerDTO {
	this := FlippingPagerDTO{}
	return &this
}

// NewFlippingPagerDTOWithDefaults instantiates a new FlippingPagerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlippingPagerDTOWithDefaults() *FlippingPagerDTO {
	this := FlippingPagerDTO{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *FlippingPagerDTO) SetTotal(v int32) {
	o.Total = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetFrom() int32 {
	if o == nil || IsNil(o.From) {
		var ret int32
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetFromOk() (*int32, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given int32 and assigns it to the From field.
func (o *FlippingPagerDTO) SetFrom(v int32) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetTo() int32 {
	if o == nil || IsNil(o.To) {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetToOk() (*int32, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *FlippingPagerDTO) SetTo(v int32) {
	o.To = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *FlippingPagerDTO) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetPagesCount() int32 {
	if o == nil || IsNil(o.PagesCount) {
		var ret int32
		return ret
	}
	return *o.PagesCount
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetPagesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PagesCount) {
		return nil, false
	}
	return o.PagesCount, true
}

// HasPagesCount returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasPagesCount() bool {
	if o != nil && !IsNil(o.PagesCount) {
		return true
	}

	return false
}

// SetPagesCount gets a reference to the given int32 and assigns it to the PagesCount field.
func (o *FlippingPagerDTO) SetPagesCount(v int32) {
	o.PagesCount = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *FlippingPagerDTO) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlippingPagerDTO) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *FlippingPagerDTO) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *FlippingPagerDTO) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o FlippingPagerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlippingPagerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !IsNil(o.PagesCount) {
		toSerialize["pagesCount"] = o.PagesCount
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableFlippingPagerDTO struct {
	value *FlippingPagerDTO
	isSet bool
}

func (v NullableFlippingPagerDTO) Get() *FlippingPagerDTO {
	return v.value
}

func (v *NullableFlippingPagerDTO) Set(val *FlippingPagerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFlippingPagerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFlippingPagerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlippingPagerDTO(val *FlippingPagerDTO) *NullableFlippingPagerDTO {
	return &NullableFlippingPagerDTO{value: val, isSet: true}
}

func (v NullableFlippingPagerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlippingPagerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


