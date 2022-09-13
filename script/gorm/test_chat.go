package gorm

// RelationBlocklist [...]
type RelationBlocklist struct {
	UID        int64 `gorm:"column:uid;type:bigint(20);not null"`
	InvitedUID int64 `gorm:"column:invited_uid;type:bigint(20);not null;comment:'拉黑uid'"`       // 拉黑uid
	CreateTime int   `gorm:"column:create_time;type:int(11);not null;comment:'创建时间'"`           // 创建时间
	DeleteTime int   `gorm:"column:delete_time;type:int(11);not null;default:0;comment:'删除时间'"` // 删除时间
}

// RelationBlocklistColumns get sql column name.获取数据库列名
var RelationBlocklistColumns = struct {
	UID        string
	InvitedUID string
	CreateTime string
	DeleteTime string
}{
	UID:        "uid",
	InvitedUID: "invited_uid",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
}
