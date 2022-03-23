# 礼物相关

## host

内网: `http://192.168.1.111:8083/`
外网测试: ``http://api-demo.tapmechat.com``

## 接口


## 获取礼物墙

** /api/v1/gift/getGiftWall

参数:

```json
param: {"page":1}  //总共两页 1,2
```

```json
{
    "code": 0,
    "data": [
        {
            "Id": 9,
            "extra": "花好月圆",
            "name": "花好月圆",
            "page": 2,
            "consume": 199,
            "icons": "http://",
            "resource_url": "http://",
            "enable": 1
        },
        {
            "Id": 10,
            "extra": "玫瑰花",
            "name": "玫瑰花",
            "page": 2,
            "consume": 299,
            "icons": "http://",
            "resource_url": "http://",
            "enable": 1
        },
        {
            "Id": 11,
            "extra": "留声机",
            "name": "留声机",
            "page": 2,
            "consume": 499,
            "icons": "http://",
            "resource_url": "http://",
            "enable": 1
        }
    ],
    "err": null,
    "time": 1644391789
}
```


## 礼物赠送

** /api/v1/gift/sendGift

参数:

```json
param: {"id":1,"recipient":21041661}  //id：礼物id , recipient:接受者id
```

```json
{
    "code": 0,
    "data": null,
    "err":  null,
    "time": 1644396232
}
```
