# CommonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewCommonResponse

`func NewCommonResponse() *CommonResponse`

NewCommonResponse instantiates a new CommonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonResponseWithDefaults

`func NewCommonResponseWithDefaults() *CommonResponse`

NewCommonResponseWithDefaults instantiates a new CommonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommonResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommonResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommonResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommonResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CommonResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CommonResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CommonResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CommonResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


