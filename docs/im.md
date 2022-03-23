# 即时通信&音视频相关接口

## host

内网: `http://192.168.1.111:8081/`
外网测试: `http://im-demo-finance.tapmechat.com/`

## 匹配，融云视屏相关接口，content中的extra字段

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配编号|
|round_no|int|轮次|
|scene|int|场景：1：速配；0：1v1|

## 接口

### im消息发送前置校验

**/api/v1/im/before/check**

参数:

|字段|值|
|---|---|
|rc_msg|RC:TxtMsg|
|peer_uid|对方的uid|
|content|文本内容，用于校验是否包含敏感词。|

```json
param: {
    "rc_msg": "RC:TxtMsg",    // 文本类型
    "peer_uid": 12,
    "content": "文本内容，用于校验是否包含敏感词。"
}
```

返回:

code:

|错误码|值|
|---|---|
|1003|余额不足|
|100002|包含敏感词|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 视频拨打事前校验

**/api/v1/video/before/check**

参数:

```json
param: {}
```

返回:

code:

|错误码|值|
|---|---|
|1003|余额不足|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 匹配发起

**/api/v1/video/pair/ask**

参数:

null

返回:

data:

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配标号

code:

|错误码|值|
|---|---|
|1003|余额不足|
|0|成功

```json
{
  "code": 0,
  "data": {
    "ask_no": "313131"
  },
  "err": null,
  "time": 1641473705
}
```

### 匹配应答

**/api/v1/video/pair/answer**

参数:

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配编号|
|round_no|int|匹配轮次|
|pass|bool|true:同意；false:拒绝|

返回:

data:

null

code:

|错误码|值|
|---|---|
|1003|余额不足|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 匹配发起方确认

**/api/v1/video/pair/confirm**

参数:

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配编号|
|pass|bool|true:喜欢；false:不喜欢|
|round_no|int|轮次|
|peer_uid|int|对方uid|
|call_id|string|通话call_id|

返回:

data:

null

code:

|错误码|值|
|---|---|
|1003|余额不足|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 匹配发起方取消

**/api/v1/video/pair/cancel**

参数:

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配编号|

返回:

data:

null

code:

|错误码|值|
|---|---|
|1003|余额不足|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 匹配状态查询

**/api/v1/video/pair/info**

参数:

|字段|类型|说明|
|---|---|---|
|ask_no|string|匹配编号|

返回:

data:

|字段|类型|说明|
|---|---|---|
|status|int8|0:匹配中；1：结束；2：通话中|
|ask_no|string|匹配编号|
|ask_uid|int|匹配发起方uid|
|call_id|string|通话call_id|
|peer_uid|int|对方的uid|

null

code:

|错误码|值|
|---|---|
|1003|余额不足|
|200003|获取配置信息失败|
|200004|无配置信息|
|0|成功

```json
{
  "code": 0,
  "data": null,
  "err": null,
  "time": 1641473705
}
```

### 匹配长链接下发消息

**obJectName:"RC:CmdMsg"**

content结构:

|字段|类型|说明|
|---|---|---|
|name|string|事件；dispatch：派单;|
|data|object||
|data.ask_no|string|匹配编号|
|data.ask_uid|int|匹配发起方uid|
|data.round_no|int|匹配轮次|

## 翻译

** /api/v1/translate/action

参数:

```json
// text : 待翻译的文字
// short : 翻译的语言简称
param:{
"text":
["hello", "world"], "short": "zh"
} 
```

```json
{
  "code": 0,
  "data": [
    "你好",
    "世界"
  ],
  "err": null,
  "time": 1644562855
}
```

### 发送小灰条 tip 消息

**/api/server/im/send/tip**
参数:

```json
param:{
"from": 100, "targets":
[], "msg": "this is a test msg."
} 
```

|字段|类型|说明|
|---|---|---|
|from|string|发送者 uid|
|targets|[]uint64|接收目标的 uid 数组|
|msg|string|消息内容|

### 发送系统消息

**/api/server/im/send/system**
参数:

```json
param:{
"from": 100, "targets":
[], "msg": "this is a test msg."
} 
```

|字段|类型|说明|
|---|---|---|
|from|string|发送者 uid|
|targets|[]uint64|接收目标的 uid 数组|
|msg|string|消息内容|

### 客户端删除匹配记录

/api/v1/video/delPairRecord 参数：

```json
param: {
"ask_no": "",
"start_at": 1 // 这里返回查询到的start时间
}
```