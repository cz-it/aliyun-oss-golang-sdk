/**
* Author: CZ cz.theng@gmail.com
 */
package ossapi

import (
	"fmt"
	"testing"
)

func TestGetBucket(t *testing.T) {
	if nil != Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := GetBucket("test-put-bucket3", L_Beijing, "", "", "", "", 0); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("GetBucket Success")
		fmt.Println(info)
	}
}
