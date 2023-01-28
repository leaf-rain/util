package struct_to_map

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var testStruct = tts{}
	testStruct.Values = &ts{
		Name: "李四",
		Age:  20,
	}
	testStruct.Map = make(map[int]interface{})
	testStruct.Map[1] = 2
	testStruct.Map[2] = 3
	err := Encode(testStruct)
	if err != nil {
		t.Errorf("failed, err:%+v", err)
	} else {
		t.Logf("success !!!")
	}
}
