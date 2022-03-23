# 金融相关接口

## 目录
|seq|ack|method|explain|check_session|
|-|-|-|-|-|
| ab607d40| |GetBalance|[查钱包余额](#查钱包余额)|true|
| 618514e7| |ActAccount|[操作钱包账户](#操作钱包账户)|true|
| 1e716564| |Income|[用户收入](#收入)|true|
| 30b2d215| |Consume|[消费](#消费)|true|
| | ||[操作钱包账户](#操作钱包账户)|true|
| | ||[操作钱包账户](#操作钱包账户)|true|
| | ||[操作钱包账户](#操作钱包账户)|true|
| | ||[操作钱包账户](#操作钱包账户)|true|
| | ||[操作钱包账户](#操作钱包账户)|true|

## 公共参数
#### balance
|参数名 |类型| 解释 |
|--|--|--|
|uid|number|用户id|
|Appid|number|渠道id |
|Coins|number|金币 |
|Diamonds|number|钻石 |
|Free|number|免费赠送 |
|PayAmount|float|历史充值金额 |
|PayCount|number|充值次数 |
|PairCards|number|匹配卡数量 |
|MsgBalance|number|剩余发消息次数 |
e.g:
```json
{
    "uid":1,
    "Appid":0,
    "Coins":1010,
    "Diamonds":0,
    "Free":0,
    "PayAmount":0,
    "PayCount": 0, 
    "PairCards": 0, 
    "MsgBalance": 0
}
```
#### fields
|参数名 |类型| 解释 |
|--|--|--|
|type|number|币种, 0:Coins,1:Diamonds,2:Free,3:PayCount,4:PayAmount,5:PairCards,7:MsgBalance|
|opt|number|操作, 0: 增加 1: 减少 2: 直接赋值|
|value|float|变动额|
e.g:
```json
[1,0,100]
```
#### logtype
|参数名 |类型| 解释 |
|--|--|--|
|logtype|number|1000:注册赠送,1001:登录奖励,1002:充值金币,1003:充值奖励,1004:文本,1005:语音,1006:礼物,1007:视频,|

### 查钱包余额
参数:
```json
param: {}
```
返回:
|参数名 |类型| 解释 |
|--|--|--|
|-|object|[balance](#balance)|
```json
{
    {{balance}}
}
```

### 操作钱包账户
|参数名 |类型| 解释 |
|--|--|--|
|logtype|number|[日志类型](#logtype)|
|fields|arry|[金额操作](#fields)|
|content|string|备注信息|

```json
param: { 
    "logtype":1000,
    "content":"系统赠送",
    "fields": [[0, 0, 100]]
}
```
返回:
|参数名 |类型| 解释 |
|--|--|--|
|-|object|[balance](#balance)|
```json
{
    {{balance}}
}
```

### 收入
参数:
|参数名|类型|解释|
|--|--|--|
|coins|number|金币|
|logtype|number|[日志类型](#logtype)|
|content|string|备注信息|
|buzid|string|||
```json
param: { "coins": 123, "logtype":10001, "content":"系统赠送" }
```
返回:
|参数名 |类型| 解释 |
|--|--|--|
|-|object|[balance](#balance)|
```json
{
    {{balance}}
}

```
### 消费
参数:
|参数名|类型|解释|
|--|--|--|
|coins|number|金币|
|logtype|number|[日志类型](#logtype)|
|content|string|备注信息|
```json
param: { "coins": 123, "logtype":10002, "content":"系统消耗" }
```

返回:
|参数名 |类型| 解释 |
|--|--|--|
|-|object|[balance](#balance)|
```json
{
    {{balance}}
}
```

### 转账接口
参数:
|参数名 |类型| 解释 |
|--|--|--|
|from|number|转出uid|
|to|number|转入uid|
|from_fields|arry|[转出金额操作](#fields)|
|to_fields|arry|[转入金额操作](#fields)|
|logtype|number|[日志类型](#logtype)|
|content|string|内容|
|buzid|string||

```json
param: { 
    "from": 1,
    "to": 2,
    "from_fields": [[ 1, 1, 10]],
    "to_fields": [[ 0, 0, 10 ]],
    "logtype":1001,
    "content":"测试转账",
    "buzid": "test_bussiness_id_001"
}
```

返回：
|参数名 |类型| 解释 |
|--|--|--|
|from|number|[balance](#balance)|
|to|number|[balance](#balance)|

```json
{
        "from": {{balance}},
        "to": {{balance}}
}
```

### 创建订单

** /api/v1/pay/mall/createOrder

参数:

```json
// amount: 充值金额   itemid: 商品Id   pay_channel: 支付方式 1Google 2Apple 3Paycenter   scene: 充值场景 1mall 2chat
 param: { "amount":12.5, "itemid": 2, "pay_channel":123, "scence": 3 }
```

返回：

```json
{
    "code": 0,
    "data": {
        "orderid": 11,
        "url": "http://payment-v1.tapmechat.com/pay/checkout/30004"
    },
    "err": null,
    "time": 1642413643
}
```

### 完成 Google 订单

** /api/v1/pay/mall/finishGoogleOrder

参数:

```json
// scene: 1:商城 2:Chat
// orderid: 订单ID
// gpid: Google订单ID
// token 方式通过 Google 云验证 
// 或者使用 signature & signdata
param: { 
    "scene": 1, 
    "token":"a1b3c3e2f2ad2378d", 
    "orderid": 123,
    "gpid": "GPA.3368-6754-5004-34365"
    "signature":"NwWm2IwQO03141KLerDBTHAlQz6l6vLsgQIBoeKAL8J6j3v2c3J+x8O8tzjrcTPanOsSDOeXLuUOgXiOaVhcgoFN9zAq3jSGvVplniYvZEkDP4T2S+VCUB0QL7\/i5Cfdm1pH3s+hsKUsoI4lfXh6rWf\/hRghDpRu0DrcCrjh1moleEWJcfw7QaALZzQSQOJHSiw7NScFkSXzEt5K3fZYEos0DIT6G2aooQap85z3I6eDH8biSMFWFFRBE6cVSsf2x3Zh4nW4wmkNssOwY4TgJmN2EDrH9BiREgwSzr4JdbtF6rj5UOzeOno2NJ4hZ39tU\/26Giwp2iRhs2vs5Nkovw==", 
    "signdata":"{\"orderId\":\"GPA.3339-6704-9049-21691\",\"packageName\":\"com.megabucks.slots\",\"productId\":\"sku_10\",\"purchaseTime\":1638169924690,\"purchaseState\":0,\"purchaseToken\":\"kmabldlbagbojbmpopcbagbd.AO-J1OyDjaOIl5bovpPPaXHq_1FqkCSgEGvR7Wto4o_nNlIL3W6alstifxeSYmQjTSFP1hROIYTfErFXa7s7Xz6attCWN0bZVg\",\"obfuscatedAccountId\":\"340\",\"obfuscatedProfileId\":\"100003\",\"acknowledged\":false}" 
    }
```

返回：

```json
{"code":0,"data":{"orderid":2, "status": 1},"err":null,"time":1641473015}
```

### 商城商品列表

** /api/v1/pay/mall/list

参数:

```json
param: {  }
```

返回：

```json
{
    "code": 0,
    "data": [
        {
            "ID": 1,
            "Appid": 12345,
            "Type": 1,
            "GoodName": "测试商品",
            "Price": 12,
            "SiteNo": "com.tapmechat.gift001",
            "BuyType": 0,
            "Num": 1,
            "Pic": "",
            "Note": "",
            "GiftNum": 0,
            "ValidDays": 0,
            "IsNew": 0,
            "Sort": 0,
            "Status": 0,
            "Tag": 0    // 商品tag, 1:普通商品，2：首充商品，3：特惠商品
        }
    ],
    "err": null,
    "time": 1641474752
}
```

## 钻石商城列表

** /api/v1/pay/mall/diamondStore

参数:

```json
param: {  }
```

返回：

```json
{
    "code": 0,
    "data": {
        "diamond": 0,   // 个人钻石数量
        "is_first_charge_user": ture, // 是否可以购买首充，true:可以，fale:不行
        "diamond_goods": [
            {
                "id": 2,
                "siteno": "com.tapmechat.dianmod001",
                "appid": 12345,
                "buy_type": 0,
                "goods_num": 1,
                "price": 12,
                "pic": "",
                "recommend": 0,
                "discount": 1.0, // 商品折扣
                "tag": 1 // 商品tag，1:普通商品,2:首充商品,3:优惠商品
            },
            {
                "id": 3,
                "siteno": "com.tapmechat.dianmod002",
                "appid": 12345,
                "buy_type": 0,
                "goods_num": 1,
                "price": 12,
                "pic": "",
                "recommend": 0,
                "discount": 1.0,  // 商品折扣
                "tag": 1 // 商品tag，1:普通商品,2:首充商品,3:优惠商品
            }
        ]
    },
    "err": null,
    "time": 1642062746
}
```


## 个人支付历史记录

** /api/v1/pay/user/paymentHistory

参数:

```json
//date:查询日期
//conds:回传条件
param: {"date":"2022-01-01","page_size":10,"page_no":1,"conds": {"date": "202202","no": 3,"size": 2 }}
```
data需要查询某日："2022-01-01", 如果日期为0则会查询当月数据
```json
{
    "code": 0,
    "data": {
        "conds": {"date": "2022-01-01","no": 3,"size": 2 },
        "history": [
            {
                "type": 1000,
                "diams": 100,
                "time": "2022-02-10T16:33:07+08:00",
                "content": "oa变更"
            }
        ]
    },
    "err": null,
    "time": 1644915425
}
```

## 个人收益历史记录

** /api/v1/pay/user/earningHistory

参数:

```json
//date:查询日期
//conds:回传条件
param: {"date":"2022-01-01","page_size":10,"page_no":1,"conds": {"date": "202202","no": 3,"size": 2 }}
```
data需要查询某日："2022-01-01", 如果日期为0则会查询当月数据
响应：
```json
{
    "code": 0,
    "data": {
        "conds": {"date": "2022-01-01","no": 3,"size": 2 },
        "history": [
            {
                "type": 1000,
                "coins": 100,
                "time": "2022-02-10T16:13:49+08:00",
                "content": "oa变更"
            }
        ]
    },
    "err": null,
    "time": 1644915160
}
```
