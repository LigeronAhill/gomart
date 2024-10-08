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

// checks if the GenerateUnitedNettingReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateUnitedNettingReportRequest{}

// GenerateUnitedNettingReportRequest Данные, необходимые для генерации отчета: идентификатор магазина, период, за который нужен отчет, а также фильтры. 
type GenerateUnitedNettingReportRequest struct {
	// Идентификатор бизнеса.
	BusinessId int64 `json:"businessId"`
	// {% note warning \"\" %}  Этот параметр устарел. Не используйте его.  {% endnote %}  Начало периода, включительно. 
	DateTimeFrom *time.Time `json:"dateTimeFrom,omitempty"`
	// {% note warning \"\" %}  Этот параметр устарел. Не используйте его.  {% endnote %}  Конец периода, включительно. Максимальный период — 1 год. 
	DateTimeTo *time.Time `json:"dateTimeTo,omitempty"`
	// Начало периода, включительно.
	DateFrom *string `json:"dateFrom,omitempty"`
	// Конец периода, включительно. Максимальный период — 1 год.
	DateTo *string `json:"dateTo,omitempty"`
	// Номер платежного поручения.
	BankOrderId *int64 `json:"bankOrderId,omitempty"`
	// Дата платежного поручения.
	BankOrderDateTime *time.Time `json:"bankOrderDateTime,omitempty"`
	// Список моделей, которые нужны в отчете. 
	PlacementPrograms []PlacementType `json:"placementPrograms,omitempty"`
	// Список ИНН, которые нужны в отчете.
	Inns []string `json:"inns,omitempty"`
	// Список магазинов, которые нужны в отчете.
	CampaignIds []int64 `json:"campaignIds,omitempty"`
}

type _GenerateUnitedNettingReportRequest GenerateUnitedNettingReportRequest

// NewGenerateUnitedNettingReportRequest instantiates a new GenerateUnitedNettingReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateUnitedNettingReportRequest(businessId int64) *GenerateUnitedNettingReportRequest {
	this := GenerateUnitedNettingReportRequest{}
	this.BusinessId = businessId
	return &this
}

// NewGenerateUnitedNettingReportRequestWithDefaults instantiates a new GenerateUnitedNettingReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateUnitedNettingReportRequestWithDefaults() *GenerateUnitedNettingReportRequest {
	this := GenerateUnitedNettingReportRequest{}
	return &this
}

// GetBusinessId returns the BusinessId field value
func (o *GenerateUnitedNettingReportRequest) GetBusinessId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetBusinessIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *GenerateUnitedNettingReportRequest) SetBusinessId(v int64) {
	o.BusinessId = v
}

// GetDateTimeFrom returns the DateTimeFrom field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetDateTimeFrom() time.Time {
	if o == nil || IsNil(o.DateTimeFrom) {
		var ret time.Time
		return ret
	}
	return *o.DateTimeFrom
}

// GetDateTimeFromOk returns a tuple with the DateTimeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetDateTimeFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTimeFrom) {
		return nil, false
	}
	return o.DateTimeFrom, true
}

// HasDateTimeFrom returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasDateTimeFrom() bool {
	if o != nil && !IsNil(o.DateTimeFrom) {
		return true
	}

	return false
}

// SetDateTimeFrom gets a reference to the given time.Time and assigns it to the DateTimeFrom field.
func (o *GenerateUnitedNettingReportRequest) SetDateTimeFrom(v time.Time) {
	o.DateTimeFrom = &v
}

// GetDateTimeTo returns the DateTimeTo field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetDateTimeTo() time.Time {
	if o == nil || IsNil(o.DateTimeTo) {
		var ret time.Time
		return ret
	}
	return *o.DateTimeTo
}

// GetDateTimeToOk returns a tuple with the DateTimeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetDateTimeToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateTimeTo) {
		return nil, false
	}
	return o.DateTimeTo, true
}

// HasDateTimeTo returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasDateTimeTo() bool {
	if o != nil && !IsNil(o.DateTimeTo) {
		return true
	}

	return false
}

// SetDateTimeTo gets a reference to the given time.Time and assigns it to the DateTimeTo field.
func (o *GenerateUnitedNettingReportRequest) SetDateTimeTo(v time.Time) {
	o.DateTimeTo = &v
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetDateFrom() string {
	if o == nil || IsNil(o.DateFrom) {
		var ret string
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetDateFromOk() (*string, bool) {
	if o == nil || IsNil(o.DateFrom) {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasDateFrom() bool {
	if o != nil && !IsNil(o.DateFrom) {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given string and assigns it to the DateFrom field.
func (o *GenerateUnitedNettingReportRequest) SetDateFrom(v string) {
	o.DateFrom = &v
}

// GetDateTo returns the DateTo field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetDateTo() string {
	if o == nil || IsNil(o.DateTo) {
		var ret string
		return ret
	}
	return *o.DateTo
}

// GetDateToOk returns a tuple with the DateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetDateToOk() (*string, bool) {
	if o == nil || IsNil(o.DateTo) {
		return nil, false
	}
	return o.DateTo, true
}

// HasDateTo returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasDateTo() bool {
	if o != nil && !IsNil(o.DateTo) {
		return true
	}

	return false
}

// SetDateTo gets a reference to the given string and assigns it to the DateTo field.
func (o *GenerateUnitedNettingReportRequest) SetDateTo(v string) {
	o.DateTo = &v
}

// GetBankOrderId returns the BankOrderId field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetBankOrderId() int64 {
	if o == nil || IsNil(o.BankOrderId) {
		var ret int64
		return ret
	}
	return *o.BankOrderId
}

// GetBankOrderIdOk returns a tuple with the BankOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetBankOrderIdOk() (*int64, bool) {
	if o == nil || IsNil(o.BankOrderId) {
		return nil, false
	}
	return o.BankOrderId, true
}

// HasBankOrderId returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasBankOrderId() bool {
	if o != nil && !IsNil(o.BankOrderId) {
		return true
	}

	return false
}

// SetBankOrderId gets a reference to the given int64 and assigns it to the BankOrderId field.
func (o *GenerateUnitedNettingReportRequest) SetBankOrderId(v int64) {
	o.BankOrderId = &v
}

// GetBankOrderDateTime returns the BankOrderDateTime field value if set, zero value otherwise.
func (o *GenerateUnitedNettingReportRequest) GetBankOrderDateTime() time.Time {
	if o == nil || IsNil(o.BankOrderDateTime) {
		var ret time.Time
		return ret
	}
	return *o.BankOrderDateTime
}

// GetBankOrderDateTimeOk returns a tuple with the BankOrderDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateUnitedNettingReportRequest) GetBankOrderDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BankOrderDateTime) {
		return nil, false
	}
	return o.BankOrderDateTime, true
}

// HasBankOrderDateTime returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasBankOrderDateTime() bool {
	if o != nil && !IsNil(o.BankOrderDateTime) {
		return true
	}

	return false
}

// SetBankOrderDateTime gets a reference to the given time.Time and assigns it to the BankOrderDateTime field.
func (o *GenerateUnitedNettingReportRequest) SetBankOrderDateTime(v time.Time) {
	o.BankOrderDateTime = &v
}

// GetPlacementPrograms returns the PlacementPrograms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateUnitedNettingReportRequest) GetPlacementPrograms() []PlacementType {
	if o == nil {
		var ret []PlacementType
		return ret
	}
	return o.PlacementPrograms
}

// GetPlacementProgramsOk returns a tuple with the PlacementPrograms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateUnitedNettingReportRequest) GetPlacementProgramsOk() ([]PlacementType, bool) {
	if o == nil || IsNil(o.PlacementPrograms) {
		return nil, false
	}
	return o.PlacementPrograms, true
}

// HasPlacementPrograms returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasPlacementPrograms() bool {
	if o != nil && !IsNil(o.PlacementPrograms) {
		return true
	}

	return false
}

// SetPlacementPrograms gets a reference to the given []PlacementType and assigns it to the PlacementPrograms field.
func (o *GenerateUnitedNettingReportRequest) SetPlacementPrograms(v []PlacementType) {
	o.PlacementPrograms = v
}

// GetInns returns the Inns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateUnitedNettingReportRequest) GetInns() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inns
}

// GetInnsOk returns a tuple with the Inns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateUnitedNettingReportRequest) GetInnsOk() ([]string, bool) {
	if o == nil || IsNil(o.Inns) {
		return nil, false
	}
	return o.Inns, true
}

// HasInns returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasInns() bool {
	if o != nil && !IsNil(o.Inns) {
		return true
	}

	return false
}

// SetInns gets a reference to the given []string and assigns it to the Inns field.
func (o *GenerateUnitedNettingReportRequest) SetInns(v []string) {
	o.Inns = v
}

// GetCampaignIds returns the CampaignIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenerateUnitedNettingReportRequest) GetCampaignIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenerateUnitedNettingReportRequest) GetCampaignIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.CampaignIds) {
		return nil, false
	}
	return o.CampaignIds, true
}

// HasCampaignIds returns a boolean if a field has been set.
func (o *GenerateUnitedNettingReportRequest) HasCampaignIds() bool {
	if o != nil && !IsNil(o.CampaignIds) {
		return true
	}

	return false
}

// SetCampaignIds gets a reference to the given []int64 and assigns it to the CampaignIds field.
func (o *GenerateUnitedNettingReportRequest) SetCampaignIds(v []int64) {
	o.CampaignIds = v
}

func (o GenerateUnitedNettingReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateUnitedNettingReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["businessId"] = o.BusinessId
	if !IsNil(o.DateTimeFrom) {
		toSerialize["dateTimeFrom"] = o.DateTimeFrom
	}
	if !IsNil(o.DateTimeTo) {
		toSerialize["dateTimeTo"] = o.DateTimeTo
	}
	if !IsNil(o.DateFrom) {
		toSerialize["dateFrom"] = o.DateFrom
	}
	if !IsNil(o.DateTo) {
		toSerialize["dateTo"] = o.DateTo
	}
	if !IsNil(o.BankOrderId) {
		toSerialize["bankOrderId"] = o.BankOrderId
	}
	if !IsNil(o.BankOrderDateTime) {
		toSerialize["bankOrderDateTime"] = o.BankOrderDateTime
	}
	if o.PlacementPrograms != nil {
		toSerialize["placementPrograms"] = o.PlacementPrograms
	}
	if o.Inns != nil {
		toSerialize["inns"] = o.Inns
	}
	if o.CampaignIds != nil {
		toSerialize["campaignIds"] = o.CampaignIds
	}
	return toSerialize, nil
}

func (o *GenerateUnitedNettingReportRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"businessId",
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

	varGenerateUnitedNettingReportRequest := _GenerateUnitedNettingReportRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateUnitedNettingReportRequest)

	if err != nil {
		return err
	}

	*o = GenerateUnitedNettingReportRequest(varGenerateUnitedNettingReportRequest)

	return err
}

type NullableGenerateUnitedNettingReportRequest struct {
	value *GenerateUnitedNettingReportRequest
	isSet bool
}

func (v NullableGenerateUnitedNettingReportRequest) Get() *GenerateUnitedNettingReportRequest {
	return v.value
}

func (v *NullableGenerateUnitedNettingReportRequest) Set(val *GenerateUnitedNettingReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateUnitedNettingReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateUnitedNettingReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateUnitedNettingReportRequest(val *GenerateUnitedNettingReportRequest) *NullableGenerateUnitedNettingReportRequest {
	return &NullableGenerateUnitedNettingReportRequest{value: val, isSet: true}
}

func (v NullableGenerateUnitedNettingReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateUnitedNettingReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


