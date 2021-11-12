package dto

import (
	"github.com/leaf-rain/util/bian/consts"
	"github.com/leaf-rain/util/tool"
	"net/url"
)

type AddressGetReq struct {
	Coin       string `json:"coin"`
	Network    string `json:"network"`
	RecvWindow int64  `json:"recvWindow"`
	Timestamp  int64  `json:"timestamp"`
}

type AddressGetResp struct {
	Coin    string `json:"coin"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

func (a *AddressGetReq) ToString() string {
	if a.RecvWindow <= 0 {
		a.RecvWindow = consts.RecvWindow
	}
	if a.Timestamp <= 0 {
		a.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	values.Add("coin", a.Coin)
	if len(a.Network) > 0 {
		values.Add("network", a.Network)
	}
	values.Add("recvWindow", tool.Int64ToString(a.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(a.Timestamp))
	return values.Encode()
}

type AccountGetReq struct {
	RecvWindow int64 `json:"recvWindow"`
	Timestamp  int64 `json:"timestamp"`
}

type AccountGetResp struct {
	MakerCommission  int        `json:"makerCommission"`
	TakerCommission  int        `json:"takerCommission"`
	BuyerCommission  int        `json:"buyerCommission"`
	SellerCommission int        `json:"sellerCommission"`
	CanTrade         bool       `json:"canTrade"`
	CanWithdraw      bool       `json:"canWithdraw"`
	CanDeposit       bool       `json:"canDeposit"`
	UpdateTime       int        `json:"updateTime"`
	AccountType      string     `json:"accountType"`
	Balances         []Balances `json:"balances"`
	Permissions      []string   `json:"permissions"`
}

type Balances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

func (a *AccountGetReq) ToString() string {
	if a.RecvWindow <= 0 {
		a.RecvWindow = consts.RecvWindow
	}
	if a.Timestamp <= 0 {
		a.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	values.Add("recvWindow", tool.Int64ToString(a.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(a.Timestamp))
	return values.Encode()
}

type TransferReq struct {
	Coin               string `json:"coin"`
	WithdrawOrderId    string `json:"withdrawOrderId"`    // 自定义提币id
	Network            string `json:"network"`            // 提币网络
	Address            string `json:"address"`            // 提币地址
	AddressTag         string `json:"addressTag"`         // 某些币种例如 XRP,XMR 允许填写次级地址标签
	Amount             string `json:"amount"`             // 提币金额
	TransactionFeeFlag bool   `json:"transactionFeeFlag"` // 当站内转账时免手续费, true: 手续费归资金转入方; false: 手续费归资金转出方; . 默认 false.
	Name               string `json:"name"`               // 地址的备注，填写该参数后会加入该币种的提现地址簿。地址簿上限为20，超出后会造成提现失败。地址中的空格需要encode成%20
	WalletType         int    `json:"walletType"`         // 表示出金使用的钱包，0为现货钱包，1为资金钱包，默认为现货钱包
	RecvWindow         int64  `json:"recvWindow"`
	Timestamp          int64  `json:"timestamp"`
}

type TransferResp struct {
	ID string `json:"id"` // 订单号
}

func (a *TransferReq) ToString() string {
	if a.RecvWindow <= 0 {
		a.RecvWindow = consts.RecvWindow
	}
	if a.Timestamp <= 0 {
		a.Timestamp = tool.GetTimeUnixMilli()
	}
	var values = url.Values{}
	if len(a.Coin) > 0 {
		values.Add("coin", a.Coin)
	}
	if len(a.WithdrawOrderId) > 0 {
		values.Add("withdrawOrderId", a.WithdrawOrderId)
	}
	if len(a.Network) > 0 {
		values.Add("network", a.Network)
	}
	if len(a.Address) > 0 {
		values.Add("address", a.Address)
	}
	if len(a.AddressTag) > 0 {
		values.Add("addressTag", a.AddressTag)
	}
	if len(a.Amount) > 0 {
		values.Add("amount", a.Amount)
	}
	if a.TransactionFeeFlag {
		values.Add("transactionFeeFlag", "true")
	}
	if len(a.Name) > 0 {
		values.Add("name", a.Name)
	}
	if a.WalletType > 0 {
		values.Add("walletType", tool.IntToStr(a.WalletType))
	}
	values.Add("recvWindow", tool.Int64ToString(a.RecvWindow))
	values.Add("timestamp", tool.Int64ToStr(a.Timestamp))
	return values.Encode()
}
