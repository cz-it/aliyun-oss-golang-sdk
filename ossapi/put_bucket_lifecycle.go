/**
* Author: CZ cz.theng@gmail.com
 */

package ossapi

import (
	"encoding/xml"
	"path"
)

const (
	LifecycleStatsEnable  = "Enabled"
	LifecycleStatsDisable = "Disabled"
)

type ExpirationDaysInfo struct {
	Days uint
}

type ExpirationDateInfo struct {
	Date string
}

type RuleInfo struct {
	ID         string
	Prefix     string
	Status     string
	Expiration ExpirationDaysInfo
}

type LifecycleConfiguration struct {
	XMLName xml.Name `xml:"LifecycleConfiguration"`
	Rule    []RuleInfo
}

func SetBucketLifecycle(name, location string, rules []RuleInfo) (ossapiError *Error) {
	host := name + "." + location + ".aliyuncs.com"
	resource := path.Join("/", name)
	info := LifecycleConfiguration{Rule: rules}
	body, err := xml.Marshal(info)
	if err != nil {
		Logger.Error("err := xml.Marshal(Info) Error %s", err.Error())
		ossapiError = OSSAPIError
		return
	}

	body = append([]byte(xml.Header), body...)
	req := &Request{
		Host:     host,
		Path:     "/?lifecycle",
		Method:   "PUT",
		Resource: resource + "/",
		SubRes:   []string{"lifecycle"},
		Body:     body,
		CntType:  "application/xml"}
	rsp, err := req.Send()
	if err != nil {
		if _, ok := err.(*Error); !ok {
			Logger.Error("GetService's Send Error:%s", err.Error())
			ossapiError = OSSAPIError
			return
		}
	}
	if rsp.Result != ESUCC {
		ossapiError = err.(*Error)
		return
	}

	return
}
