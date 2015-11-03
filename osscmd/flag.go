/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"os"
)

var Flag FlagInfo

const (
	helpStr = `
	osscmd is a tool for ossapi.

	Usage:

		osscmd command [arguments]

		The commands are:

		    init : init osscmd's evniroment 
			bucket: buckets' tool
	`
	initHelpStr = `
	Usage: init -i id -s secort 
	    init osscmd with Access Key ID id and Access Key Secort secort
	Commands:
		-i id : Access Key ID
		-s secort : Access Key Secort
		-h/--hlep : Show Help

	`
)

type FlagInfo struct {
	Version bool
	ID      string
	Secort  string
}

func init() {
	flag.Usage = func() {
		fmt.Printf("OSS API's Command Tool\n")
		fmt.Println(helpStr)
	}

	flag.BoolVar(&Flag.Version, "v", false, "Show OSS API Command Tool's Version")
	flag.StringVar(&Flag.ID, "i", "", " Access Key ID")
	flag.StringVar(&Flag.Secort, "s", "", " Access Key Secort")
}

func Usage() {
	fmt.Println(helpStr)
	os.Exit(0)
}

func InitUsage() {
	fmt.Println(initHelpStr)
	os.Exit(0)
}

func parseArgs() {
	var err error
	if len(os.Args) < 2 {
		Usage()
	}
	if "-h" == os.Args[1] || "help" == os.Args[1] {
		Usage()
	} else if "-v" == os.Args[1] || "version" == os.Args[1] {
		fmt.Println("Cur Version:", ossapi.Version())
		return
	} else if "init" == os.Args[1] {
		if len(os.Args) < 3 {
			Usage()
		}
		if "-h" == os.Args[2] || "--help" == os.Args[2] {
			InitUsage()
		}
		flag.CommandLine.Parse(os.Args[2:])
		if Flag.ID == "" || Flag.Secort == "" {
			flag.Usage()
			return
		}
		if err = writeCfg(Flag.ID, Flag.Secort); err != nil {
			fmt.Println("[FAIL]: Init Error ", err.Error())
		} else {
			fmt.Println("[SUCC]: Init Success!")
		}
		return
	} else if "bucket" == os.Args[1] {
		if len(os.Args) < 3 {
			Usage()
		}
		if err = Bucket(os.Args[:]); err != nil {
			Usage()
		}
	} else if "cors" == os.Args[1] {
		if len(os.Args) < 3 {
			Usage()
		}
		if err = CORS(os.Args[:]); err != nil {
			Usage()
		}
	} else {
		Usage()
	}
}
