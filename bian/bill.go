package bian

import (
	"context"
	"github.com/leaf-rain/util/bian/dto"
	"github.com/leaf-rain/util/log"
	"go.uber.org/zap"
	"net/http"
)

const (
	UrlRecordChargeGet = "/sapi/v1/capital/deposit/hisrec"
)

// RecordChargeGet 获取充值历史记录
func (b binance) RecordChargeGet(ctx context.Context, req *dto.RecordChargeGetReq) ([]*dto.RecordChargeGetResp, error) {
	var err error
	var result []*dto.RecordChargeGetResp
	var data = req.ToString()
	url := b.BaseUrl + UrlRecordChargeGet
	err = b.BinanceRequest(ctx, url, data, &result, http.MethodGet)
	if err != nil {
		log.GetLogger().Error("[RecordChargeGet] BinanceGet failed",
			zap.Any("req", req),
			zap.Any("url", url),
			zap.String("data", data),
			zap.Error(err))
	}
	return result, err
}

// RecordCashGet 获取提现历史记录
func (b binance) RecordCashGet(ctx context.Context, req *dto.RecordCashGetReq) ([]*dto.RecordCashGetResp, error) {
	var err error
	var result []*dto.RecordCashGetResp
	var data = req.ToString()
	url := b.BaseUrl + UrlRecordChargeGet
	err = b.BinanceRequest(ctx, url, data, &result, http.MethodGet)
	if err != nil {
		log.GetLogger().Error("[RecordChargeGet] BinanceGet failed",
			zap.Any("req", req),
			zap.Any("url", url),
			zap.String("data", data),
			zap.Error(err))
	}
	return result, err
}
