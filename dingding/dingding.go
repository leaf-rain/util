package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leaf-rain/util/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Config struct {
	Secret    string   // 密钥
	Urls      string   // 路径
	AtMobiles []string // @的手机号（固定）
	AtUserIds []string // @的用户id（固定）
}

type Option struct {
	AtMobiles []string // @的手机号（额外添加）
	AtUserIds []string // @的用户id（额外添加）
}

type DingSrv struct {
	opt  *Config
	lock *sync.Mutex // 锁
}

func NewDingSrv(opt *Config) DingSrv {
	return DingSrv{opt: opt, lock: new(sync.Mutex)}
}

type ResultForDingDing struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (d DingSrv) DingdingSend(message string, opt ...Option) (bool, error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	var reqUrl = d.opt.Urls + sign(d.opt.Secret)
	var AtMobiles []string
	var AtUserIds []string
	for _, item := range opt {
		AtMobiles = append(d.opt.AtMobiles, item.AtMobiles...)
		AtUserIds = append(d.opt.AtUserIds, item.AtUserIds...)
	}
	var request, err = newRequestBody(AtMobiles, AtUserIds, message, false)
	if err != nil {
		log.GetLogger().Error("[DingdingSend] newRequestBody failed",
			zap.Any("setting", d.opt),
			zap.Any("message", message),
			zap.Error(err))
		return false, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(request))
	if err != nil {
		log.GetLogger().Error("[DingdingSend] NewRequest failed",
			zap.Any("setting", d.opt),
			zap.Any("message", message),
			zap.Error(err))
		return false, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		log.GetLogger().Error("[DingdingSend] request Do failed",
			zap.Any("setting", d.opt),
			zap.Any("message", message),
			zap.Error(err))
		return false, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.GetLogger().Error("[DingdingSend] ioutil.ReadAll failed",
			zap.Any("setting", d.opt),
			zap.Any("message", message),
			zap.Error(err))
		return false, err
	}
	var sta ResultForDingDing
	err = json.Unmarshal(body, &sta)
	if err != nil {
		log.GetLogger().Error("[DingdingSend] json.Unmarshal failed",
			zap.Any("setting", d.opt),
			zap.Any("message", message),
			zap.Error(err))
		return false, err
	}
	if sta.Errcode == 0 || sta.Errmsg == "ok" {
		log.GetLogger().Info("[DingdingSend] success",
			zap.Any("message", message))
		return true, nil
	}
	return false, errors.New("Error response:---" + string(body))
}

func newRequestBody(atm, atu []string, data string, isAtAll bool) (string, error) {
	reqBody := struct {
		At struct {
			AtMobiles []string `json:"atMobiles"`
			AtUserIds []string `json:"atUserIds"`
			IsAtAll   bool     `json:"isAtAll"`
		} `json:"at"`
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		Msgtype string `json:"msgtype"`
	}{}
	reqBody.Text.Content = data
	reqBody.Msgtype = "text"
	reqBody.At.AtMobiles = atm
	reqBody.At.AtUserIds = atu
	reqBody.At.IsAtAll = isAtAll
	reqData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	return string(reqData), nil
}

func sign(secret string) string {
	timestamp := fmt.Sprint(time.Now().UnixNano() / 1000000)
	secStr := timestamp + "\n" + secret
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(secStr))
	sum := h.Sum(nil)
	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)
	return "&timestamp=" + timestamp + "&sign=" + urlEncode
}
