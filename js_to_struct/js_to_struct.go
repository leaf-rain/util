package js_to_struct

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/leaf-rain/util/js_to_struct/listener"
	"github.com/leaf-rain/util/js_to_struct/parser"
	"github.com/leaf-rain/util/js_to_struct/util"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func JsonToStructForFolder(folderPath, outPath, pkgName string) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}
	if strings.HasPrefix(outPath, "./") {
		pwd, _ := os.Getwd()
		outPath = pwd + strings.TrimPrefix(outPath, ".")
	}
	for _, item := range files {
		fullname := folderPath + item.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if item.IsDir() {
			err = JsonToStructForFolder(fullname, outPath, pkgName)
			if err != nil {
				return err
			}
		} else if path.Ext(fullname) == ".json" {
			jts := NewJsonToStruct(fullname, outPath, pkgName, "")
			err = jts.ToStruct()
			if err != nil {
				fmt.Println("err:", err)
			}
		}
	}
	return nil
}

type JsonToStruct struct {
	ConfPath   string // 配置路径
	OutPath    string // 输出路径
	PkgName    string // 输出包名
	StructName string // 输出结构体名称
}

func NewJsonToStruct(confPath, outPath, pkgName, StructName string) *JsonToStruct {
	if confPath == "" || pkgName == "" {
		return nil
	}
	var result = &JsonToStruct{
		ConfPath:   confPath,
		OutPath:    outPath,
		PkgName:    pkgName,
		StructName: StructName,
	}
	if outPath == "" || path.Ext(outPath) == "" {
		result.AutoOutPath()
	}
	if StructName == "" {
		result.AutoStructName()
	}
	return result
}

func (jts *JsonToStruct) AutoOutPath() {
	var fileDir string // 如果没有指定输出文件路径，则默认输出在配置文件中
	if jts.OutPath != "" {
		fileDir, _ = path.Split(jts.OutPath)
	} else {
		fileDir, _ = path.Split(jts.ConfPath)
	}
	jts.OutPath = fileDir + strings.TrimSuffix(path.Base(jts.ConfPath), path.Ext(jts.ConfPath)) + ".go"
}

// 如果不指定结构体名称的话自动生成结构体名称
func (jts *JsonToStruct) AutoStructName() {
	jts.StructName = util.CamelString(strings.TrimSuffix(path.Base(jts.ConfPath), path.Ext(jts.ConfPath)))
}

// 输出到指定目录
func (jts *JsonToStruct) ToStruct() error {
	if jts.OutPath == "" || path.Ext(jts.OutPath) == "" {
		jts.AutoOutPath()
	}
	if jts.StructName == "" {
		jts.AutoStructName()
	}
	fileDataByte, err := ioutil.ReadFile(jts.ConfPath)
	if err != nil {
		return err
	}
	fileDataStr := string(fileDataByte)
	fileDataStr = strings.Trim(fileDataStr, " ") // 去除首尾空格
	inputSteam := antlr.NewInputStream(fileDataStr)
	lex := parser.NewJSONLexer(inputSteam)
	stream := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	fileParser := parser.NewJSONParser(stream)
	fileListener := listener.NewJsonToGoListener(&listener.GoTarget{
		ConfPath: jts.ConfPath,
		PkgName:  jts.PkgName,
		ConfName: jts.StructName,
	})
	antlr.ParseTreeWalkerDefault.Walk(fileListener, fileParser.Json())
	err = util.PutGoLang(jts.OutPath, fileListener.JsonStr)
	if err != nil {
		return err
	}
	return nil
}
