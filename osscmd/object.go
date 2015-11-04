/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/object"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	objectHelp = `
	Usage: object [] 
	   	use object and options to list/new/qurey etc 
	Commands:
		-c : create a bucket
		-s : set bucket's attributes
		-d : --logging --website --lifecycle
		-q : query bucket's attributes --logging --website --lifecycle --acl --location --referer

		-p : permission
		--encoding: encoding of content
		--expire: expire time format is "Fri, 28 Feb 2012 05:38:42 GMT"
		--file: file to upload
		--type: content file type
	`
)

type ObjectFlagInfo struct {
	Encoding string
	Expire   string
	File     string
	Type     string
}

var (
	ObjectFlag ObjectFlagInfo
)

func init() {
	flag.StringVar(&ObjectFlag.Encoding, "encoding", "", "encoding of content")
	flag.StringVar(&ObjectFlag.Expire, "expire", "", "expire time format is `Fri, 28 Feb 2012 05:38:42 GMT`")
	flag.StringVar(&ObjectFlag.File, "file", "", "file to upload")
	flag.StringVar(&ObjectFlag.Type, "type", "", "content file type")
}

func Object(args []string) (err error) {
	if err = readCfg(); err != nil {
		fmt.Println("You May Havn't Init . Use osscmd init First!")
		os.Exit(0)
	}
	ossapi.Init(accessKeyID, accessKeySecret)
	var e *ossapi.Error
	if "-c" == args[2] {
		flag.CommandLine.Parse(args[3:])
		var loc, per string
		if BucketFlag.Location == "hangzhou" {
			loc = bucket.L_Hangzhou
		} else if BucketFlag.Location == "beijin" {
			loc = bucket.L_Beijing
		} else if BucketFlag.Location == "shenzhen" {
			loc = bucket.L_Shenzhen
		} else if BucketFlag.Location == "hongkong" {
			loc = bucket.L_HongKong
		} else if BucketFlag.Location == "qingdao" {
			loc = bucket.L_Qingdao
		} else if BucketFlag.Location == "shanghai" {
			loc = bucket.L_Shanghai
		} else {
			loc = bucket.L_Hangzhou
		}
		if BucketFlag.Permission == "RW" {
			per = bucket.P_PublicRW
		} else if BucketFlag.Permission == "PT" {
			per = bucket.P_Private
		} else if BucketFlag.Permission == "RO" {
			per = bucket.P_PublicReadOnly
		} else {
			per = bucket.P_Private
		}

		fd, err := os.Open(ObjectFlag.File)
		if err != nil {
			Exit(e.Error())
		}

		body, err := ioutil.ReadAll(fd)
		if err != nil {
			Exit(e.Error())
		}

		objInfo := &object.ObjectInfo{
			CacheControl:       "no-cache",
			ContentDisposition: "attachment;",
			ContentEncoding:    ObjectFlag.Encoding,
			Expires:            ObjectFlag.Expire,
			Encryption:         "AES256",
			ACL:                per,
			ObjName:            filepath.Base(ObjectFlag.File),
			Location:           loc,
			Body:               body,
			Type:               ObjectFlag.Type,
			BucketName:         BucketFlag.Bucket}
		e = object.Create(objInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Create Object" + ObjectFlag.File + " Success !")
	} else if "-q" == args[2] {

	} else {
		fmt.Println(objectHelp)
		os.Exit(-1)
	}
	return
}
