# CreateRecordResponseResponseAddedRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**RData** | Pointer to [**CreateRecordResponseResponseAddedRecordRData**](CreateRecordResponseResponseAddedRecordRData.md) |  | [optional] 
**DnssecStatus** | Pointer to **string** |  | [optional] 
**LastUsedOn** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRecordResponseResponseAddedRecord

`func NewCreateRecordResponseResponseAddedRecord() *CreateRecordResponseResponseAddedRecord`

NewCreateRecordResponseResponseAddedRecord instantiates a new CreateRecordResponseResponseAddedRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordResponseResponseAddedRecordWithDefaults

`func NewCreateRecordResponseResponseAddedRecordWithDefaults() *CreateRecordResponseResponseAddedRecord`

NewCreateRecordResponseResponseAddedRecordWithDefaults instantiates a new CreateRecordResponseResponseAddedRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *CreateRecordResponseResponseAddedRecord) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CreateRecordResponseResponseAddedRecord) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CreateRecordResponseResponseAddedRecord) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CreateRecordResponseResponseAddedRecord) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *CreateRecordResponseResponseAddedRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRecordResponseResponseAddedRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRecordResponseResponseAddedRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRecordResponseResponseAddedRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateRecordResponseResponseAddedRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRecordResponseResponseAddedRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRecordResponseResponseAddedRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateRecordResponseResponseAddedRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTtl

`func (o *CreateRecordResponseResponseAddedRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateRecordResponseResponseAddedRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateRecordResponseResponseAddedRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateRecordResponseResponseAddedRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetRData

`func (o *CreateRecordResponseResponseAddedRecord) GetRData() CreateRecordResponseResponseAddedRecordRData`

GetRData returns the RData field if non-nil, zero value otherwise.

### GetRDataOk

`func (o *CreateRecordResponseResponseAddedRecord) GetRDataOk() (*CreateRecordResponseResponseAddedRecordRData, bool)`

GetRDataOk returns a tuple with the RData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRData

`func (o *CreateRecordResponseResponseAddedRecord) SetRData(v CreateRecordResponseResponseAddedRecordRData)`

SetRData sets RData field to given value.

### HasRData

`func (o *CreateRecordResponseResponseAddedRecord) HasRData() bool`

HasRData returns a boolean if a field has been set.

### GetDnssecStatus

`func (o *CreateRecordResponseResponseAddedRecord) GetDnssecStatus() string`

GetDnssecStatus returns the DnssecStatus field if non-nil, zero value otherwise.

### GetDnssecStatusOk

`func (o *CreateRecordResponseResponseAddedRecord) GetDnssecStatusOk() (*string, bool)`

GetDnssecStatusOk returns a tuple with the DnssecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecStatus

`func (o *CreateRecordResponseResponseAddedRecord) SetDnssecStatus(v string)`

SetDnssecStatus sets DnssecStatus field to given value.

### HasDnssecStatus

`func (o *CreateRecordResponseResponseAddedRecord) HasDnssecStatus() bool`

HasDnssecStatus returns a boolean if a field has been set.

### GetLastUsedOn

`func (o *CreateRecordResponseResponseAddedRecord) GetLastUsedOn() string`

GetLastUsedOn returns the LastUsedOn field if non-nil, zero value otherwise.

### GetLastUsedOnOk

`func (o *CreateRecordResponseResponseAddedRecord) GetLastUsedOnOk() (*string, bool)`

GetLastUsedOnOk returns a tuple with the LastUsedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedOn

`func (o *CreateRecordResponseResponseAddedRecord) SetLastUsedOn(v string)`

SetLastUsedOn sets LastUsedOn field to given value.

### HasLastUsedOn

`func (o *CreateRecordResponseResponseAddedRecord) HasLastUsedOn() bool`

HasLastUsedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


