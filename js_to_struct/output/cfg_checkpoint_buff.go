package js_to_struct

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

var cfgCheckpointBuffModel = new(CfgCheckpointBuff)
var cfgCheckpointBuffLock = new(sync.RWMutex)
var cfgCheckpointBuffOnce = new(sync.Once)

type CfgCheckpointBuff []CfgCheckpointBuffInfo

type CfgCheckpointBuffInfo struct {
	Effects []struct {
		Desc   string  `json:"desc"`
		Type   string  `json:"type"`
		Values []int64 `json:"values"`
	} `json:"effects"`
	AttachRole int64  `json:"attach_role"`
	Id         int64  `json:"id"`
	Name       string `json:"name"`
}

func GetCfgCheckpointBuff(jsPath string) CfgCheckpointBuff {
	cfgCheckpointBuffOnce.Do(func() {
		cfgCheckpointBuffModel.loadConfig(jsPath)
	})
	cfgCheckpointBuffLock.Lock()
	defer cfgCheckpointBuffLock.Unlock()
	return *cfgCheckpointBuffModel
}
func (model CfgCheckpointBuff) loadConfig(jsPath string) {
	var fullPath = path.Join(jsPath, path.Base("./cfg_checkpoint_buff.json"))
	fileDataByte, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileDataByte, &cfgCheckpointBuffModel)
	if err != nil {
		panic(err)
	}
	if strings.HasPrefix(fullPath, "./") {
		pwd, _ := os.Getwd()
		fullPath = pwd + strings.TrimPrefix(fullPath, ".")
	}
	var v = viper.New()
	v.SetConfigFile(fullPath)
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		fileDataByte, err = ioutil.ReadFile(fullPath)
		if err != nil {
			log.Printf("errors!, config ioutil.ReadFile %s update faild", e.Name)
		}
		cfgCheckpointBuffLock.Lock()
		defer cfgCheckpointBuffLock.Unlock()
		cfgCheckpointBuffModel = new(CfgCheckpointBuff)
		err = json.Unmarshal(fileDataByte, &cfgCheckpointBuffModel)
		if err != nil {
			log.Printf("errors!, config json.Unmarshal %s update faild", e.Name)
		}
	})
	v.WatchConfig()
}
