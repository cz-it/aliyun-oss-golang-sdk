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
	headers := map[string]string{"Content-XX": "adf"}
	req := &Request{
		Host:      "oss.aliyuncs.com",
		Path:      "/",
		ExtHeader: headers,
		SubRes:    []string{"aa"},
		Body:      []byte("aaa"),
		CntType:   "text/html",
		Method:    "GET",
		Resource:  "/"}
	req.AddXOSS("oss-xx", "abc")
	rsp, err := req.Send()

	fmt.Println("====Method Error=======")
	req.Method = "error"
	rsp, err = req.Send()

	fmt.Println("====URL Error=======")
	req.Host = "nimeidenimei"
	rsp, err = req.Send()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
