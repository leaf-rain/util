# 用户相关接口

## host

内网: `http://192.168.1.111:8080/`
外网测试: `http://api-demo.tapmechat.com/`

## 接口

### 登录

见: [README.md](README.md)

### 用户资料信息

** /api/v1/user/getUserInfo

参数:

```json
param: { "uid": 123 }
```

返回：

```json
{"code":0,"data":{"uid":1,"Appid":0,"nickname":"USER_1","borndate":0,"HeadPic":"","Gender":0,"CountryCode":"","Lang":"","RegTime":1640841882,"LoginTime":1640941172,"RegIP":"127.0.0.1 ","LoginIP":"127.0.0.1","ImToken":"G6fZum4Pk8BSc9zU88tKqzGk1MVUtX+v@1apu.sg.rongnav.com;1apu.sg.rongcfg.com","Status":0,"RegDevid":"","RegIdfa":"","RegPlace":"","RegAppid":0,"LoginDevid":"","LoginIdfa":"","LoginPlace":"","LoginDevname":"","VerApp":"","VerRes":"","SysVersion":"","SysTimezone":"0","SysLang":""},"err":null,"time":1641475508}
```

### 验证Session

** /api/v1/user/checkSession

参数:

```json
param: {}
```

返回：

```json
{
    "code":0,
    "data":{
        "info": userinfo, // 格式见登录接口
        "account": useraccount, // 格式见登录接口
        "imToken":"xxxxxxxxxxxxxxxxxxx"
    },
    "err":null,
    "time":1234567
}
```

### 在线用户列表

** /api/v1/user/onlineUserList

参数:

```json
// 性别和显示数量
{
      "gender": 1,
      "type":0, //0所用用户，1新用户，2充值用户
      "no": 1, // 页码
      "size":10, // 每页显示数据
      "time": 0, // 起始时间
      "cond": "[1,1,10,0]"  // 回传查询条件，默认为空
  }
```

返回：

```json
{
    "code":0,
    "data":{
        "online": [
            {
                "login": 1234567890, // 登录时间
                "heart": 1234567890, // 当前时间
                "time": 120,
                "gender":1,
                "headpic":"",
                "nickname":"test001",
                "country":"china",
                "lang":"zh",
            }
        ],
        "cond": "[1,1,10,0]"
        
    },
    "err":null,
    "time":1234567
}
```

### 心跳包

**/api/v1/user/heartbeat**

参数:

```json
param: {  }
```

返回：

```json
// LoginTime int64  `json:"login"`
// HeartTime int64  `json:"heart"`
// TotalTime int64  `json:"time"`
// Gender    int    `json:"gender"`
// HeadPic   string `json:"headpic"`
// Nickname  string `json:"nickname"`
// Country   string `json:"country"`
// Lang      string `json:"lang"`
{
    "code":0,
    "data":{
        "login": 1234567890, // 登录时间
        "heart": 1234567890, // 当前时间
        "time": 120,
        "gender":1,
        "headpic":"",
        "nickname":"test001",
        "country":"china",
        "lang":"zh/cn",
    },
    "err":null,
    "time":1234567
}

```

### 首次登录用户完善信息接口

** /api/v1/user/onPerfect

参数:

```json
//都是必传字段
param: {
    "head_pic":"http://",
    "nick_name":"nn", 
    "age":30,
    "region":"India",
    "intro":"hello",
    "album_param":[
        {
            "source_type":0,
            "url":"http://",
            "media_type":0
        },{
            "source_type":0,
            "url":"http://",
            "media_type":0
        },{
            "source_type":0,
            "url":"http://",
            "media_type":0
        },{
            "source_type":0,
            "url":"http://",
            "media_type":1
        }
    ]
}
```

```json
返回：
{
    "code": 0,
    "data": null,
    "err": null,
    "time": 1641894530
}
```

### 用户信息修改接口

**/api/v1/user/onUpdate**

参数:

```json
//如果只是单独修改单个字段，其他字段可不传
//增加 geeting_words,geeting_pic 上传修改
//增加 birt_date
//增加第二语言
param: {
    "birt_date":"1990-10-10",
    "geeting_words":"hello",
    "geeting_pic":"http://",
    "head_pic":"http://",
    "nick_name":"nn", 
    "age":30,
    "region":"India",
    "language":"English",
    "intro":"hello",
    "nd_lang":"India"
```

```json
返回：
{
    "code": 0,
    "data": null,
    "err": null,
    "time": 1641894530
}
```

### 用户相册获取

** /api/v1/user/getUserAlbum

参数:

```json
param: {"uid":20}
```

返回: 

```json
{
    "code": 0,
    "data": [
        {
            "aid": 7,
            "uid": 9123690,
            "source_type": 1,
            "url": "http://",
            "media_type": 1,
            "status": 0,  //0：未审核，1审核失败，2审核通过
        },
        {
            "aid": 8,
            "uid": 9123690,
            "source_type": 1,
            "url": "http://",
            "media_type": 1,
            "status": 0,
        }
    ],
    "err": null,
    "time": 1641896884
}

```

### 相册更新

** /api/v1/user/updateAlbums

参数:

```js
//source_type图片来源 ; media_type相册类型 0:pic,1:video;
//图片视频分别最大6张
param:  
    {
    "upload":[{"source_type":1,"url":"http://","cover":"http://", "media_type":0}],
    "delete":[11,12]
    }
```

```js
{
    "code": 0,
    "data": null,
    "err": null,
    "time": 1641956943
}
```




### 随机获取一个相册视频 

**/api/v1/user/randomVideo**

参数:

```json
param: {}
```

```json
{
    "code": 0,
    "data": {
        "url": "http://"
    },
    "err": null,
    "time": 1646033731
}
```


### 设置繁忙状态或取消繁忙

**/api/server/user/setBusyStatus**

```json
param: { uid: xxx, is_busy: true }
```

返回:

```json
{
    "code": 0,
    "data": [],
    "err":null
}
```

### 获取配置信息

**/api/v1/config/getSystemConfig**

```json
param: { }
```

返回:

```json
{
    "code": 0,
    "data": {
        "ApiHosts": {
            "ApiUser": "user-demo.tapmechat.com",
            "ApiFinance": "finance-demo.tapmechat.com",
            "ApiIM": "im-demo.tapmechat.com",
            "TcpURL": "",
            "TcpURL2": "",
            "CdnURL": "img.tapmechat.com"
        },
        "Switchs": {
            "FBLogin": "0",
            "GoogleLogin": "true",
            "IsPreview": "true",
            "VisitorLogin": "true"
        },
        "Version": "1.2.3.4",
        "Download": "",
        "FbLink": "",
        "PackageName": "com.tapmechat.chat",
        "ClientType": 2,
        "CoinsConsume": {
            "Text": 1,
            "Voice": 1,
            "Img": 1,
            "Video": 30,
            "Pair": 8,
            "FistSeconds": 15
        }
    },
    "err": null,
    "time": 1645245717
}
```

### 随机在线用户列表

**/api/server/user/onlineRandomUser**

参数:

```json
// 性别和显示数量
param: { "gender":1, "size":20 }
```

返回:

```json
{
    "code": 0,
    "data": [
        21,
        456
    ],
    "err": null,
    "time": 1642404720
}
```

### 用户喜欢

**/api/v1/user/likeTo**

参数:

```json
param: {"like_uid":18}
```

返回:

```json
{
    "code": 0,
    "data": {"uid": 0, "star": 0},
    "err": null,
    "time": 1642404720
}
```

### 用户取消喜欢
**/api/v1/user/likeCancel**

参数:

```json
param: {"like_uid":18}
```

返回:

```json
{
    "code": 0,
    "data": {"uid": 0, "star": 0},
    "err": null,
    "time": 1642404720
}
```

### 获取用户喜欢的列表
**/api/v1/user/getMyLike**
参数:

```json
param: {"page_size":10,"page_no":1}
```
返回:

```json
{
    "code": 0,
    "data": [
        {
            "uid": 20,
            "nick_name": "nn",
            "head_pic": "http://",
            "gender": 0,
            "age": 33,
            "country_code": "India",
            "lang": "English",
            "status": 1,
            "star": 1 // 用户喜欢状态，1：我喜欢的，2：喜欢我的，3：互相喜欢的
        }
    ],
    "err": null,
    "time": 1642411540
}
```

### 获取喜欢用户的列表
**/api/v1/user/getLikeMe**
参数:

```json
param: {"page_size":10,"page_no":1}
```

返回:

```json
{
    "code": 0,
    "data": [
        {
            "uid": 20,
            "nick_name": "nn",
            "head_pic": "http://",
            "gender": 0,
            "age": 33,
            "country_code": "India",
            "lang": "English",
            "status": 1,
            "star": 1 // 用户喜欢状态，1：我喜欢的，2：喜欢我的，3：互相喜欢的
        }
    ],
    "err": null,
    "time": 1642411540
}
```
### 查询是否喜欢的状态
**/api/v1/user/checkUserLike**
参数:

```json
param: {"uid":19}
```

返回:

```json
{
    "code": 0,
    "data": {
        "like": true
    },
    "err": null,
    "time": 1642646593
}
```

### 获取关注自己用户数量
**/api/v1/user/getLikeMeCount**
参数:

```json
param: {}
```

返回:

```json
{
    "code": 0,
    "data": {
        "count": 1
    },
    "err": null,
    "time": 1642646593
}
```

### 获取我关注用户数量
**/api/v1/user/getMyLikeCount**
参数:

```json
param: {}
```

返回:

```json
{
    "code": 0,
    "data": {
        "count": 1
    },
    "err": null,
    "time": 1642646593
}
```

### 获取用户喜欢信息
**/api/v1/user/getLikeByUser**
参数:

```json
param: {"uid"：88888888}
```

返回:

```json
{
    "code": 0,
    "data": {
        "follow": 123,
        "following": 123,
        "is_like": 1
    },
    "err": null,
    "time": 1642646593
}
```
is_like: // 用户喜欢状态，1：我喜欢的，2：喜欢我的，3：互相喜欢的


### 分页拿在线用户UID

**/api/server/user/getOnlineUids**

参数:

```json
// 性别和显示数量
{
      "gender": 1,
      "no": 1, // 页码
      "size":10, // 每页显示数据
      "time": 0, // 起始时间
      "cond": "[1,1,10,0]"  // 回传查询条件，默认为空
  }
```

返回:

```json
{
    "code": 0,
    "data": {
        "cond": "[1,1,10,0]",
        "uids": [
            21,
            456
        ]
    },
    "err": null,
    "time": 1642472645
}
```



## 通话记录

**/api/v1/user/getCallInfo**

参数:

```json
param: {"page_size":10,"page_no":1}
```

```json
{
    "code": 0,
    "data": [
        {
            "uid": 42,
            "nick_name": "USER_42",
            "head_pic": "",
            "gender": 1,
            "age": 0,
            "country_code": "",
            "lang": "EN",
            "status": 0
        },
        {
            "uid": 42,
            "nick_name": "USER_42",
            "head_pic": "",
            "gender": 1,
            "age": 0,
            "country_code": "",
            "lang": "EN",
            "status": 0
        }
    ],
    "err": null,
    "time": 1643022539
}
```

## 获取语言列表

**/api/v1/user/getLangConfig**

参数:

```json
param: {}
```

```json
{
    "code": 0,
    "data": {
        "af South African": "af",  //全称:简称
        "af-ZA South African": "af-ZA",
        "ar Arabic": "ar",
        "ar-AE Arabic (UAE)": "ar-AE",
        "ar-BH Arabic (Bahrain)": "ar-BH",
        "ar-DZ Arabic (Algeria)": "ar-DZ",
        "ar-EG Arabic (Egypt)": "ar-EG",
        "ar-IQ Arabic (Iraq)": "ar-IQ",
        "ar-JO Arabic (JORDAN)": "ar-JO",
    },
    "err": null,
    "time": 1643197724
}
```


## 账号绑定接口

**/api/v1/user/bind**

参数:

```json
param: {
  platform: platform,m             // 绑定平台 / facebook / google
  fb_appid: fb_appid,              //facebook appid
  fb_openid: openid,               //facebook openid
  fb_access_token: access_token,   //facebook token
  google_openid: google_openid,
  google_access_token: google_access_token,
}
```

```json
{
    "code": 0,
    "data":null,
    "err": null,
    "time": 1643197724
}
```

## 用户反馈接口

**/api/v1/user/feedback**

参数:

```json
  param: { "content": "hello!!"}
```

```json
{
    "code": 0,
    "data":null,
    "err": null,
    "time": 1643197724
}
```

## 获取在线用户数量

**/api/server/user/getOnlineUserCount**

参数:

```json
  param: {}
```

```json
{
    "code": 0,
    "data": {
        "famale": 0,
        "male": 1
    },
    "err": null,
    "time": 1645427768
}
```

## 增加用户短语

**/api/v1/user/addShortcut**

参数:

```json
  param: {"short":"hello", "type": 1}
```
type:  类型，1：文本，2：语音

```json
{
    "code": 0,
    "data": null,
    "err": null,
    "time": 1645512176
}
```
## 删除用户短语

**/api/v1/user/delShortcut**

参数:

```json
  param: {"id":1}
```

```json
{
    "code": 0,
    "data":null,
    "err": null,
    "time": 1645512176
}
```

## 获取用户短语列表

**/api/v1/user/shortcutList**

参数:

```json
  param: {}
```

```json
{
    "code": 0,
    "data": {
        "text": [{"id":1, "uid":88888888, "short":"内容", "type":1, "status":1}],
        "voice": [{"id":1, "uid":88888888, "short":"内容", "type":2, "status":1}],
    },
    "err": null,
    "time": 1645512176
}
```
type:  类型，1：文本，2：语音
status: 审核状态，1：审核中，2：审核成功（没有审核失败，审核失败直接删除）

## 获取用户随机短语

**/api/v1/user/randomShortcut**

参数:

```json
  param: {}
```

```json
{
    "code": 0,
    "data": {
        "id":1, "uid":88888888, "short":"内容", "type":2, "status":1
    },
    "err": null,
    "time": 1645512176
}
```
type:  类型，1：文本，2：语音
status: 审核状态，1：审核中，2：审核成功（没有审核失败，审核失败直接删除）

## 获取用户匹配记录

__/api/v1/user/getPairRoundRecordByPage__

参数：
```json
param:{page_on: 1, page_size: 1}
```
响应：
```json
{
    "code": 0,
    "data": [
      {
        "uid": 42,
        "nick_name": "USER_42",
        "head_pic": "", // 头像
        "country_code": "", 
        "picture": "", //随机相册图片
        "duration": 1,
        "record": { // 记录信息
          "status": 1,// 通话状态，'0:默认状态；1：成功；2：失败；3：超时；',
          "call_id": 2, // 通话id
          "ask_no": "",
          "ask_uid": 1, // 请求id
          "peer_uid": 2, // 匹配id
          "start_at": 2012, // 开始时间 秒级时间戳
          "end_at": 2011 // 结束时间 秒级时间戳
          "create_at": 2022 // 创建时间 秒级时间戳
        }
      }
    ],
    "err": null,
    "time": 1645512176
}
```

## 根据uid获取用户的在线状态

**/api/v1/user/getUidsStatus**

参数:

```json
  param: {
    "uids":[
        {
            "uid":91182089,
            "gender":1
        }
    ]
  }
```
```json
 //0 不在线，1 ，在线，2繁忙
{
    "code": 0,
    "data": [
        {
            "uid": 91182089,
            "status": 0 
        }
    ],
    "err": null,
    "time": 1646203660
}
```

## 获取区域

**/api/v1/user/getCountryConfig**

参数:

```json
  param: {}
```
```json
{
    "code": 0,
    "data": {
        "country": {
            "ABU_DHABI": "https://img.tapmechat.com/country/abu_dhabi.png",
            "AFGHANISTAN": "https://img.tapmechat.com/country/afghanistan.png",
            "ALBANIA": "https://img.tapmechat.com/country/albania.png",
            "ALGERIA": "https://img.tapmechat.com/country/algeria.png",
            "AMERICAN_SAMOA": "https://img.tapmechat.com/country/american_samoa.png",
            "ANDORRA": "https://img.tapmechat.com/country/andorra.png",
            "ANGOLA": "https://img.tapmechat.com/country/angola.png",
            "ANGUILLA": "https://img.tapmechat.com/country/anguilla.png",
            "ANTIGUA_AND_BARBUDA": "https://img.tapmechat.com/country/antigua_and_barbuda.png",
        }
    },
    "err": null,
    "time": 1646209202
}
```




## 举报用户

**/api/v1/user/report**

参数:

```json
  param: {
    "report_uid": 0, // 举报用户id
    "type": 0, // 举报类型
    "location": 8, // 举报位置
    "content": "举报内容" // 最长长度500，上报具体内容（最好是json数据，方便解析）
}
```

```json
{
    "code": 0, 
    "data":null,
    "err": null,
    "time": 1643197724
}
```

## 添加黑名单

**/api/v1/user/addBlocklist**

参数:

```json
  param: {
    "block_uid": 0, // 拉黑用户id
}
```

```json
{
    "code": 0, 
    "data":null,
    "err": null,
    "time": 1643197724
}
```

## 数据上报

**/api/v1/user/inform**

参数:

```json
  param: {
    "uid": 0, // 上报uid(可空)
    "type": 0, // 上报类型，1软件闪退
    "content": "参数解析错误" // 最长长度500，上报具体内容（最好是json数据，方便解析）
}
```

```json
{
    "code": 0, 
    "data":null,
    "err": null,
    "time": 1643197724
}
```
## 随机几张在线用户图片

**/api/v1/user/randomOnlinePic**

参数:

```json
  param: {}
```

```json
{
    "code": 0, 
    "data": {
        "pic": [{
            "https://img.tapmechat.com/country/abu_dhabi.png",
        }], 
    "err": null,
    "time": 1643197724
}
```

## 随机获取问候语

参数:

```json
  param: {}
```

```json
{
    "code": 0, 
    "data": {
        "msg": "hello_word", 
    "err": null,
    "time": 1643197724
}
```

## 设置匹配开关(获取匹配开关状态再userinfo里面 MatchOnOff 字段)

**/api/v1/user/setMatchOnoff**

参数:

```json
  param: {"onoff":true}
```

```json
{
    "code": 0, 
    "data": {
        "onoff":true,
    }
    "err": null,
    "time": 1643197724
}
```

