# GetGoodsFeedbackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ApiResponseStatusType**](ApiResponseStatusType.md) |  | [optional] 
**Result** | Pointer to [**GoodsFeedbackListDTO**](GoodsFeedbackListDTO.md) |  | [optional] 

## Methods

### NewGetGoodsFeedbackResponse

`func NewGetGoodsFeedbackResponse() *GetGoodsFeedbackResponse`

NewGetGoodsFeedbackResponse instantiates a new GetGoodsFeedbackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoodsFeedbackResponseWithDefaults

`func NewGetGoodsFeedbackResponseWithDefaults() *GetGoodsFeedbackResponse`

NewGetGoodsFeedbackResponseWithDefaults instantiates a new GetGoodsFeedbackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetGoodsFeedbackResponse) GetStatus() ApiResponseStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGoodsFeedbackResponse) GetStatusOk() (*ApiResponseStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGoodsFeedbackResponse) SetStatus(v ApiResponseStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGoodsFeedbackResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetGoodsFeedbackResponse) GetResult() GoodsFeedbackListDTO`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGoodsFeedbackResponse) GetResultOk() (*GoodsFeedbackListDTO, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGoodsFeedbackResponse) SetResult(v GoodsFeedbackListDTO)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGoodsFeedbackResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


