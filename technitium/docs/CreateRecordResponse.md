# CreateRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**CreateRecordResponseResponse**](CreateRecordResponseResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRecordResponse

`func NewCreateRecordResponse() *CreateRecordResponse`

NewCreateRecordResponse instantiates a new CreateRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordResponseWithDefaults

`func NewCreateRecordResponseWithDefaults() *CreateRecordResponse`

NewCreateRecordResponseWithDefaults instantiates a new CreateRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *CreateRecordResponse) GetResponse() CreateRecordResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CreateRecordResponse) GetResponseOk() (*CreateRecordResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CreateRecordResponse) SetResponse(v CreateRecordResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *CreateRecordResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *CreateRecordResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRecordResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRecordResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRecordResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CreateRecordResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateRecordResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateRecordResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateRecordResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


