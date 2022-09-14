package listener

import (
	"fmt"
)

type GoTarget struct {
	ConfPath string
	PkgName  string
	ConfName string
}

func headFunc(pkgName string) string {
	return fmt.Sprintf(`package  %s 
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
)`, pkgName)
}

func initFunc(modelName string) string {
	return fmt.Sprintf("\nfunc init() {\n\t   %s.loadConfig() \n }", modelName)
}

func varFunc(modelName, confName, lockName string) string {
	return fmt.Sprintf("\n var %s = new(%s) \nvar %s = new(sync.RWMutex)\n", modelName, confName, lockName)
}

func modelFunc(confName, valStr string) string {
	return "\ntype " + confName + valStr + "\n"
}

func getFunc(modelName, lockName, confName string) string {
	return fmt.Sprintf(`
func Get%s() %s {
	%s.Lock()
	defer %s.Unlock()
	return *%s
}`, confName, confName, lockName, lockName, modelName)
}

func loadFunc(modelName, lockName, confName, confPath string) string {
	return fmt.Sprintf(`
func (model %s) loadConfig() { 
	var fullPath = "%s"
	fileDataByte, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileDataByte, &%s)
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
			log.Printf("errors!, config ioutil.ReadFile %%s update faild", e.Name)
		}
		%s.Lock()
		defer %s.Unlock()
		err = json.Unmarshal(fileDataByte, &%s)
		if err != nil {
			log.Printf("errors!, config json.Unmarshal %%s update faild", e.Name)
		}
	})
	v.WatchConfig()
}`, confName, confPath, modelName, lockName, lockName, modelName)
}

// func (t*GoTarget)ExitJson is called when production json is exited.
func (t *GoTarget) ExitJson(typeStr, valStr string) string {
	var modelName = FirstLower(suffix(t.ConfName))
	var lockName = FirstLower(suffixLock(t.ConfName))
	//return "package Generated \n type AutoGenerated " + valStr + "\n"
	return headFunc(t.PkgName) + initFunc(modelName) + varFunc(modelName, t.ConfName, lockName) +
		modelFunc(t.ConfName, valStr) + getFunc(modelName, lockName, t.ConfName) + loadFunc(modelName, lockName, t.ConfName, t.ConfPath)
}

// func (t*GoTarget)ExitObj is called when production obj is exited.
func (t *GoTarget) PreExitObj(typeStr, valStr string) string {
	return " struct {\n"
}
func (t *GoTarget) ExitObj(typeStr, valStr string, isEnd bool, bIsMap bool) string {
	return valStr + "\n"
}
func (t *GoTarget) PostExitObj(typeStr, valStr string) string {
	//println("##obj",valStr)
	return " \tlock     *sync.RWMutex \n}"
}

// func (t*GoTarget)ExitPair is called when production pair is exited.
func (t *GoTarget) ExitPair(index int, keyStr, typeStr, valStr, valType string) (string, string) {
	/*if typeStr=="struct"{
		return captical(stripQuotes(keyStr)) +valStr
	}
	if typeStr=="array"{
		return captical(stripQuotes(keyStr)) +valStr
	}

	*/
	return "", captical(stripQuotes(keyStr)) + " " + valStr + " `json:\"" + stripQuotes(keyStr) + "\"`" //+typeStr
}

// func (t*GoTarget)ExitArr is called when production arr is exited.
func (t *GoTarget) ExitArr(typeStr, valStr string) string {
	if valStr == "" {
		return "[] interface{}"
	}
	//println("##arr",valStr)
	return "[]" + valStr
}

// func (t*GoTarget)ExitValue is called when production value is exited.
func (t *GoTarget) ExitValue(typeStr, valStr string) string {
	if typeStr == "null" {
		return "interface{}"
	}
	return typeStr
}
