package js_to_struct

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/leaf-rain/util/js_to_struct/listener"
	"github.com/leaf-rain/util/js_to_struct/parser"
	"io/ioutil"
	"os"
	"os/exec"
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
	jts.StructName = CamelString(strings.TrimSuffix(path.Base(jts.ConfPath), path.Ext(jts.ConfPath)))
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
	err = PutGoLang(jts.OutPath, fileListener.JsonStr)
	if err != nil {
		return err
	}
	return nil
}

// 蛇形转驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func PutGoLang(name, content string) error {
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			if _, err = os.Stat(path.Dir(name)); err != nil {
				if os.IsNotExist(err) {
					_ = os.MkdirAll(path.Dir(name), 0755)
				}
			}
			f, err = os.Create(name)
		}

	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	cmd := exec.Command("gofmt", "-w", f.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}
