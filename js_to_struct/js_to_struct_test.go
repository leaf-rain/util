package js_to_struct

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
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

func TestJsonToStructForFolder(t *testing.T) {
	err := JsonToStructForFolder("./", "./output/", "js_to_struct")
	if err != nil {
		t.Errorf("failed, err:%v", err)
	} else {
		t.Logf("success")
	}
}

func TestTest(t *testing.T) {
	var _, err = os.Stat("/Users/dartou")
	t.Log(os.IsNotExist(err))
	t.Log(path.Base("/Users/dartou/test.json"))

}

func BenchmarkSprint2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var o = Online{
			Uid:    1,
			Sid:    2,
			GameId: 3,
			Gtype:  4,
			Isseat: "5",
			Ip:     "6",
		}
		i, _ := json.Marshal(o)
		_ = json.Unmarshal(i, &o)
	}
}

func BenchmarkSprint1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var o = Online{
			Uid:    1,
			Sid:    2,
			GameId: 3,
			Gtype:  4,
			Isseat: "5",
			Ip:     "6",
		}
		str := o.ToString()
		o = StringToOnline(str)
	}
}

type Online struct {
	Uid    uint32
	Sid    uint32
	GameId uint32
	Gtype  int32
	Isseat string
	Ip     string
}

func (o Online) ToString() string {
	return fmt.Sprintf("%d;%d;%d;%d;%s;%s", o.Uid, o.Sid, o.GameId, o.Gtype, o.Isseat, o.Ip)
}

func StringToOnline(str string) Online {
	slice := strings.Split(str, ";")
	var uid, _ = strconv.ParseUint(slice[0], 10, 64)
	var sid, _ = strconv.ParseUint(slice[1], 10, 64)
	var gameId, _ = strconv.ParseUint(slice[2], 10, 64)
	var gtype, _ = strconv.ParseInt(slice[3], 10, 64)
	return Online{
		Uid:    uint32(uid),
		Sid:    uint32(sid),
		GameId: uint32(gameId),
		Gtype:  int32(gtype),
		Isseat: slice[4],
		Ip:     slice[5],
	}
}
