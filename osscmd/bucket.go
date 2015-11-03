/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service"
	"os"
)

const (
	bucketHelp = `
	Usage: bucket [-l] 
	   	use bucket and options to list/new/qurey etc 
	Commands:
		-l : list bucket

		-c : create a bucket
		-s : set bucket's attributes
		-d : --logging --website --lifecycle
		-q : query bucket's attributes --logging --website --lifecycle --acl --location --referer

		-b : bucket's name
		-a : area such as hangzhou/shenzhen/beijin
		-g : set logging  --log-possition --log-prefix
		-w : set website  --index --error
		-r : set referer
	`
)

type BucketFlagInfo struct {
	Bucket       string
	Location     string
	Permission   string
	LogPossition string
	Log          bool
	IsLog        bool
	LogPrefix    string
	Website      bool
	IsWebsite    bool
	WebIndex     string
	WebError     string
	Referer      bool
	IsLifecycle  bool
	IsLocation   bool
	IsACL        bool
	IsReferer    bool
}

var (
	ArrErr     = errors.New("Arg Error")
	BucketFlag BucketFlagInfo
)

func init() {
	flag.StringVar(&BucketFlag.Bucket, "b", "", "Bucket Name")
	flag.StringVar(&BucketFlag.Location, "a", "", "Area Name such as hangzhou/beijin/shenzhen")
	flag.StringVar(&BucketFlag.Permission, "p", "", "Permission such as RO/RW/PT")
	flag.StringVar(&BucketFlag.LogPossition, "log-possition", "", "Which bucket to store logs")
	flag.StringVar(&BucketFlag.LogPrefix, "log-prefix", "", "log file's prefix name")
	flag.BoolVar(&BucketFlag.Log, "g", false, "set log ")
	flag.BoolVar(&BucketFlag.Referer, "r", false, "set referer")
	flag.BoolVar(&BucketFlag.Website, "w", false, "set website")
	flag.BoolVar(&BucketFlag.IsWebsite, "website", false, "delete/get website")
	flag.BoolVar(&BucketFlag.IsLog, "logging", false, "delete/get log")
	flag.BoolVar(&BucketFlag.IsLifecycle, "lifecycle", false, "delete/get lifecycle")
	flag.StringVar(&BucketFlag.WebIndex, "index", "", "index page")
	flag.StringVar(&BucketFlag.WebError, "error", "", "error page")
	flag.BoolVar(&BucketFlag.IsLocation, "location", false, "get location")
	flag.BoolVar(&BucketFlag.IsACL, "acl", false, "get acl info")
	flag.BoolVar(&BucketFlag.IsReferer, "referer", false, "get referer info")
}

func Exit(msg string) {
	fmt.Println("[FAIL]:", msg)
	os.Exit(-1)
}

func Bucket(args []string) (err error) {
	if err = readCfg(); err != nil {
		fmt.Println("You May Havn't Init . Use osscmd init First!")
		os.Exit(0)
	}
	ossapi.Init(accessKeyID, accessKeySecret)
	var e *ossapi.Error
	if "-l" == args[2] {
		info, e := service.QueryBucketsDefault()
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Owner:", info.Owner.DisplayName)
		for idx, bkt := range info.Buckets.Bucket {
			fmt.Printf("	Bucket[%v]:", idx)
			fmt.Println(bkt)
		}
	} else if "-c" == args[2] {
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
		flag.CommandLine.Parse(args[3:])
		var loc string
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
		if BucketFlag.IsACL {
			info, e := bucket.QueryACL(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Owner is :", info.Owner)
			fmt.Println("Grant is :", info.AccessControlList.Grant)
		} else if BucketFlag.IsLocation {
			info, e := bucket.QueryLocation(BucketFlag.Bucket)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Location is ", info)
		} else if BucketFlag.IsLifecycle {
			info, e := bucket.QueryLifecycle(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			for idx, rule := range info {
				fmt.Printf("Rule [%d]:", idx)
				fmt.Println(rule)
			}
		} else if BucketFlag.IsLog {
			info, e := bucket.QueryLogging(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			if info == nil {
				fmt.Println("Bucket has not config logging")

			} else {
				fmt.Println("Target Bucket", info.TargetBucket)
				fmt.Println("Target Prefix", info.TargetPrefix)
			}
		} else if BucketFlag.IsReferer {
			info, e := bucket.QueryReferer(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("If Allow Empyt", info.AllowEmptyReferer)
			fmt.Println("White List", info.RefererList.Referer)
		} else if BucketFlag.IsWebsite {
			info, e := bucket.QueryWebsite(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Index is :", info.IndexDocument.Suffix)
			fmt.Println("404page is :", info.ErrorDocument.Key)
		} else {
			info, e := bucket.QueryObjects(BucketFlag.Bucket, loc, "", "", "", "", 0)
			if e != nil {
				Exit(e.Error())
			}
			for idx, obj := range info.Contents {
				fmt.Printf("Objects [%d]", idx)
				fmt.Println(" are:", obj)
			}
		}
	} else if "-d" == args[2] {
		flag.CommandLine.Parse(args[3:])
		var loc string
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
		if BucketFlag.IsLog {
			e = bucket.DeleteLogging(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Delete Bucket " + BucketFlag.Bucket + "  Logging Success !")
		} else if BucketFlag.IsWebsite {
			e = bucket.DeleteWebsite(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Delete Bucket " + BucketFlag.Bucket + "  Website Success !")
		} else if BucketFlag.IsLifecycle {
			e = bucket.DeleteLifecycle(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Delete Bucket " + BucketFlag.Bucket + "  LifeCycle Success !")
		} else {
			e = bucket.Delete(BucketFlag.Bucket, loc)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Delete Bucket " + BucketFlag.Bucket + " Success !")
		}
	} else if "-s" == args[2] {
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
		if BucketFlag.Permission != "" {
			e = bucket.SetACL(BucketFlag.Bucket, loc, per)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Set Bucket " + BucketFlag.Bucket + ":" + BucketFlag.Permission + " Success !")
		} else if BucketFlag.Log {

			if BucketFlag.LogPossition == "" {
				e = bucket.CloseLogging(BucketFlag.Bucket, loc)
			} else {
				e = bucket.OpenLogging(BucketFlag.Bucket, loc, BucketFlag.LogPossition, BucketFlag.LogPrefix)
			}
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Set Bucket " + BucketFlag.Bucket + "log  Success !")
		} else if BucketFlag.Website {
			e = bucket.SetWebsite(BucketFlag.Bucket, loc, BucketFlag.WebIndex, BucketFlag.WebError)
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Set Bucket " + BucketFlag.Bucket + " website  Success !")
		} else if BucketFlag.Referer {
			if len(flag.Args()) > 0 {
				e = bucket.SetReferer(BucketFlag.Bucket, loc, true, flag.Args())
			} else {
				e = bucket.SetReferer(BucketFlag.Bucket, loc, true, flag.Args())
			}
			if e != nil {
				Exit(e.Error())
			}
			fmt.Println("Set Bucket " + BucketFlag.Bucket + " refer Success !")
		} else {
			fmt.Println(bucketHelp)
			os.Exit(-1)
		}

	} else {
		fmt.Println(bucketHelp)
		os.Exit(-1)
	}
	return
}
