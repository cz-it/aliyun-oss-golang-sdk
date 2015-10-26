/**
* Author: CZ cz.theng@gmail.com
 */
package bucket

import (
	"fmt"
	"testing"
)

func TestGetBucketLocation(t *testing.T) {
	if nil != Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := GetBucketLocation("test-put-bucket2"); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetBucketLocation Success")
		fmt.Println(info)
	}
}
