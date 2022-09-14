package js_to_struct

import (
	"fmt"
	js_to_struct "github.com/leaf-rain/util/js_to_struct/output"
	"os"
	"path"
	"testing"
	"time"
)

func TestNewJsonToStruct(t *testing.T) {
	jts := NewJsonToStruct("./cfg_checkpoint_buff.json", "", "js_to_struct", "")
	err := jts.ToStruct()
	if err != nil {
		t.Errorf("failed, err:%v", err)
	} else {
		t.Logf("success")
	}
}

func TestJsonToStructForFolder(t *testing.T) {
	err := JsonToStructForFolder("./", "./output/", "js_to_struct")
	if err != nil {
		t.Errorf("failed, err:%v", err)
	} else {
		t.Logf("success")
	}
	for {
		time.Sleep(time.Second * 2)
		fmt.Println(js_to_struct.CfgCheckpointBuffModel)
	}
}

func TestTest(t *testing.T) {
	var _, err = os.Stat("/Users/dartou")
	t.Log(os.IsNotExist(err))
	t.Log(path.Base("/Users/dartou/test.json"))

}
