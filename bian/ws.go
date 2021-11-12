package bian

import (
	"context"
	"github.com/leaf-rain/util/bian/dto"
	"github.com/leaf-rain/util/log"
	"go.uber.org/zap"
	"net/http"
)

const (
	UrlListenKey = "/api/v3/userDataStream"
)

func (b binance) JobStart(ctx context.Context) {
	// 获取listen key

	// 建立长连接

	// 重试机制
}

func (b binance) JobClose(ctx context.Context) { // 销毁listen key
	_ = b.ListenKeyDestroy(ctx, &dto.ListenKeyValue{ListenKey: b.ListenKey})
}

// ListenKeyGenerate 生成 Listen Key
func (b binance) ListenKeyGenerate(ctx context.Context) (string, error) {
	var err error
	var result dto.ListenKeyValue
	url := b.BaseUrl + UrlListenKey
	err = b.BinanceRequest(ctx, url, "", &result, http.MethodPost)
	if err != nil {
		log.GetLogger().Error("[ListenKeyGenerate] BinanceRequest failed",
			zap.Any("url", url),
			zap.Error(err))
		return "", err
	}
	return result.ListenKey, nil
}

// ListenKeyRenew 续约 Listen Key
func (b binance) ListenKeyRenew(ctx context.Context, req *dto.ListenKeyValue) error {
	var err error
	url := b.BaseUrl + UrlListenKey
	url = url + "?listenKey=" + req.ListenKey
	err = b.BinanceRequest(ctx, url, "", nil, http.MethodPut)
	if err != nil {
		log.GetLogger().Error("[ListenKeyGenerate] BinanceRequest failed",
			zap.Any("url", url),
			zap.Error(err))
		return err
	}
	return nil
}

// ListenKeyDestroy 销毁 Listen Key
func (b binance) ListenKeyDestroy(ctx context.Context, req *dto.ListenKeyValue) error {
	var err error
	url := b.BaseUrl + UrlListenKey
	url = url + "?listenKey=" + req.ListenKey
	err = b.BinanceRequest(ctx, url, "", nil, http.MethodPut)
	if err != nil {
		log.GetLogger().Error("[ListenKeyGenerate] BinanceRequest failed",
			zap.Any("url", url),
			zap.Error(err))
		return err
	}
	return nil
}
