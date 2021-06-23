# \DopaPersonalVerifyVerificationLaserCodeApi

All URIs are relative to *https://ws.ega.or.th/ws/dopa/verification*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonalGet**](DopaPersonalVerifyVerificationLaserCodeApi.md#PersonalGet) | **Get** /personal | การเรียกตรวจสอบข้อมูลบุคคลว่าตรงกับที่มีอยู่ในฐานข้อมูลทะเบียนราษฎร์หรือไม่


# **PersonalGet**
> Verify PersonalGet(ctx, consumerKey, token, citizenID, firstName, lastName, bEBirthDate, laserCode, optional)
การเรียกตรวจสอบข้อมูลบุคคลว่าตรงกับที่มีอยู่ในฐานข้อมูลทะเบียนราษฎร์หรือไม่

การเรียกตรวจสอบข้อมูลบุคคลของบุคคลนั้น ๆ ว่าตรงกับที่มีอยู่ใน ฐานข้อมูลทะเบียนราษฎร์ กรมการปกครอง หรือไม่ โดยใช้เลขหลังบัตรประชาชน (Laser Code)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| ชุดรหัสที่ สรอ. ออกให้ เพื่อความปลอดภัยในการเรียกใช้งาน API | 
  **token** | **string**| TOKEN ที่ได้รับกลับมาจาก สรอ. ในการ Validate API (https://app.swaggerhub.com/apis/dga/Validate/1.0) | 
  **citizenID** | **string**| หมายเลขบัตรประจำตัวประชาชน 13 หลัก | 
  **firstName** | **string**| ชื่อจริง (ภาษาไทย) | 
  **lastName** | **string**| นามสกุล (ภาษาไทย) | 
  **bEBirthDate** | **string**| วัน/เดือน/ปีเกิด (พ.ศ.)  โดยมีรูปแบบเป็น YYYYMMDD | 
  **laserCode** | **string**| หมายเลขหลังบัตรประจำตัวประชาชน | 
 **optional** | ***DopaPersonalVerifyVerificationLaserCodeApiPersonalGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DopaPersonalVerifyVerificationLaserCodeApiPersonalGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **contentType** | **optional.String**|  | [default to application/x-www-form-urlencoded; charset&#x3D;utf-8]

### Return type

[**Verify**](Verify.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

