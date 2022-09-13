package gorm

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RelationBlocklistMgr struct {
	*_BaseMgr
}

// RelationBlocklistMgr open func
func RelationBlocklistMgr(db *gorm.DB) *_RelationBlocklistMgr {
	if db == nil {
		panic(fmt.Errorf("RelationBlocklistMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RelationBlocklistMgr{_BaseMgr: &_BaseMgr{DB: db.Table("relation_blocklist"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RelationBlocklistMgr) GetTableName() string {
	return "relation_blocklist"
}

// Reset 重置gorm会话
func (obj *_RelationBlocklistMgr) Reset() *_RelationBlocklistMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RelationBlocklistMgr) Get() (result RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RelationBlocklistMgr) Gets() (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RelationBlocklistMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUID uid获取
func (obj *_RelationBlocklistMgr) WithUID(uid int64) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithInvitedUID invited_uid获取 拉黑uid
func (obj *_RelationBlocklistMgr) WithInvitedUID(invitedUID int64) Option {
	return optionFunc(func(o *options) { o.query["invited_uid"] = invitedUID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_RelationBlocklistMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_RelationBlocklistMgr) WithDeleteTime(deleteTime int) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// GetByOption 功能选项模式获取
func (obj *_RelationBlocklistMgr) GetByOption(opts ...Option) (result RelationBlocklist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RelationBlocklistMgr) GetByOptions(opts ...Option) (results []*RelationBlocklist, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUID 通过uid获取内容
func (obj *_RelationBlocklistMgr) GetFromUID(uid int64) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找
func (obj *_RelationBlocklistMgr) GetBatchFromUID(uids []int64) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromInvitedUID 通过invited_uid获取内容 拉黑uid
func (obj *_RelationBlocklistMgr) GetFromInvitedUID(invitedUID int64) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`invited_uid` = ?", invitedUID).Find(&results).Error

	return
}

// GetBatchFromInvitedUID 批量查找 拉黑uid
func (obj *_RelationBlocklistMgr) GetBatchFromInvitedUID(invitedUIDs []int64) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`invited_uid` IN (?)", invitedUIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_RelationBlocklistMgr) GetFromCreateTime(createTime int) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_RelationBlocklistMgr) GetBatchFromCreateTime(createTimes []int) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_RelationBlocklistMgr) GetFromDeleteTime(deleteTime int) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`delete_time` = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量查找 删除时间
func (obj *_RelationBlocklistMgr) GetBatchFromDeleteTime(deleteTimes []int) (results []*RelationBlocklist, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RelationBlocklist{}).Where("`delete_time` IN (?)", deleteTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
