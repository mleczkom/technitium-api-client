# CreateZoneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**CreateZoneResponseResponse**](CreateZoneResponseResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateZoneResponse

`func NewCreateZoneResponse() *CreateZoneResponse`

NewCreateZoneResponse instantiates a new CreateZoneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateZoneResponseWithDefaults

`func NewCreateZoneResponseWithDefaults() *CreateZoneResponse`

NewCreateZoneResponseWithDefaults instantiates a new CreateZoneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *CreateZoneResponse) GetResponse() CreateZoneResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CreateZoneResponse) GetResponseOk() (*CreateZoneResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CreateZoneResponse) SetResponse(v CreateZoneResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *CreateZoneResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *CreateZoneResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateZoneResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateZoneResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateZoneResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CreateZoneResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CreateZoneResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CreateZoneResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CreateZoneResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


