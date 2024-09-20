# ListZonesResponseResponseZonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**DnssecStatus** | Pointer to **string** |  | [optional] 
**SoaSerial** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **string** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**SyncFailed** | Pointer to **bool** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewListZonesResponseResponseZonesInner

`func NewListZonesResponseResponseZonesInner() *ListZonesResponseResponseZonesInner`

NewListZonesResponseResponseZonesInner instantiates a new ListZonesResponseResponseZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListZonesResponseResponseZonesInnerWithDefaults

`func NewListZonesResponseResponseZonesInnerWithDefaults() *ListZonesResponseResponseZonesInner`

NewListZonesResponseResponseZonesInnerWithDefaults instantiates a new ListZonesResponseResponseZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListZonesResponseResponseZonesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListZonesResponseResponseZonesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListZonesResponseResponseZonesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListZonesResponseResponseZonesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListZonesResponseResponseZonesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListZonesResponseResponseZonesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListZonesResponseResponseZonesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListZonesResponseResponseZonesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInternal

`func (o *ListZonesResponseResponseZonesInner) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ListZonesResponseResponseZonesInner) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ListZonesResponseResponseZonesInner) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ListZonesResponseResponseZonesInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDnssecStatus

`func (o *ListZonesResponseResponseZonesInner) GetDnssecStatus() string`

GetDnssecStatus returns the DnssecStatus field if non-nil, zero value otherwise.

### GetDnssecStatusOk

`func (o *ListZonesResponseResponseZonesInner) GetDnssecStatusOk() (*string, bool)`

GetDnssecStatusOk returns a tuple with the DnssecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecStatus

`func (o *ListZonesResponseResponseZonesInner) SetDnssecStatus(v string)`

SetDnssecStatus sets DnssecStatus field to given value.

### HasDnssecStatus

`func (o *ListZonesResponseResponseZonesInner) HasDnssecStatus() bool`

HasDnssecStatus returns a boolean if a field has been set.

### GetSoaSerial

`func (o *ListZonesResponseResponseZonesInner) GetSoaSerial() int32`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *ListZonesResponseResponseZonesInner) GetSoaSerialOk() (*int32, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *ListZonesResponseResponseZonesInner) SetSoaSerial(v int32)`

SetSoaSerial sets SoaSerial field to given value.

### HasSoaSerial

`func (o *ListZonesResponseResponseZonesInner) HasSoaSerial() bool`

HasSoaSerial returns a boolean if a field has been set.

### GetExpiry

`func (o *ListZonesResponseResponseZonesInner) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ListZonesResponseResponseZonesInner) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ListZonesResponseResponseZonesInner) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ListZonesResponseResponseZonesInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetIsExpired

`func (o *ListZonesResponseResponseZonesInner) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *ListZonesResponseResponseZonesInner) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *ListZonesResponseResponseZonesInner) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *ListZonesResponseResponseZonesInner) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetSyncFailed

`func (o *ListZonesResponseResponseZonesInner) GetSyncFailed() bool`

GetSyncFailed returns the SyncFailed field if non-nil, zero value otherwise.

### GetSyncFailedOk

`func (o *ListZonesResponseResponseZonesInner) GetSyncFailedOk() (*bool, bool)`

GetSyncFailedOk returns a tuple with the SyncFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFailed

`func (o *ListZonesResponseResponseZonesInner) SetSyncFailed(v bool)`

SetSyncFailed sets SyncFailed field to given value.

### HasSyncFailed

`func (o *ListZonesResponseResponseZonesInner) HasSyncFailed() bool`

HasSyncFailed returns a boolean if a field has been set.

### GetLastModified

`func (o *ListZonesResponseResponseZonesInner) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ListZonesResponseResponseZonesInner) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ListZonesResponseResponseZonesInner) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ListZonesResponseResponseZonesInner) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetDisabled

`func (o *ListZonesResponseResponseZonesInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ListZonesResponseResponseZonesInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ListZonesResponseResponseZonesInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ListZonesResponseResponseZonesInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


