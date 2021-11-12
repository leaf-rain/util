package tool

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	var data = make(map[string]interface{})
	data["id_card"] = 100045
	data["cid"] = 3
	//data["hash"] = "0x1a5e0415eb1754776b90a8662778e009e52d81d1351cbc0c9e5b8492dc5d9c76"
	//data["net_work"] = "frt"
	data["phone"] = "13609245992"
	//data["contract"] = "0xbf89666bca17bb9c58e65d8e71dd3ae1adbab91e"
	//data["net_work"] = "frt"
	//data["announcer"] = "0x7f009de9f7687ba4458224b461eff75e683d4ed0"
	//data["artwork_name"] = "12312312"
	//data["artwork_name"] = "12312312"
	//data["material"] = "12312312"
	//data["spec"] = "12312312"
	//data["artwork_year"] = "12312312"
	//data["introduce"] = "12312312"
	//data["image_url"] = "12312312"

	//data["uid"] = 254911
	//data["cid"] = 2
	//data["power"] = 1000
	//data["remark"] = "添加用户备注1"

	//data["uid"] = 2549
	//data["cid"] = 2

	//data["timestamp"] = fmt.Sprint(GetTimeUnixMilli())
	//data["timestamp"] = fmt.Sprint(GetTimeUnixMilli())
	data["timestamp"] = "1634031449792"
	signStart := JoinStringsInASCII(data, "&", false, "sign")
	signEnd := fmt.Sprintf("%s&sign=%s", signStart, "123456")
	md5Key := GetMD5Encode(signEnd)
	data["sign"] = md5Key
	s, _ := json.Marshal(data)
	fmt.Println(string(s))
	fmt.Println(Sign(data, "123456"))
}
