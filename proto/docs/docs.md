# Protocol Documentation
<a name="top"/>

## Table of Contents

- [cache.proto](#cache.proto)
    - [FlushResponse](#cache.FlushResponse)
    - [GetRequest](#cache.GetRequest)
    - [GetResponse](#cache.GetResponse)
    - [SetRequest](#cache.SetRequest)
    - [SetResponse](#cache.SetResponse)
    - [UnsetRequest](#cache.UnsetRequest)
    - [UnsetResponse](#cache.UnsetResponse)
  
  
  
    - [Cache](#cache.Cache)
  

- [Scalar Value Types](#scalar-value-types)



<a name="cache.proto"/>
<p align="right"><a href="#top">Top</a></p>

## cache.proto



<a name="cache.FlushResponse"/>

### FlushResponse
Defines the return(Response) to the Flush method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Cache was flushed or not |






<a name="cache.GetRequest"/>

### GetRequest
Defines the Request to the Get method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Unique value identifier. |






<a name="cache.GetResponse"/>

### GetResponse
Defines the return(Response) to the Get method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  | Value of a given key. |






<a name="cache.SetRequest"/>

### SetRequest
Defines the Request to the Set method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Unique value identifier. |
| value | [bytes](#bytes) |  | Value to be set. |






<a name="cache.SetResponse"/>

### SetResponse
Defines the return(Response) to the Set method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Says whether value has been set or not. |






<a name="cache.UnsetRequest"/>

### UnsetRequest
Defines the Request to the Unset method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Key to remove from cache. |






<a name="cache.UnsetResponse"/>

### UnsetResponse
Defines the return(Response) value to the Unset method of Cache.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | Key was removed or not. |





 

 

 


<a name="cache.Cache"/>

### Cache
Service for handling caching request

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#cache.GetRequest) | [GetResponse](#cache.GetRequest) | Used to retrieve a key from the cache service. |
| Set | [SetRequest](#cache.SetRequest) | [SetResponse](#cache.SetRequest) | Used to set a new key from the cache service. |
| Unset | [UnsetRequest](#cache.UnsetRequest) | [UnsetResponse](#cache.UnsetRequest) | Used to remove a key from the cache service. |
| Flush | [.google.protobuf.Empty](#google.protobuf.Empty) | [FlushResponse](#google.protobuf.Empty) | Used tot delete all the values from the cache service. |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

