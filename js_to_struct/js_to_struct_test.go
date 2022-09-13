package js_to_struct

import (
	"testing"
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
