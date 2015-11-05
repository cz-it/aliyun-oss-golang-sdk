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
	    -q : Query Object's data
		-n : create a object 
		-c : copy a object 
		-a : append object
		-d : delete object
		-p : permission
		-s : set object's attribute
		-m: query object's brief meta info

		--object: name of object
		--position: position to append
		--encoding: encoding of content
		--source : source object of copy
		--directive: copy function: REPLACE/COPY
		--expire: expire time format is "Fri, 28 Feb 2012 05:38:42 GMT"
		--file: file to upload
		--type: content file type
	`
)

type ObjectFlagInfo struct {
	Encoding  string
	Expire    string
	File      string
	Type      string
	Directive string
	Source    string
	Position  int
	Head      bool
}

var (
	ObjectFlag ObjectFlagInfo
)

func init() {
	flag.StringVar(&ObjectFlag.Encoding, "encoding", "", "encoding of content")
	flag.StringVar(&ObjectFlag.Expire, "expire", "", "expire time format is `Fri, 28 Feb 2012 05:38:42 GMT`")
	flag.StringVar(&ObjectFlag.File, "file", "", "file to upload")
	flag.StringVar(&ObjectFlag.Type, "type", "", "content file type")
	flag.StringVar(&ObjectFlag.Directive, "directive", "COPY", "copy function: REPLACE/COPY")
	flag.StringVar(&ObjectFlag.Source, "source", "", "source object of copy")
	flag.IntVar(&ObjectFlag.Position, "position", 0, "content file type")
	flag.BoolVar(&ObjectFlag.Head, "head", false, "")
}

func Object(args []string) (err error) {
	if err = readCfg(); err != nil {
		fmt.Println("You May Havn't Init . Use osscmd init First!")
		os.Exit(0)
	}
	ossapi.Init(accessKeyID, accessKeySecret)
	var e *ossapi.Error
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
	if "-n" == args[2] {
		fd, err := os.Open(ObjectFlag.File)
		if err != nil {
			Exit(err.Error())
		}
		body, err := ioutil.ReadAll(fd)
		if err != nil {
			Exit(err.Error())
		}
		objInfo := &object.ObjectInfo{
			CacheControl:       "no-cache",
			ContentDisposition: "attachment;",
			ContentEncoding:    ObjectFlag.Encoding,
			Expires:            ObjectFlag.Expire,
			Encryption:         "AES256",
			ACL:                per,
			Body:               body,
			Type:               ObjectFlag.Type}
		e = object.Create(filepath.Base(ObjectFlag.File), BucketFlag.Bucket, loc, objInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Create Object" + ObjectFlag.File + " Success !")
	} else if "-c" == args[2] {
		copyInfo := &object.CopyInfo{
			ObjectName: CORSFlag.Object,
			BucketName: BucketFlag.Bucket,
			Location:   loc,
			Source:     ObjectFlag.Source,
			Directive:  ObjectFlag.Directive,
			ACL:        per,
			Encryption: "AES256"}
		info, e := object.Copy(copyInfo, nil)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Copy Object Success , New Object is ", info)
	} else if "-a" == args[2] {
		fd, err := os.Open(ObjectFlag.File)
		if err != nil {
			Exit(e.Error())
		}
		body, err := ioutil.ReadAll(fd)
		if err != nil {
			Exit(e.Error())
		}
		objInfo := &object.AppendObjInfo{ObjectInfo: object.ObjectInfo{
			CacheControl:       "no-cache",
			ContentDisposition: "attachment;",
			ContentEncoding:    ObjectFlag.Encoding,
			Expires:            ObjectFlag.Expire,
			Encryption:         "AES256",
			Body:               body,
			ACL:                per,
			Type:               ObjectFlag.Type},
			Position: uint64(ObjectFlag.Position)}
		info, e := object.Append(CORSFlag.Object, BucketFlag.Bucket, loc, objInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Append Success. resuult:", info)
	} else if "-d" == args[2] {
		e = object.Delete(CORSFlag.Object, BucketFlag.Bucket, loc)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Delete " + CORSFlag.Object + " Success ")
	} else if "-m" == args[2] {
		info, e := object.QueryMeta(CORSFlag.Object, BucketFlag.Bucket, loc, nil)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Meta "+CORSFlag.Object+"  is ", info)
	} else if "-s" == args[2] {
		e := object.SetACL(CORSFlag.Object, BucketFlag.Bucket, loc, per)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Set Object's ACL Success")
	} else if "-q" == args[2] {
		if BucketFlag.IsACL {
			info, e := object.QueryACL(CORSFlag.Object, BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Object's ACL Owner:", info.Owner)
			fmt.Println("Object's ACL:", info.AccessControlList)
		} else {
			data, e := object.Query(CORSFlag.Object, BucketFlag.Bucket, loc, nil, nil)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Data:", string(data))
		}
	} else {
		fmt.Println(objectHelp)
		os.Exit(-1)
	}
	return
}
