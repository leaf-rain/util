package struct_to_map

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var testStruct = Request{}
	testStruct.Ping = "ping"
	result := Encode(&testStruct)
	if result == nil {
		t.Errorf("failed, err:%+v", result)
	} else {
		t.Logf("success !!! result:%+v", result)
	}
}
