/*
 * Validate
 *
 * Obtain authentication token for further api requests. / ยืนยันตนกับ DGA API Platform เพื่อใช้ในการเข้าถึง API ตามสิทธิที่หน่วยงานได้รับ
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TokenResult struct {
	// TOKEN ที่ระบบของ สรอ. สร้างให้และส่งกลับมา เพื่อนำไปใช้ในการพัฒนา API เรียกใช้ข้อมูล
	Result string `json:"Result,omitempty"`
}
