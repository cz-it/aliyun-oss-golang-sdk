/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service"
	"os"
)

const (
	bucketHelp = `
	Usage: bucket [-l] 
	   	use bucket and options to list/new/qurey etc 
	Commands:
		-l : list bucket
		-b : bucket's name
		-l : location such as hangzhou/shenzhen/beijin
	`
)

type BucketFlagInfo struct {
	Bucket   string
	Location string
}

var (
	ArrErr     = errors.New("Arg Error")
	BucketFlag BucketFlagInfo
)

func init() {
	flag.StringVar(&BucketFlag.Bucket, "b", "", "Bucket Name")
	flag.StringVar(&BucketFlag.Location, "l", "", "Location Name such as hangzhou/beijin/shenzhen")
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

	if "-l" == args[2] {
		info, e := service.GetServiceDefault()
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Owner:", info.Owner.DisplayName)
		for idx, bucket := range info.Buckets.Bucket {
			fmt.Printf("	Bucket[%v]:", idx)
			fmt.Println(bucket)
		}
	}
	return
}
