# 新接口
**http://api-v1.touchat.live/api/v1/client**

 方法名参考  [method.md](method.md)  
`request` 统一使用 header: `"Content-Type": "application/json`, `JSON` 固定格式如下：

```js
{              
  "time":1630995979,             //请求时间
  "ver_app":"1.0.4.1",           //版本(包)
  "ver_res":"1.2.4",             //版本(资源)
  "mtkey": mtkey,                //签名key
  "appid":200,                   //渠道id
  "uid":9123690,                 //uid
  "sesskey": sesskey,            //session
  "lang":"EN",                   //语言
  "param": param,                //参数
  "pack": 1,                     // 包类别， 1: 男包, 2: 女包  0: 不区分
  "sig":sig,                     //签名校验
  "method":"3af7e5ef"            //方法名
}
```


## 举例，登陆调用

** /api/v1/client

request:

```json
{
  "appid": 10002,
  "lang": "zh",
  "mtkey": "MTKEY",
  "pack": 2,
  "param": {
    "deviceid": "503762fe84bab6e4bf98d5137a900e33",
    "devname": "SM-A205F",
    "fb_access_token": "",
    "fb_appid": "",
    "fb_openid": "",
    "gender": 2,
    "google_access_token": "",
    "google_openid": "",
    "idfa": "",
    "platform": "visitor",
    "siteuid": "",
    "sys_lang": "zh",
    "sys_timezone": "CN",
    "sys_version": "android_30"
  },
  "sesskey": "d0d12ab0ff8ff36b050c81b37573e9a8",
  "sig": "5f569473c061679ce83e143a95e6cd1d",
  "time": 1644660047639,
  "uid": 0,
  "ver_app": "1.0.0",
  "ver_res": "1.0.0",
  "method":"3af7e5ef" 
}
```

response:

```json
{
    "code": 0,
    "data": {
        "account": {
            "uid": 42513290,
            "Appid": 10002,
            "Coins": 0,
            "Diamonds": 800,
            "PayCount": 0,
            "PayAmount": 0
        },
        "guid": "65affb19e1b3fa860ce063500ae27452",
        "imToken": "IfNaIYmkslWciXqBzI5+KDUhZAT2gVMq1rsW+ZI6jQQ=@1apu.sg.rongnav.com;1apu.sg.rongcfg.com",
        "info": {
            "uid": 42513290,
            "Appid": 10002,
            "nickname": "USER_42513290",
            "borndate": 0,
            "HeadPic": "",
            "Gender": 2,
            "Age": 0,
            "CountryCode": "",
            "Lang": "zh",
            "SecondLang": "",
            "GeetingWords": "",
            "GeetingPic": "",
            "Intro": "",
            "RegTime": 1644660324,
            "LoginTime": 1646877975,
            "RegIP": "127.0.0.1",
            "LoginIP": "127.0.0.1",
            "ImToken": "IfNaIYmkslWciXqBzI5+KDUhZAT2gVMq1rsW+ZI6jQQ=@1apu.sg.rongnav.com;1apu.sg.rongcfg.com",
            "Status": 0,
            "RegDevid": "503762fe84bab6e4bf98d5137a900e33",
            "RegIdfa": "",
            "RegPlace": "",
            "RegAppid": 10002,
            "LoginDevid": "503762fe84bab6e4bf98d5137a900e33",
            "LoginIdfa": "503762fe84bab6e4bf98d5137a900e33",
            "LoginPlace": "",
            "LoginDevname": "SM-A205F",
            "LoginSid": 0,
            "VerApp": "1.0.0",
            "VerRes": "1.0.0",
            "SysVersion": "android_30",
            "SysTimezone": "CN",
            "SysLang": "zh",
            "IsOnline": false,
            "IsVerifyPkg": false
        },
        "session": "2e1153b66fe105afe7e8ab35d20fbc19",
        "uid": 42513290
    },
    "err": null,
    "time": 1646877975
}
```