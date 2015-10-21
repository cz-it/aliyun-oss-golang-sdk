/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import ()

func signature(req *Request) (sig string, err error) {
	sigStr := "ET\n"
	sigStr += "" + "\n"
	sigStr += "\n"
	sigStr += req.Date + "\n"
	sigStr += "/"
	sig, err = Base64AndHmacSha1([]byte(accessKeySecret), []byte(sigStr))
	if err != nil {
		return
	}
	return
}

type Buckets struct {
}

func GetServiceWith(prefix, marker string, maxKeys int) (buckets *Buckets, err error) {
	args := ""
	path := "/"
	if "" != prefix {
		args += "prefix=" + prefix
	}
	if "" != marker {
		args += "marker=" + marker
	}
	if 0 < maxKeys && maxKeys <= 1000 {
		args += "maxkeys=" + string(maxKeys)
	}

	if "" != args {
		path += "?" + args
	}
	req := &Request{Host: "oss.aliyuncs.com", Path: "/"}
	rsp, err := req.Send(signature)
	/*
		body := make([]byte, 10000)
		rsp.httpRsp.Body.Read(body)
		fmt.Println(string(body))
	*/
	if rsp.Result != ESUCC {
		return
	}
	return
}

func GetService() (buckets *Buckets, err error) {
	buckets, err = GetServiceWith("", "", 0)
	return
}
