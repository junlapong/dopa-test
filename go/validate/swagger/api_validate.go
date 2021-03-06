
/*
 * Validate
 *
 * Obtain authentication token for further api requests. / ยืนยันตนกับ DGA API Platform เพื่อใช้ในการเข้าถึง API ตามสิทธิที่หน่วยงานได้รับ
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ValidateApiService service

/*
ValidateApiService Validate
ตรวจสอบความถูกต้องของ Consumer Key และ Consumer Secret เพื่อขอ Token และเปิด Session ในการเรียกใช้ Service 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param consumerKey ชุดรหัสที่ สรอ. ออกให้ เพื่อความปลอดภัยในการเรียกใช้งาน API
 * @param consumerSecret ชุดรหัสที่ สรอ. ออกให้ เพื่อความปลอดภัยในการเรียกใช้งาน API
 * @param agentID - เลขบัตรประชาชน 13 หลัก ของผู้ที่เข้ามาดูข้อมูลในระบบ   - เป็นค่าที่สามารถบ่งบอกได้ว่า User ใดที่เข้ามาใช้งาน Service ใน Session นั้น ๆ - เป็นค่า Unique ID ของแต่ละ Session โดยทาง สรอ. กำหนดไว้ให้  
 * @param optional nil or *ValidateApiValidateGetOpts - Optional Parameters:
     * @param "ContentType" (optional.String) - 

@return TokenResult
*/

type ValidateApiValidateGetOpts struct { 
	ContentType optional.String
}

func (a *ValidateApiService) ValidateGet(ctx context.Context, consumerKey string, consumerSecret string, agentID string, localVarOptionals *ValidateApiValidateGetOpts) (TokenResult, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue TokenResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("ConsumerSecret", parameterToString(consumerSecret, ""))
	localVarQueryParams.Add("AgentID", parameterToString(agentID, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Consumer-Key"] = parameterToString(consumerKey, "")
	if localVarOptionals != nil && localVarOptionals.ContentType.IsSet() {
		localVarHeaderParams["Content-Type"] = parameterToString(localVarOptionals.ContentType.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v TokenResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

