package gorm

// PeerData202201 peer端数据表(按10周一张表)
type PeerData202201 struct {
	ID                       uint32  `gorm:"primaryKey;column:id;type:int(10) unsigned;not null"`                                        // ID
	UID                      int64   `gorm:"uniqueIndex:peer_data202201_unique_uid_createdate;column:uid;type:bigint(20);not null"`      // uid
	Level                    int8    `gorm:"column:level;type:tinyint(10);not null;default:0"`                                           // 主播等级
	GuildID                  int64   `gorm:"index:peer_data202201_index_guild_id;column:guild_id;type:bigint(20);not null;default:0"`    // 工会id
	GoldBalance              float64 `gorm:"column:gold_balance;type:double;not null;default:0"`                                         // 当日金币余额
	GoldIncome               float64 `gorm:"column:gold_income;type:double;not null;default:0"`                                          // 金币收益
	GoldIncomeForCall        float64 `gorm:"column:gold_income_for_call;type:double;not null;default:0"`                                 // 金币收益 1v1视频通话
	GoldIncomeForGift        float64 `gorm:"column:gold_income_for_gift;type:double;not null;default:0"`                                 // 金币收益 礼物
	GoldIncomeForReward      float64 `gorm:"column:gold_income_for_reward;type:double;not null;default:0"`                               // 金币收益 奖励
	GoldIncomeForCallTimes   int     `gorm:"column:gold_income_for_call_times;type:int(10);not null;default:0"`                          // 金币收益 1v1视频通话次数
	GoldIncomeForGiftTimes   int     `gorm:"column:gold_income_for_gift_times;type:int(10);not null;default:0"`                          // 金币收益 礼物次数
	GoldIncomeForRewardTimes int     `gorm:"column:gold_income_for_reward_times;type:int(10);not null;default:0"`                        // 金币收益 奖励次数
	OnlineCount              int     `gorm:"column:online_count;type:int(10);not null;default:0"`                                        // 在线时长
	CallCount                int     `gorm:"column:call_count;type:int(10);not null;default:0"`                                          // 通话时长
	VideoCount               int     `gorm:"column:video_count;type:int(10);not null;default:0"`                                         // 1v1视频通话时长
	PeerCount                int     `gorm:"column:peer_count;type:int(10);not null;default:0"`                                          // 匹配通话时长
	FreeRate                 float64 `gorm:"column:free_rate;type:double;not null;default:0"`                                            // 空闲率
	CallArouse               int     `gorm:"column:call_arouse;type:int(10);not null;default:0"`                                         // 通话唤起次数
	CallSend                 int     `gorm:"column:call_send;type:int(10);not null;default:0"`                                           // 通话被叫次数
	CallConnected            int     `gorm:"column:call_connected;type:int(10);not null;default:0"`                                      // 通话接通次数
	FreeConnectRate          float64 `gorm:"column:free_connect_rate;type:double;not null;default:0"`                                    // 通话接通率
	PeerSend                 int     `gorm:"column:peer_send;type:int(10);not null;default:0"`                                           // 通话唤起次数
	PeerReply                int     `gorm:"column:peer_reply;type:int(10);not null;default:0"`                                          // 通话被叫次数
	PeerConnected            int     `gorm:"column:peer_connected;type:int(10);not null;default:0"`                                      // 通话接通次数
	PeerConnectRate          float64 `gorm:"column:peer_connect_rate;type:double;not null;default:0"`                                    // 通话接通率
	CmsCount                 int     `gorm:"column:cms_count;type:int(10);not null;default:0"`                                           // 短信发送次数
	CreateDate               int     `gorm:"uniqueIndex:peer_data202201_unique_uid_createdate;column:create_date;type:int(11);not null"` // 创建日期，格式：20060102
}

// PeerData202201Columns get sql column name.获取数据库列名
var PeerData202201Columns = struct {
	ID                       string
	UID                      string
	Level                    string
	GuildID                  string
	GoldBalance              string
	GoldIncome               string
	GoldIncomeForCall        string
	GoldIncomeForGift        string
	GoldIncomeForReward      string
	GoldIncomeForCallTimes   string
	GoldIncomeForGiftTimes   string
	GoldIncomeForRewardTimes string
	OnlineCount              string
	CallCount                string
	VideoCount               string
	PeerCount                string
	FreeRate                 string
	CallArouse               string
	CallSend                 string
	CallConnected            string
	FreeConnectRate          string
	PeerSend                 string
	PeerReply                string
	PeerConnected            string
	PeerConnectRate          string
	CmsCount                 string
	CreateDate               string
}{
	ID:                       "id",
	UID:                      "uid",
	Level:                    "level",
	GuildID:                  "guild_id",
	GoldBalance:              "gold_balance",
	GoldIncome:               "gold_income",
	GoldIncomeForCall:        "gold_income_for_call",
	GoldIncomeForGift:        "gold_income_for_gift",
	GoldIncomeForReward:      "gold_income_for_reward",
	GoldIncomeForCallTimes:   "gold_income_for_call_times",
	GoldIncomeForGiftTimes:   "gold_income_for_gift_times",
	GoldIncomeForRewardTimes: "gold_income_for_reward_times",
	OnlineCount:              "online_count",
	CallCount:                "call_count",
	VideoCount:               "video_count",
	PeerCount:                "peer_count",
	FreeRate:                 "free_rate",
	CallArouse:               "call_arouse",
	CallSend:                 "call_send",
	CallConnected:            "call_connected",
	FreeConnectRate:          "free_connect_rate",
	PeerSend:                 "peer_send",
	PeerReply:                "peer_reply",
	PeerConnected:            "peer_connected",
	PeerConnectRate:          "peer_connect_rate",
	CmsCount:                 "cms_count",
	CreateDate:               "create_date",
}
