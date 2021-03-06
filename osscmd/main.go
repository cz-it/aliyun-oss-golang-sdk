/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"encoding/xml"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func init() {
	usr, err := user.Current()
	if err != nil {
		return
	}
	configPath = path.Join(usr.HomeDir, configPathFile)
}

const (
	configPathFile = ".osscmd/config.xml"
)

// ConfigInfo is configure
type ConfigInfo struct {
	XMLName         xml.Name `xml:"Config"`
	AccessKeyID     string   `xml:"AccessKeyID"`
	AccessKeySecret string   `xml:"AccessKeySecret"`
}

var (
	configPath      string
	accessKeyID     string
	accessKeySecret string
)

func readCfg() (err error) {
	cfg := new(ConfigInfo)
	cfgData := make([]byte, 10240)
	fd, err := os.Open(configPath)
	_, err = fd.Read(cfgData)
	if err != nil {
		return
	}
	err = xml.Unmarshal(cfgData, cfg)
	if err != nil {
		return
	}
	accessKeyID = cfg.AccessKeyID
	accessKeySecret = cfg.AccessKeySecret
	return
}

func writeCfg(ID, Secret string) (err error) {
	fd, err := os.Create(configPath)
	if err != nil {
		return
	}
	cfg := &ConfigInfo{AccessKeyID: ID, AccessKeySecret: Secret}
	cfgData, err := xml.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return
	}
	_, err = fd.Write(cfgData)
	if err != nil {
		return
	}
	return
}

func main() {
	dirPath := filepath.Dir(configPath)
	os.MkdirAll(dirPath, 0777)
	parseArgs()
}
