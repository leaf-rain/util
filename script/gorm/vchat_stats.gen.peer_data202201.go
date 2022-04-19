package gorm

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _PeerData202201Mgr struct {
	*_BaseMgr
}

// PeerData202201Mgr open func
func PeerData202201Mgr(db *gorm.DB) *_PeerData202201Mgr {
	if db == nil {
		panic(fmt.Errorf("PeerData202201Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PeerData202201Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("peer_data202201"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PeerData202201Mgr) GetTableName() string {
	return "peer_data202201"
}

// Reset 重置gorm会话
func (obj *_PeerData202201Mgr) Reset() *_PeerData202201Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PeerData202201Mgr) Get() (result PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PeerData202201Mgr) Gets() (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PeerData202201Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_PeerData202201Mgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUID uid获取 uid
func (obj *_PeerData202201Mgr) WithUID(uid int64) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithLevel level获取 主播等级
func (obj *_PeerData202201Mgr) WithLevel(level int8) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithGuildID guild_id获取 工会id
func (obj *_PeerData202201Mgr) WithGuildID(guildID int64) Option {
	return optionFunc(func(o *options) { o.query["guild_id"] = guildID })
}

// WithGoldBalance gold_balance获取 当日金币余额
func (obj *_PeerData202201Mgr) WithGoldBalance(goldBalance float64) Option {
	return optionFunc(func(o *options) { o.query["gold_balance"] = goldBalance })
}

// WithGoldIncome gold_income获取 金币收益
func (obj *_PeerData202201Mgr) WithGoldIncome(goldIncome float64) Option {
	return optionFunc(func(o *options) { o.query["gold_income"] = goldIncome })
}

// WithGoldIncomeForCall gold_income_for_call获取 金币收益 1v1视频通话
func (obj *_PeerData202201Mgr) WithGoldIncomeForCall(goldIncomeForCall float64) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_call"] = goldIncomeForCall })
}

// WithGoldIncomeForGift gold_income_for_gift获取 金币收益 礼物
func (obj *_PeerData202201Mgr) WithGoldIncomeForGift(goldIncomeForGift float64) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_gift"] = goldIncomeForGift })
}

// WithGoldIncomeForReward gold_income_for_reward获取 金币收益 奖励
func (obj *_PeerData202201Mgr) WithGoldIncomeForReward(goldIncomeForReward float64) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_reward"] = goldIncomeForReward })
}

// WithGoldIncomeForCallTimes gold_income_for_call_times获取 金币收益 1v1视频通话次数
func (obj *_PeerData202201Mgr) WithGoldIncomeForCallTimes(goldIncomeForCallTimes int) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_call_times"] = goldIncomeForCallTimes })
}

// WithGoldIncomeForGiftTimes gold_income_for_gift_times获取 金币收益 礼物次数
func (obj *_PeerData202201Mgr) WithGoldIncomeForGiftTimes(goldIncomeForGiftTimes int) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_gift_times"] = goldIncomeForGiftTimes })
}

// WithGoldIncomeForRewardTimes gold_income_for_reward_times获取 金币收益 奖励次数
func (obj *_PeerData202201Mgr) WithGoldIncomeForRewardTimes(goldIncomeForRewardTimes int) Option {
	return optionFunc(func(o *options) { o.query["gold_income_for_reward_times"] = goldIncomeForRewardTimes })
}

// WithOnlineCount online_count获取 在线时长
func (obj *_PeerData202201Mgr) WithOnlineCount(onlineCount int) Option {
	return optionFunc(func(o *options) { o.query["online_count"] = onlineCount })
}

// WithCallCount call_count获取 通话时长
func (obj *_PeerData202201Mgr) WithCallCount(callCount int) Option {
	return optionFunc(func(o *options) { o.query["call_count"] = callCount })
}

// WithVideoCount video_count获取 1v1视频通话时长
func (obj *_PeerData202201Mgr) WithVideoCount(videoCount int) Option {
	return optionFunc(func(o *options) { o.query["video_count"] = videoCount })
}

// WithPeerCount peer_count获取 匹配通话时长
func (obj *_PeerData202201Mgr) WithPeerCount(peerCount int) Option {
	return optionFunc(func(o *options) { o.query["peer_count"] = peerCount })
}

// WithFreeRate free_rate获取 空闲率
func (obj *_PeerData202201Mgr) WithFreeRate(freeRate float64) Option {
	return optionFunc(func(o *options) { o.query["free_rate"] = freeRate })
}

// WithCallArouse call_arouse获取 通话唤起次数
func (obj *_PeerData202201Mgr) WithCallArouse(callArouse int) Option {
	return optionFunc(func(o *options) { o.query["call_arouse"] = callArouse })
}

// WithCallSend call_send获取 通话被叫次数
func (obj *_PeerData202201Mgr) WithCallSend(callSend int) Option {
	return optionFunc(func(o *options) { o.query["call_send"] = callSend })
}

// WithCallConnected call_connected获取 通话接通次数
func (obj *_PeerData202201Mgr) WithCallConnected(callConnected int) Option {
	return optionFunc(func(o *options) { o.query["call_connected"] = callConnected })
}

// WithFreeConnectRate free_connect_rate获取 通话接通率
func (obj *_PeerData202201Mgr) WithFreeConnectRate(freeConnectRate float64) Option {
	return optionFunc(func(o *options) { o.query["free_connect_rate"] = freeConnectRate })
}

// WithPeerSend peer_send获取 通话唤起次数
func (obj *_PeerData202201Mgr) WithPeerSend(peerSend int) Option {
	return optionFunc(func(o *options) { o.query["peer_send"] = peerSend })
}

// WithPeerReply peer_reply获取 通话被叫次数
func (obj *_PeerData202201Mgr) WithPeerReply(peerReply int) Option {
	return optionFunc(func(o *options) { o.query["peer_reply"] = peerReply })
}

// WithPeerConnected peer_connected获取 通话接通次数
func (obj *_PeerData202201Mgr) WithPeerConnected(peerConnected int) Option {
	return optionFunc(func(o *options) { o.query["peer_connected"] = peerConnected })
}

// WithPeerConnectRate peer_connect_rate获取 通话接通率
func (obj *_PeerData202201Mgr) WithPeerConnectRate(peerConnectRate float64) Option {
	return optionFunc(func(o *options) { o.query["peer_connect_rate"] = peerConnectRate })
}

// WithCmsCount cms_count获取 短信发送次数
func (obj *_PeerData202201Mgr) WithCmsCount(cmsCount int) Option {
	return optionFunc(func(o *options) { o.query["cms_count"] = cmsCount })
}

// WithCreateDate create_date获取 创建日期，格式：20060102
func (obj *_PeerData202201Mgr) WithCreateDate(createDate int) Option {
	return optionFunc(func(o *options) { o.query["create_date"] = createDate })
}

// GetByOption 功能选项模式获取
func (obj *_PeerData202201Mgr) GetByOption(opts ...Option) (result PeerData202201, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PeerData202201Mgr) GetByOptions(opts ...Option) (results []*PeerData202201, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 ID
func (obj *_PeerData202201Mgr) GetFromID(id uint32) (result PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_PeerData202201Mgr) GetBatchFromID(ids []uint32) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUID 通过uid获取内容 uid
func (obj *_PeerData202201Mgr) GetFromUID(uid int64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找 uid
func (obj *_PeerData202201Mgr) GetBatchFromUID(uids []int64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 主播等级
func (obj *_PeerData202201Mgr) GetFromLevel(level int8) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 主播等级
func (obj *_PeerData202201Mgr) GetBatchFromLevel(levels []int8) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromGuildID 通过guild_id获取内容 工会id
func (obj *_PeerData202201Mgr) GetFromGuildID(guildID int64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`guild_id` = ?", guildID).Find(&results).Error

	return
}

// GetBatchFromGuildID 批量查找 工会id
func (obj *_PeerData202201Mgr) GetBatchFromGuildID(guildIDs []int64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`guild_id` IN (?)", guildIDs).Find(&results).Error

	return
}

// GetFromGoldBalance 通过gold_balance获取内容 当日金币余额
func (obj *_PeerData202201Mgr) GetFromGoldBalance(goldBalance float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_balance` = ?", goldBalance).Find(&results).Error

	return
}

// GetBatchFromGoldBalance 批量查找 当日金币余额
func (obj *_PeerData202201Mgr) GetBatchFromGoldBalance(goldBalances []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_balance` IN (?)", goldBalances).Find(&results).Error

	return
}

// GetFromGoldIncome 通过gold_income获取内容 金币收益
func (obj *_PeerData202201Mgr) GetFromGoldIncome(goldIncome float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income` = ?", goldIncome).Find(&results).Error

	return
}

// GetBatchFromGoldIncome 批量查找 金币收益
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncome(goldIncomes []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income` IN (?)", goldIncomes).Find(&results).Error

	return
}

// GetFromGoldIncomeForCall 通过gold_income_for_call获取内容 金币收益 1v1视频通话
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForCall(goldIncomeForCall float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_call` = ?", goldIncomeForCall).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForCall 批量查找 金币收益 1v1视频通话
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForCall(goldIncomeForCalls []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_call` IN (?)", goldIncomeForCalls).Find(&results).Error

	return
}

// GetFromGoldIncomeForGift 通过gold_income_for_gift获取内容 金币收益 礼物
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForGift(goldIncomeForGift float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_gift` = ?", goldIncomeForGift).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForGift 批量查找 金币收益 礼物
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForGift(goldIncomeForGifts []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_gift` IN (?)", goldIncomeForGifts).Find(&results).Error

	return
}

// GetFromGoldIncomeForReward 通过gold_income_for_reward获取内容 金币收益 奖励
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForReward(goldIncomeForReward float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_reward` = ?", goldIncomeForReward).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForReward 批量查找 金币收益 奖励
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForReward(goldIncomeForRewards []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_reward` IN (?)", goldIncomeForRewards).Find(&results).Error

	return
}

// GetFromGoldIncomeForCallTimes 通过gold_income_for_call_times获取内容 金币收益 1v1视频通话次数
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForCallTimes(goldIncomeForCallTimes int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_call_times` = ?", goldIncomeForCallTimes).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForCallTimes 批量查找 金币收益 1v1视频通话次数
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForCallTimes(goldIncomeForCallTimess []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_call_times` IN (?)", goldIncomeForCallTimess).Find(&results).Error

	return
}

// GetFromGoldIncomeForGiftTimes 通过gold_income_for_gift_times获取内容 金币收益 礼物次数
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForGiftTimes(goldIncomeForGiftTimes int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_gift_times` = ?", goldIncomeForGiftTimes).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForGiftTimes 批量查找 金币收益 礼物次数
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForGiftTimes(goldIncomeForGiftTimess []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_gift_times` IN (?)", goldIncomeForGiftTimess).Find(&results).Error

	return
}

// GetFromGoldIncomeForRewardTimes 通过gold_income_for_reward_times获取内容 金币收益 奖励次数
func (obj *_PeerData202201Mgr) GetFromGoldIncomeForRewardTimes(goldIncomeForRewardTimes int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_reward_times` = ?", goldIncomeForRewardTimes).Find(&results).Error

	return
}

// GetBatchFromGoldIncomeForRewardTimes 批量查找 金币收益 奖励次数
func (obj *_PeerData202201Mgr) GetBatchFromGoldIncomeForRewardTimes(goldIncomeForRewardTimess []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`gold_income_for_reward_times` IN (?)", goldIncomeForRewardTimess).Find(&results).Error

	return
}

// GetFromOnlineCount 通过online_count获取内容 在线时长
func (obj *_PeerData202201Mgr) GetFromOnlineCount(onlineCount int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`online_count` = ?", onlineCount).Find(&results).Error

	return
}

// GetBatchFromOnlineCount 批量查找 在线时长
func (obj *_PeerData202201Mgr) GetBatchFromOnlineCount(onlineCounts []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`online_count` IN (?)", onlineCounts).Find(&results).Error

	return
}

// GetFromCallCount 通过call_count获取内容 通话时长
func (obj *_PeerData202201Mgr) GetFromCallCount(callCount int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_count` = ?", callCount).Find(&results).Error

	return
}

// GetBatchFromCallCount 批量查找 通话时长
func (obj *_PeerData202201Mgr) GetBatchFromCallCount(callCounts []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_count` IN (?)", callCounts).Find(&results).Error

	return
}

// GetFromVideoCount 通过video_count获取内容 1v1视频通话时长
func (obj *_PeerData202201Mgr) GetFromVideoCount(videoCount int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`video_count` = ?", videoCount).Find(&results).Error

	return
}

// GetBatchFromVideoCount 批量查找 1v1视频通话时长
func (obj *_PeerData202201Mgr) GetBatchFromVideoCount(videoCounts []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`video_count` IN (?)", videoCounts).Find(&results).Error

	return
}

// GetFromPeerCount 通过peer_count获取内容 匹配通话时长
func (obj *_PeerData202201Mgr) GetFromPeerCount(peerCount int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_count` = ?", peerCount).Find(&results).Error

	return
}

// GetBatchFromPeerCount 批量查找 匹配通话时长
func (obj *_PeerData202201Mgr) GetBatchFromPeerCount(peerCounts []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_count` IN (?)", peerCounts).Find(&results).Error

	return
}

// GetFromFreeRate 通过free_rate获取内容 空闲率
func (obj *_PeerData202201Mgr) GetFromFreeRate(freeRate float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`free_rate` = ?", freeRate).Find(&results).Error

	return
}

// GetBatchFromFreeRate 批量查找 空闲率
func (obj *_PeerData202201Mgr) GetBatchFromFreeRate(freeRates []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`free_rate` IN (?)", freeRates).Find(&results).Error

	return
}

// GetFromCallArouse 通过call_arouse获取内容 通话唤起次数
func (obj *_PeerData202201Mgr) GetFromCallArouse(callArouse int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_arouse` = ?", callArouse).Find(&results).Error

	return
}

// GetBatchFromCallArouse 批量查找 通话唤起次数
func (obj *_PeerData202201Mgr) GetBatchFromCallArouse(callArouses []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_arouse` IN (?)", callArouses).Find(&results).Error

	return
}

// GetFromCallSend 通过call_send获取内容 通话被叫次数
func (obj *_PeerData202201Mgr) GetFromCallSend(callSend int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_send` = ?", callSend).Find(&results).Error

	return
}

// GetBatchFromCallSend 批量查找 通话被叫次数
func (obj *_PeerData202201Mgr) GetBatchFromCallSend(callSends []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_send` IN (?)", callSends).Find(&results).Error

	return
}

// GetFromCallConnected 通过call_connected获取内容 通话接通次数
func (obj *_PeerData202201Mgr) GetFromCallConnected(callConnected int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_connected` = ?", callConnected).Find(&results).Error

	return
}

// GetBatchFromCallConnected 批量查找 通话接通次数
func (obj *_PeerData202201Mgr) GetBatchFromCallConnected(callConnecteds []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`call_connected` IN (?)", callConnecteds).Find(&results).Error

	return
}

// GetFromFreeConnectRate 通过free_connect_rate获取内容 通话接通率
func (obj *_PeerData202201Mgr) GetFromFreeConnectRate(freeConnectRate float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`free_connect_rate` = ?", freeConnectRate).Find(&results).Error

	return
}

// GetBatchFromFreeConnectRate 批量查找 通话接通率
func (obj *_PeerData202201Mgr) GetBatchFromFreeConnectRate(freeConnectRates []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`free_connect_rate` IN (?)", freeConnectRates).Find(&results).Error

	return
}

// GetFromPeerSend 通过peer_send获取内容 通话唤起次数
func (obj *_PeerData202201Mgr) GetFromPeerSend(peerSend int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_send` = ?", peerSend).Find(&results).Error

	return
}

// GetBatchFromPeerSend 批量查找 通话唤起次数
func (obj *_PeerData202201Mgr) GetBatchFromPeerSend(peerSends []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_send` IN (?)", peerSends).Find(&results).Error

	return
}

// GetFromPeerReply 通过peer_reply获取内容 通话被叫次数
func (obj *_PeerData202201Mgr) GetFromPeerReply(peerReply int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_reply` = ?", peerReply).Find(&results).Error

	return
}

// GetBatchFromPeerReply 批量查找 通话被叫次数
func (obj *_PeerData202201Mgr) GetBatchFromPeerReply(peerReplys []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_reply` IN (?)", peerReplys).Find(&results).Error

	return
}

// GetFromPeerConnected 通过peer_connected获取内容 通话接通次数
func (obj *_PeerData202201Mgr) GetFromPeerConnected(peerConnected int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_connected` = ?", peerConnected).Find(&results).Error

	return
}

// GetBatchFromPeerConnected 批量查找 通话接通次数
func (obj *_PeerData202201Mgr) GetBatchFromPeerConnected(peerConnecteds []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_connected` IN (?)", peerConnecteds).Find(&results).Error

	return
}

// GetFromPeerConnectRate 通过peer_connect_rate获取内容 通话接通率
func (obj *_PeerData202201Mgr) GetFromPeerConnectRate(peerConnectRate float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_connect_rate` = ?", peerConnectRate).Find(&results).Error

	return
}

// GetBatchFromPeerConnectRate 批量查找 通话接通率
func (obj *_PeerData202201Mgr) GetBatchFromPeerConnectRate(peerConnectRates []float64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`peer_connect_rate` IN (?)", peerConnectRates).Find(&results).Error

	return
}

// GetFromCmsCount 通过cms_count获取内容 短信发送次数
func (obj *_PeerData202201Mgr) GetFromCmsCount(cmsCount int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`cms_count` = ?", cmsCount).Find(&results).Error

	return
}

// GetBatchFromCmsCount 批量查找 短信发送次数
func (obj *_PeerData202201Mgr) GetBatchFromCmsCount(cmsCounts []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`cms_count` IN (?)", cmsCounts).Find(&results).Error

	return
}

// GetFromCreateDate 通过create_date获取内容 创建日期，格式：20060102
func (obj *_PeerData202201Mgr) GetFromCreateDate(createDate int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`create_date` = ?", createDate).Find(&results).Error

	return
}

// GetBatchFromCreateDate 批量查找 创建日期，格式：20060102
func (obj *_PeerData202201Mgr) GetBatchFromCreateDate(createDates []int) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`create_date` IN (?)", createDates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PeerData202201Mgr) FetchByPrimaryKey(id uint32) (result PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByPeerData202201UniqueUIDCreatedate primary or index 获取唯一内容
func (obj *_PeerData202201Mgr) FetchUniqueIndexByPeerData202201UniqueUIDCreatedate(uid int64, createDate int) (result PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`uid` = ? AND `create_date` = ?", uid, createDate).First(&result).Error

	return
}

// FetchIndexByPeerData202201IndexGuildID  获取多个内容
func (obj *_PeerData202201Mgr) FetchIndexByPeerData202201IndexGuildID(guildID int64) (results []*PeerData202201, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PeerData202201{}).Where("`guild_id` = ?", guildID).Find(&results).Error

	return
}
