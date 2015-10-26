/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	ArgError.Error()
}

func TestDo(t *testing.T) {
	req := &Request{Host: "oss.aliyuncs.com", Path: "/", Method: "GET", Resource: "/"}
	rsp, err := req.Send()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
