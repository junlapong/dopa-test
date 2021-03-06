/*
 * กรมการปกครอง / ตรวจสอบข้อมูลโดยใช้ข้อมูลบนบัตรฯ และเลขหลังบัตร
 *
 * หน่วยงานสามารถยืนยันตัวบุคคลโดยใช้ข้อมูลจากบัตรประชาชน และเลขหลังบัตร ผ่าน Government API ที่ สรอ. จัดเตรียมไว้ให้ ดังนี้
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Verify struct {
	// ผลการตรวจสอบ (true คือ ตรงกับฐานข้อมูลทะเบียนราษฎร์ หรือ false คือ ไม่ตรงกับฐานข้อมูลทะเบียนราษฎร์)
	Result bool `json:"Result,omitempty"`
}
