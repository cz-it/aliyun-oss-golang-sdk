/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"errors"
	//"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/service"
	"os"
)

var (
	ArrErr = errors.New("Arg Error")
)

type BucketFlagInfo struct {
	ID     string
	Secort string
}

func init() {
	//flag.StringVar(&Flag.Secort, "s", "", " Access Key Secort")
}

func Bucket(args []string) (err error) {
	if "-l" == args[2] {
		if err = readCfg(); err != nil {
			fmt.Println("You May Havn't Init . Use osscmd init First!")
			os.Exit(0)
		}

		ossapi.Init(accessKeyID, accessKeySecret)
		info, e := service.GetService()
		fmt.Println("id:", accessKeyID, " And key:", accessKeySecret)
		fmt.Println("info:", info)
		fmt.Println("err:", e)
	}
	return
}
