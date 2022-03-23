# 1. 接口目录
| 服务 | 接口文档 | 内网地址 | 线上正式服地址 |
|---|---|---|---|
| user    | [user.md](user.md) | http://192.168.1.111:8080/ | http://api-v1.touchat.live | 
| cms     | [cms.md](cms.md) | http://192.168.1.111:8084/ | http://api-v1.touchat.live | 
| finance | [finance.md](finance.md) | http://192.168.1.111:8082/ | http://api-v1.touchat.live |
| imcenter | [imcenter.md](imcenter.md) | http://192.168.1.111:8081/ | http://api-v1.touchat.live |
| gift    | [user.md](gift.md) | http://192.168.1.111:8083/ | http://api-v1.touchat.live |
| ecode    | [ecode.md](ecode.md) |||

# 2. 接口请求格式

接口请求 URL 格式:
`scheme://host/api/v1/$apipath`

接口请求使用JSON格式， `request` 统一使用 header: `"Content-Type": "application/json`, `JSON` 固定格式如下：

### 参数：
|参数名 |类型| 解释 |
|--|--|--|
|time| number|请求时间 |
|ver_app |string |版本(包) |
|ver_res|string |版本(资源) |
|mtkey|string |签名key |
|appid|number |渠道id |
|uid |number|uid |
|sesskey|string |session |
|lang|string |语言 |
|pack |number|包类别， 1: 男包, 2: 女包  0: 不区分 |
|sig|string |签名校验 |
|param |object|参数 |
### e.g:
```json
{              
  "time":1630995979,             
  "ver_app":"1.0.4.1",         
  "ver_res":"1.2.4",            
  "mtkey": "mtkey",            
  "appid":200,                
  "uid":9123690,               
  "sesskey": "sesskey",           
  "lang":"EN",                
  "param": {},               
  "pack": 1,                   
  "sig":"sig"              
}

```

数据返回格式：
|参数名 |类型| 解释 |
|--|--|--|
|code|number|业务码（为0时请求正常）|
|data|object|响应体|
|err|string|错误信息|
|time|number|响应时间（时间戳）|

```json
{
    "code":0,
    "data":{},
    "err":null,
    "time":1640762242
}
```

# 3.SIGN加密算法

所有post 数据除 `sig` 字段外，需要参与sign加密。

将`json` 的所有 key 取出，按照ascii 从小到大排序， 将KEY / Value 使用 `=` 连接， value使用 mtkey加盐，最后将所有字符串连接并`md5` ，得到的值为 sig 对应的值。

假如请求的 参数如下:

```json
var body = {              
  "time":1630995979,             
  "ver_app":"1.0.4.1",           
  "ver_res":"1.2.4",             
  "mtkey": "MTKEY",            
  "appid":200,                   
  "uid":9123690,                
  "sesskey": "e024d6febaddef74fb3a87c1d66c4edf",           
  "lang":"EN",                  
  "param": { "sys_version":"1.0.10", "coins": 100000, "isMale": true }, 
  "sig":"",
}
```

```js
const headflag = "MH"
function genSign() {
  var mtkey = body.mtkey;
  
  // sig 和 mtkey字段不进行组装
  delete body.sig;
  delete body.mtkey;
  
  var str = getSignString(body, mtkey);
  return md5(str);
}

function getSignString(obj, mtkey) {
  var ret = "";
  if (obj == null || typeof obj == "boolean") {
    return ret;
  }
  if (typeof obj == "string" || typeof obj == "number") {
    return headflag + mtkey + obj;
  }
  if (typeof obj == "object") {
    var keys = Object.keys(obj);
    keys = keys.sort();
    for (let i = 0; i < keys.length; i++) {
      var k = keys[i];
      str += k + "=" + getSignString(obj[k], mtkey)
    }
  }
  return ret
}
```

# 4. 示例:登录接口

**参数**

```json
{
    "appid": 10001,
    "lang": "zh",
    "mtkey": "MTKEY",
    "pack": 1,
    "param": {
        "platform": "visitor",
        "siteuid": "123456789",
        "devname": "android",
        "sys_version": "1.0.0",
        "deviceid": "123456789"
    },
    "sesskey": "",
    "sig": "f75fc3064a2607e575b50b00ade14496",
    "time": 1646188653211,
    "uid": 88888888,
    "ver_app": "1.0.0",
    "ver_res": "1.0.0",
    "method": "3af7e5ef"
}
```

**返回**

```json
{
    "code": 0,
    "data": {
        "account": {
            "Appid": 0,
            "Coins": 3200,
            "Diamonds": 1500,
            "Free": 4,
            "MsgBalance": 4,
            "PairCards": 14,
            "PayAmount": 0,
            "PayCount": 4,
            "uid": 87207658
        },
        "guid": "123456789",
        "imToken": "GjxJ1zuevVuXeMhX61q6djUhZAT2gVMqDVBoXkxavDs=@1apu.sg.rongnav.com;1apu.sg.rongcfg.com",
        "info": {
            "Age": 25,
            "Appid": 10001,
            "CountryCode": "",
            "GeetingPic": "",
            "GeetingWords": "",
            "Gender": 0,
            "HeadPic": "",
            "ImToken": "GjxJ1zuevVuXeMhX61q6djUhZAT2gVMqDVBoXkxavDs=@1apu.sg.rongnav.com;1apu.sg.rongcfg.com",
            "Intro": "",
            "IsOnline": false,
            "IsVerifyPkg": false,
            "Lang": "zh",
            "Level": 0,
            "LoginDevid": "123456789",
            "LoginDevname": "android",
            "LoginIP": "",
            "LoginIdfa": "123456789",
            "LoginPlace": "",
            "LoginSid": 0,
            "LoginTime": 1647914231,
            "MatchOnOff": false,
            "RegAppid": 10001,
            "RegDevid": "123456789",
            "RegIP": "",
            "RegIdfa": "",
            "RegPlace": "",
            "RegTime": 1647328124,
            "SecondLang": "",
            "Status": 0,
            "SysLang": "",
            "SysTimezone": "",
            "SysVersion": "1.0.0",
            "VerApp": "1.0.0",
            "VerRes": "1.0.0",
            "borndate": 858409724,
            "nickname": "USER_87207658",
            "uid": 87207658
        },
        "session": "300208ccae302843eb24ac9965faa7d6",
        "uid": 87207658
    },
    "err": "",
    "time": 0
}
```
