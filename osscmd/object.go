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
	"os"
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

	`
)

type ObjectFlagInfo struct {
	Bucket      string
	Location    string
	Permission  string
	Log         bool
	IsLog       bool
	LogPrefix   string
	Website     bool
	IsWebsite   bool
	WebIndex    string
	WebError    string
	Referer     bool
	IsLifecycle bool
	IsLocation  bool
	IsACL       bool
	IsReferer   bool
}

var (
	ObjectFlag ObjectFlagInfo
)

func init() {
	flag.StringVar(&ObjectFlag.Bucket, "b", "", "Bucket Name")
	flag.StringVar(&ObjectFlag.Location, "a", "", "Area Name such as hangzhou/beijin/shenzhen")
	flag.StringVar(&ObjectFlag.Permission, "p", "", "Permission such as RO/RW/PT")
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
		e = bucket.Create(BucketFlag.Bucket, loc, per)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Create Bucket " + BucketFlag.Bucket + " Success !")
	} else if "-q" == args[2] {

	} else {
		fmt.Println(bucketHelp)
		os.Exit(-1)
	}
	return
}
