package js_to_struct

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func init() {
	cfgCheckpointBuffModel.loadConfig()
}

var cfgCheckpointBuffModel = new(CfgCheckpointBuff)
var cfgCheckpointBuffLock = new(sync.RWMutex)

type CfgCheckpointBuff []CfgCheckpointBuffInfo

type CfgCheckpointBuffInfo struct {
	ScientificId         int64    `json:"scientificId"`
	SkillsId             int64    `json:"skillsId"`
	MaxLv                int64    `json:"maxLv"`
	PreconditionSkills   []int64  `json:"preconditionSkills"`
	PreconditionMaterial []string `json:"preconditionMaterial"`
	PreconditionBuild    []int64  `json:"preconditionBuild"`
}

func GetCfgCheckpointBuff() CfgCheckpointBuff {
	cfgCheckpointBuffLock.Lock()
	defer cfgCheckpointBuffLock.Unlock()
	return *cfgCheckpointBuffModel
}
func (model CfgCheckpointBuff) loadConfig() {
	var fullPath = "./cfg_checkpoint_buff.json"
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
		err = json.Unmarshal(fileDataByte, &cfgCheckpointBuffModel)
		if err != nil {
			log.Printf("errors!, config json.Unmarshal %s update faild", e.Name)
		}
	})
	v.WatchConfig()
}
