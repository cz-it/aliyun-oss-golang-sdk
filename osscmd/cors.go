/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/cors"
	"os"
)

const (
	corsHelp = `
	Usage: cors 
		Set Bucket's CORS
	Commands:
		-c : create bucket's cors
		-q : query bucket's cors
		-d : delete bucket's cors
		-o : option bucket's cors

		--origin: option's origin 
		--method: option's method
		--headers: option's headers
		--object : object's name
	`
)

// CORSFlagInfo is cors flag
type CORSFlagInfo struct {
	Bucket   string
	Location string
	Origin   string
	Method   string
	Headers  string
	Object   string
}

var (
	// CORSFlag is flag
	CORSFlag CORSFlagInfo
)

func init() {
	flag.StringVar(&CORSFlag.Origin, "origin", "", "option's origin")
	flag.StringVar(&CORSFlag.Method, "method", "", "option's method")
	flag.StringVar(&CORSFlag.Headers, "headers", "", "option's headers")
	flag.StringVar(&CORSFlag.Object, "object", "", "obejct's name")
}

// CORS is cors cmd
func CORS(args []string) (err error) {
	if err = readCfg(); err != nil {
		fmt.Println("You May Havn't Init . Use osscmd init First!")
		os.Exit(0)
	}
	ossapi.Init(accessKeyID, accessKeySecret)
	//var e *ossapi.Error
	if "-q" == args[2] {
		flag.CommandLine.Parse(args[3:])
		var loc string
		if BucketFlag.Location == "hangzhou" {
			loc = bucket.LHangzhou
		} else if BucketFlag.Location == "beijin" {
			loc = bucket.LBeijing
		} else if BucketFlag.Location == "shenzhen" {
			loc = bucket.LShenzhen
		} else if BucketFlag.Location == "hongkong" {
			loc = bucket.LHongKong
		} else if BucketFlag.Location == "qingdao" {
			loc = bucket.LQingdao
		} else if BucketFlag.Location == "shanghai" {
			loc = bucket.LShanghai
		} else {
			loc = bucket.LHangzhou
		}
		info, e := cors.Query(BucketFlag.Bucket, loc)
		if e != nil {
			Exit(e.Error())
		}
		for idx, rule := range info {
			fmt.Printf("Rule [%d]: \n", idx)
			fmt.Println("    AllowedOrigin:", rule.AllowedOrigin)
			fmt.Println("    AllowedMethod:", rule.AllowedMethod)
			fmt.Println("    AllowedHeader:", rule.AllowedHeader)
			fmt.Println("    ExposeHeader:", rule.ExposeHeader)
			fmt.Println("    MaxAgeSeconds:", rule.MaxAgeSeconds)
		}
	} else if "-d" == args[2] {
		flag.CommandLine.Parse(args[3:])
		var loc string
		if BucketFlag.Location == "hangzhou" {
			loc = bucket.LHangzhou
		} else if BucketFlag.Location == "beijin" {
			loc = bucket.LBeijing
		} else if BucketFlag.Location == "shenzhen" {
			loc = bucket.LShenzhen
		} else if BucketFlag.Location == "hongkong" {
			loc = bucket.LHongKong
		} else if BucketFlag.Location == "qingdao" {
			loc = bucket.LQingdao
		} else if BucketFlag.Location == "shanghai" {
			loc = bucket.LShanghai
		} else {
			loc = bucket.LHangzhou
		}
		e := cors.Delete(BucketFlag.Bucket, loc)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Delete CORS Success.")
	} else if "-o" == args[2] {
		flag.CommandLine.Parse(args[3:])
		var loc string
		if BucketFlag.Location == "hangzhou" {
			loc = bucket.LHangzhou
		} else if BucketFlag.Location == "beijin" {
			loc = bucket.LBeijing
		} else if BucketFlag.Location == "shenzhen" {
			loc = bucket.LShenzhen
		} else if BucketFlag.Location == "hongkong" {
			loc = bucket.LHongKong
		} else if BucketFlag.Location == "qingdao" {
			loc = bucket.LQingdao
		} else if BucketFlag.Location == "shanghai" {
			loc = bucket.LShanghai
		} else {
			loc = bucket.LHangzhou
		}
		reqInfo := &cors.OptionReqInfo{
			Origin:  CORSFlag.Origin,
			Headers: CORSFlag.Headers,
			Method:  CORSFlag.Method}
		info, e := cors.Option(CORSFlag.Object, BucketFlag.Bucket, loc, reqInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Option Info:", info)
	} else {
		fmt.Println(corsHelp)
		os.Exit(-1)
	}
	return
}
