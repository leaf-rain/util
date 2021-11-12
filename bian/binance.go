package bian

import (
	"context"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/leaf-rain/util/log"
	"go.uber.org/zap"
	"net/http"
)

const baseUrl = "https://api.binance.com"

type binance struct {
	BaseUrl   string `json:"base_url"`
	WsUrl     string `json:"ws_url"`
	ApiKey    string `json:"api_key"`
	SecretKey string `json:"secret_key"`
	ListenKey string `json:"listen_key"`
}

func NewBinance() binance {

	return binance{
		BaseUrl:   "https://api.binance.com",
		WsUrl:     "wss://stream.binance.com:9443",
		ApiKey:    "459GGkaLLZuP7leFGfqF9FrmZPXix7d6Z30wgWLoPuQqUVv8Q1UyIVBSwTS98T2F",
		SecretKey: "",
	}
}

func (b binance) BinanceRequest(ctx context.Context, path string, data string, result interface{}, method string) (err error) {
	var request *http.Request
	var response *http.Response
	var respJson *simplejson.Json
	var dataJson []byte
	var code int
	if len(data) > 0 {
		sign := b.SignGet(data, "")
		data = data + "&signature=" + sign
		path = path + "?" + data
	}
	client := &http.Client{}
	request, err = http.NewRequest(method, path, nil)
	if err != nil {
		log.GetLogger().Error("[BinanceGet] http.NewRequest failed",
			zap.Any("path", path),
			zap.Any("data", data),
			zap.Error(err))
		return
	}
	request.Header.Add("X-MBX-APIKEY", b.ApiKey)
	response, err = client.Do(request)
	if err != nil {
		log.GetLogger().Error("[BinanceGet] client.Do failed",
			zap.Any("path", path),
			zap.Any("result", result),
			zap.Any("data", data),
			zap.Error(err))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.GetLogger().Error("[BinanceGet] StatusCode != 200",
			zap.Any("path", path),
			zap.Any("data", data),
			zap.Any("code", response.StatusCode))
		return CodeCheck(code)
	}
	respJson, err = simplejson.NewFromReader(response.Body)
	code, _ = respJson.Get("code").Int()
	if code != 0 {
		log.GetLogger().Error("[BinanceGet] error code",
			zap.Any("path", path),
			zap.Any("data", data),
			zap.Any("code", code))
		return CodeCheck(code)
	}
	if result != nil {
		dataJson, err = respJson.MarshalJSON()
		if err != nil {
			log.GetLogger().Error("[BinanceGet] respJson.MarshalJSON failed",
				zap.Any("path", path),
				zap.Any("result", result),
				zap.Any("data", data),
				zap.Error(err))
			return
		}
		err = json.Unmarshal(dataJson, result)
		if err != nil {
			log.GetLogger().Error("[BinanceGet] json.Unmarshal failed",
				zap.Any("path", path),
				zap.Any("dataJson", string(dataJson)),
				zap.Any("data", data),
				zap.Error(err))
			return err
		}
	}
	return nil
}
