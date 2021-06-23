# \ValidateApi

All URIs are relative to *https://virtserver.swaggerhub.com/dga/Validate/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateGet**](ValidateApi.md#ValidateGet) | **Get** /validate | Validate


# **ValidateGet**
> TokenResult ValidateGet(ctx, consumerKey, consumerSecret, agentID, optional)
Validate

ตรวจสอบความถูกต้องของ Consumer Key และ Consumer Secret เพื่อขอ Token และเปิด Session ในการเรียกใช้ Service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| ชุดรหัสที่ สรอ. ออกให้ เพื่อความปลอดภัยในการเรียกใช้งาน API | 
  **consumerSecret** | **string**| ชุดรหัสที่ สรอ. ออกให้ เพื่อความปลอดภัยในการเรียกใช้งาน API | 
  **agentID** | **string**| - เลขบัตรประชาชน 13 หลัก ของผู้ที่เข้ามาดูข้อมูลในระบบ   - เป็นค่าที่สามารถบ่งบอกได้ว่า User ใดที่เข้ามาใช้งาน Service ใน Session นั้น ๆ - เป็นค่า Unique ID ของแต่ละ Session โดยทาง สรอ. กำหนดไว้ให้   | 
 **optional** | ***ValidateApiValidateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ValidateApiValidateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **optional.String**|  | [default to application/x-www-form-urlencoded; charset&#x3D;utf-8]

### Return type

[**TokenResult**](TokenResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

