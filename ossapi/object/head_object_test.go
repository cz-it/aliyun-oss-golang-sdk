/**
* Author: CZ cz.theng@gmail.com
 */
package object

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestBriefObject(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if info, err := BriefObject("append2", "test-object-hz", ossapi.L_Hangzhou, nil); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("CopyObject Success")
		fmt.Println(info)
	}
	if info, err := BriefObject("append2", "test-object-hz", ossapi.L_Hangzhou, &BriefConnInfo{MatchEtag: "append2"}); err != nil {
		fmt.Println(err.ErrNo, err.HttpStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("CopyObject Success")
		fmt.Println(info)
	}
}
