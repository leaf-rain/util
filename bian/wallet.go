package bian

import (
	"context"
	"github.com/leaf-rain/util/bian/dto"
	"github.com/leaf-rain/util/log"
	"go.uber.org/zap"
	"net/http"
)

const (
	UrlAddressGet  = "/sapi/v1/capital/deposit/address"
	UrlAccountGet  = "/api/v3/account"
	UrlTransferGet = "/sapi/v1/capital/withdraw/apply"
)

// AddressGet 获取交易地址
func (b binance) AddressGet(ctx context.Context, req *dto.AddressGetReq) (*dto.AddressGetResp, error) {
	var err error
	var result dto.AddressGetResp
	var data = req.ToString()
	url := b.BaseUrl + UrlAddressGet
	err = b.BinanceRequest(ctx, url, data, &result, http.MethodGet)
	if err != nil {
		log.GetLogger().Error("[AddressGet] BinanceGet failed",
			zap.Any("req", req),
			zap.Any("url", url),
			zap.String("data", data),
			zap.Error(err))
	}
	return &result, err
}

// AccountGet 获取帐号信息
func (b binance) AccountGet(ctx context.Context, req *dto.AccountGetReq) (*dto.AccountGetResp, error) {
	var err error
	var result dto.AccountGetResp
	var data = req.ToString()
	url := b.BaseUrl + UrlAccountGet
	err = b.BinanceRequest(ctx, url, data, &result, http.MethodGet)
	if err != nil {
		log.GetLogger().Error("[AddressGet] BinanceGet failed",
			zap.Any("req", req),
			zap.Any("url", url),
			zap.String("data", data),
			zap.Error(err))
	}
	return &result, err
}

// Transfer // 提币
func (b binance) Transfer(ctx context.Context, req *dto.TransferReq) (*dto.TransferResp, error) {
	var err error
	var result dto.TransferResp
	var data = req.ToString()
	url := b.BaseUrl + UrlTransferGet
	err = b.BinanceRequest(ctx, url, data, &result, http.MethodGet)
	if err != nil {
		log.GetLogger().Error("[AddressGet] BinanceGet failed",
			zap.Any("req", req),
			zap.Any("url", url),
			zap.String("data", data),
			zap.Error(err))
	}
	return &result, err
}
