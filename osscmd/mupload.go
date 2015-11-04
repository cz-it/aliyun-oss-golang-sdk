/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/bucket"
	"github.com/cz-it/aliyun-oss-golang-sdk/ossapi/mupload"
	"io/ioutil"
	"os"
	"strings"
)

const (
	muploadHelp = `
	Usage: mupload 
	 	Upload multiple parts to a object. Such as a big file.	
	Commands:
		-i : create bucket's cors
		-a : add upload part
		-c : complete

		--file : data from file
		--number: part's number
	`
	muploadCfgFile = ".osscmd/mupload.cfg"
)

type MuploadFlagInfo struct {
	ID     string
	Number int
}

var (
	MuploadFlag MuploadFlagInfo
)

func init() {
	flag.StringVar(&MuploadFlag.ID, "id", "", "upload context's ID")
	flag.IntVar(&MuploadFlag.Number, "number", 0, "part's number")
}

type PartInfo struct {
	Number int
	ETag   string
}

type UploadInfo struct {
	ID       string
	Key      string
	PartInfo []PartInfo
}

func Mupload(args []string) (err error) {
	if err = readCfg(); err != nil {
		fmt.Println("You May Havn't Init . Use osscmd init First!")
		os.Exit(0)
	}
	ossapi.Init(accessKeyID, accessKeySecret)
	//var e *ossapi.Error
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
	if "-i" == args[2] {
		initInfo := &mupload.InitInfo{
			CacheControl:       "no-cache",
			ContentDisposition: "attachment",
			ContentEncoding:    ObjectFlag.Encoding,
			Expires:            ObjectFlag.Expire,
			Encryption:         "AES256"}
		info, e := mupload.Init(CORSFlag.Object, BucketFlag.Bucket, loc, initInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Init Mupload Success:")
		fmt.Println("Key is:", info.Key)
		fmt.Println("Id is: ", info.UploadId)
		rstData := &UploadInfo{
			ID:  info.UploadId,
			Key: info.Key}
		body, err := xml.MarshalIndent(rstData, "", "  ")
		if err != nil {
			Exit(err.Error())
		}
		fd, err := os.Create(muploadCfgFile)
		defer fd.Close()
		if err != nil {
			Exit(err.Error())
		}
		fd.Write(append([]byte(xml.Header), body...))
	} else if "-c" == args[2] {
		cfgFd, err := os.Open(muploadCfgFile)
		if err != nil {
			Exit(err.Error())
		}
		cfgData, err := ioutil.ReadAll(cfgFd)
		defer cfgFd.Close()
		if err != nil {
			Exit(err.Error())
		}
		cfgInfo := new(UploadInfo)
		err = xml.Unmarshal(cfgData, cfgInfo)
		if err != nil {
			Exit(err.Error())
		}
		cfgFd.Close()
		var parts []mupload.PartInfo
		for _, i := range cfgInfo.PartInfo {
			parts = append(parts, mupload.PartInfo{ETag: i.ETag, PartNumber: i.Number})
		}
		partsInfo := &mupload.PartsInfo{Part: parts}
		rstInfo, e := mupload.Complete(CORSFlag.Object, BucketFlag.Bucket, loc, cfgInfo.ID, partsInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Success Done Complete ETag:", rstInfo.ETag, "  Key:", rstInfo.Key)
	} else if "-a" == args[2] {
		fd, err := os.Open(ObjectFlag.File)
		defer fd.Close()
		if err != nil {
			Exit(err.Error())
		}
		body, err := ioutil.ReadAll(fd)
		if err != nil {
			Exit(err.Error())
		}
		cfgFd, err := os.Open(muploadCfgFile)
		if err != nil {
			Exit(err.Error())
		}
		cfgData, err := ioutil.ReadAll(cfgFd)
		defer cfgFd.Close()
		if err != nil {
			Exit(err.Error())
		}
		cfgInfo := new(UploadInfo)
		err = xml.Unmarshal(cfgData, cfgInfo)
		if err != nil {
			Exit(err.Error())
		}
		cfgFd.Close()
		partInfo := &mupload.UploadPartInfo{
			ObjectName: CORSFlag.Object,
			BucketName: BucketFlag.Bucket,
			Location:   loc,
			UploadID:   cfgInfo.ID,
			PartNumber: MuploadFlag.Number,
			Data:       body,
			CntType:    ObjectFlag.Type}
		info, e := mupload.Append(partInfo)
		if e != nil {
			Exit(e.Error())
		}
		fmt.Println("Add Part Success, Tag is ", info.Etag)
		cfgFd, err = os.Create(muploadCfgFile)
		defer cfgFd.Close()
		if err != nil {
			Exit(err.Error())
		}
		cfgInfo.PartInfo = append(cfgInfo.PartInfo, PartInfo{Number: MuploadFlag.Number, ETag: strings.Trim(info.Etag, "\"")})
		cfgData, err = xml.MarshalIndent(cfgInfo, "", "  ")
		if err != nil {
			Exit(err.Error())
		}
		cfgFd.Write(append([]byte(xml.Header), cfgData...))
	} else {
		fmt.Println(corsHelp)
		os.Exit(-1)
	}
	return
}
