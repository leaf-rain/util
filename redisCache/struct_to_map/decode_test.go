package struct_to_map

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_decodeStruct(t *testing.T) {
	var testStruct = tts{}
	var f = reflect.ValueOf(&testStruct.Values)
	f = f.Elem()
	var js, _ = json.Marshal(&ts{
		Name: "张三",
		Age:  100,
	})
	err := decodeJson(f, string(js))
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!! ts:%+v", testStruct.Values)
	}
}
