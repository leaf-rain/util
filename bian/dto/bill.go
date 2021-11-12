package dto

import (
	"github.com/leaf-rain/util/bian/consts"
	"github.com/leaf-rain/util/tool"
	"net/url"
)

type RecordChargeGetReq struct {
	Coin       string             `json:"coin"`
	Status     RecordChargeStatus `json:"status"`
	StartTime  int64              `json:"startTime"`
	EndTime    int64              `json:"endTime"`
	Offset     int                `json:"offset"`
	Limit      int                `json:"limit"`
	RecvWindow int64              `json:"recvWindow"`
	Timestamp  int64              `json:"timestamp"`
}

type RecordChargeStatus int

const (
	RecordStatus_Pending  RecordChargeStatus = 0 // pending
	RecordStatus_Credited RecordChargeStatus = 6 // 质押中，无法提取
	RecordStatus_Success  RecordChargeStatus = 1 // 成功
)

func (r RecordChargeGetReq) ToString() string {
	if r.RecvWindow <= 0 {
		r.RecvWindow = consts.RecvWindow
	}
	if r.Timestamp <= 0 {
		r.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	if len(r.Coin) > 0 {
		values.Add("coin", r.Coin)
	}
	if r.Status > 0 {
		values.Add("status", tool.IntToStr(int(r.Status)))
	}
	if r.StartTime > 0 {
		values.Add("startTime", tool.Int64ToStr(r.StartTime))
	}
	if r.EndTime > 0 {
		values.Add("endTime", tool.Int64ToStr(r.EndTime))
	}
	if r.Offset > 0 {
		values.Add("offset", tool.IntToStr(r.Offset))
	}
	if r.Limit > 0 {
		values.Add("limit", tool.IntToStr(r.Limit))
	}
	values.Add("recvWindow", tool.Int64ToString(r.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(r.Timestamp))
	return values.Encode()
}

type RecordChargeGetResp struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxID          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int    `json:"transferType"`
	ConfirmTimes  string `json:"confirmTimes"`
	UnlockConfirm int    `json:"unlockConfirm"` // 解锁需要的网络确认次数
	WalletType    int    `json:"walletType"`
}

type RecordCashGetReq struct {
	Coin            string           `json:"coin"`
	WithdrawOrderId string           `json:"withdrawOrderId"`
	Status          RecordCashStatus `json:"status"`
	StartTime       int64            `json:"startTime"`
	EndTime         int64            `json:"endTime"`
	Offset          int              `json:"offset"`
	Limit           int              `json:"limit"`
	RecvWindow      int64            `json:"recvWindow"`
	Timestamp       int64            `json:"timestamp"`
}

type RecordCashStatus int

const (
	RecordCashStatus_SendEmail  RecordCashStatus = 0 // 已发送确认Email
	RecordCashStatus_Cancel     RecordCashStatus = 1 // 已被用户取消
	RecordCashStatus_Waiting    RecordCashStatus = 2 // 等待确认
	RecordCashStatus_Refuse     RecordCashStatus = 3 // 被拒绝
	RecordCashStatus_Processing RecordCashStatus = 4 // 处理中
	RecordCashStatus_Failed     RecordCashStatus = 5 // 提现交易失败
	RecordCashStatus_Success    RecordCashStatus = 6 // 提现完成
)

func (r RecordCashGetReq) ToString() string {
	if r.RecvWindow <= 0 {
		r.RecvWindow = consts.RecvWindow
	}
	if r.Timestamp <= 0 {
		r.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	if len(r.Coin) > 0 {
		values.Add("coin", r.Coin)
	}
	if r.Status > 0 {
		values.Add("status", tool.IntToStr(int(r.Status)))
	}
	if r.StartTime > 0 {
		values.Add("startTime", tool.Int64ToStr(r.StartTime))
	}
	if r.EndTime > 0 {
		values.Add("endTime", tool.Int64ToStr(r.EndTime))
	}
	if r.Offset > 0 {
		values.Add("offset", tool.IntToStr(r.Offset))
	}
	if r.Limit > 0 {
		values.Add("limit", tool.IntToStr(r.Limit))
	}
	values.Add("recvWindow", tool.Int64ToString(r.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(r.Timestamp))
	return values.Encode()
}

type RecordCashGetResp struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`    // 提现转出金额
	ApplyTime       string `json:"applyTime"` // UTC 时间
	Coin            string `json:"coin"`
	ID              string `json:"id"`              // 该笔提现在币安的id
	WithdrawOrderID string `json:"withdrawOrderId"` // 自定义ID, 如果没有则不返回该字段
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"` // 1: 站内转账, 0: 站外转账
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"` // 手续费
	ConfirmNo       int    `json:"confirmNo"`      // 提现确认数
	TxID            string `json:"txId"`           // 提现交易id
}
