/**
* Author: CZ cz.theng@gmail.com
 */

package bucket

import (
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"testing"
)

func TestSetBucketWebsite(t *testing.T) {
	if nil != ossapi.Init("v8P430U3UcILP6KA", "EB9v8yL2aM07YOgtO1BdfrXtdxa4A1") {
		t.Fail()
	}
	if err := SetWebsite("test-put-bucket4", LBeijing, "index.html", ""); err != nil {
		fmt.Println(err.ErrNo, err.HTTPStatus, err.ErrMsg, err.ErrDetailMsg)
	} else {
		t.Log("SetBucketWebSite Success")
	}
}
