package main

import (
	"fmt"

	validate "github.com/junlapong/dopa-test/validate/swagger"
	verify "github.com/junlapong/dopa-test/verify/swagger"
)

func main() {
	ver := verify.APIClient{}
	// c.DopaPersonalVerifyVerificationLaserCodeApi.PersonalGet()
	fmt.Printf("%T\n", ver)

	val := validate.APIClient{}
	fmt.Printf("%T\n", val)

}
