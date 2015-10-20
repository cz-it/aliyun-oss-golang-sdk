/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import ()

func GetService() int {
	req := &Request{Host: "oss.aliyuncs.com", Path: "/"}
	req.Signature()
	req.Send()
	return 0
}
