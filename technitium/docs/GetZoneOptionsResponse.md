# GetZoneOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**GetZoneOptionsResponseResponse**](GetZoneOptionsResponseResponse.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewGetZoneOptionsResponse

`func NewGetZoneOptionsResponse() *GetZoneOptionsResponse`

NewGetZoneOptionsResponse instantiates a new GetZoneOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneOptionsResponseWithDefaults

`func NewGetZoneOptionsResponseWithDefaults() *GetZoneOptionsResponse`

NewGetZoneOptionsResponseWithDefaults instantiates a new GetZoneOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *GetZoneOptionsResponse) GetResponse() GetZoneOptionsResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GetZoneOptionsResponse) GetResponseOk() (*GetZoneOptionsResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GetZoneOptionsResponse) SetResponse(v GetZoneOptionsResponseResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GetZoneOptionsResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *GetZoneOptionsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetZoneOptionsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetZoneOptionsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetZoneOptionsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetZoneOptionsResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetZoneOptionsResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetZoneOptionsResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetZoneOptionsResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


