package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateCyymallUserRequest request params
type CreateCyymallUserRequest struct {
	Type                      int    `json:"type" binding:""` // 用户类型：0=管理员，1=普通用户3员工账号
	Username                  string `json:"username" binding:""`
	Password                  string `json:"password" binding:""`
	AuthKey                   string `json:"authKey" binding:""`
	AccessToken               string `json:"accessToken" binding:""`
	Addtime                   int    `json:"addtime" binding:""`
	IsDelete                  int    `json:"isDelete" binding:""`
	WechatOpenID              string `json:"wechatOpenId" binding:""`  // 微信openid
	WechatUnionID             string `json:"wechatUnionId" binding:""` // 微信用户union id
	Nickname                  string `json:"nickname" binding:""`      // 昵称
	AvatarUrl                 string `json:"avatarUrl" binding:""`
	StoreID                   int    `json:"storeId" binding:""`       // 商城id
	IsDistributor             int    `json:"isDistributor" binding:""` // 是否是分销商 0--不是 1--是 2--申请中
	ParentID                  int    `json:"parentId" binding:""`      // 父级ID
	Time                      int    `json:"time" binding:""`          // 成为分销商的时间
	TotalPrice                string `json:"totalPrice" binding:""`    // 累计佣金
	Price                     string `json:"price" binding:""`         // 可提现佣金
	IsClerk                   int    `json:"isClerk" binding:""`       // 是否是核销员 0--不是 1--是
	We7Uid                    int    `json:"we7Uid" binding:""`        // 微擎账户id
	ShopID                    int    `json:"shopId" binding:""`
	Level                     int    `json:"level" binding:""`                // 会员等级
	Integral                  uint   `json:"integral" binding:""`             // 用户当前积分
	TotalIntegral             uint   `json:"totalIntegral" binding:""`        // 用户总获得积分
	Money                     string `json:"money" binding:""`                // 余额
	ContactWay                string `json:"contactWay" binding:""`           // 联系方式
	Comments                  string `json:"comments" binding:""`             // 备注
	Binding                   string `json:"binding" binding:""`              // 授权手机号
	WechatPlatformOpenID      string `json:"wechatPlatformOpenId" binding:""` // 微信公众号openid
	Platform                  int    `json:"platform" binding:""`             // 小程序平台 0 => 微信, 1 => 支付宝
	Blacklist                 int    `json:"blacklist" binding:""`            // 黑名单 0.否 | 1.是
	ParentUserID              int    `json:"parentUserId" binding:""`         // 可能成为上级的ID
	Appcid                    string `json:"appcid" binding:""`
	IsAdmin                   uint   `json:"isAdmin" binding:""`
	TuanPrice                 string `json:"tuanPrice" binding:""`
	IsDelivery                int    `json:"isDelivery" binding:""` // 1是配送员0不是
	MchID                     int    `json:"mchId" binding:""`
	RealName                  string `json:"realName" binding:""`       // 真实姓名
	RealCode                  string `json:"realCode" binding:""`       // 身份证号
	RealJustPic               string `json:"realJustPic" binding:""`    // 身份证正面照
	RealBackPic               string `json:"realBackPic" binding:""`    // 身份证反面照
	IsReal                    int    `json:"isReal" binding:""`         // 是否已实名注册 2审核中3审核未通过1审核通过
	LivePrice                 string `json:"livePrice" binding:""`      // 直播佣金
	Fyjine                    string `json:"fyjine" binding:""`         // 待返总金额
	Yifan                     string `json:"yifan" binding:""`          // 已返金额
	DeliveryOffice            int    `json:"deliveryOffice" binding:""` // 骑手是否在线（0、在线，1离线）
	SubsidyPrice              string `json:"subsidyPrice" binding:""`   // 补贴金额
	DeliveryMchID             int    `json:"deliveryMchId" binding:""`
	ClerkMchID                uint   `json:"clerkMchId" binding:""`                // 核销员入驻商id
	IsGroupCentre             int    `json:"isGroupCentre" binding:""`             // 是否团长身份 0--否，1--是
	Sex                       int    `json:"sex" binding:""`                       // 性别0-未知,1男,2女
	BirthDate                 string `json:"birthDate" binding:""`                 // 出生日期
	IsVgoods                  int    `json:"isVgoods" binding:""`                  // 是否允许发布抖品：0=否，1=是
	ShareCatID                int    `json:"shareCatId" binding:""`                // 分销商分组id
	UserLabel                 int    `json:"userLabel" binding:""`                 // 用户标签
	ConsumeRebateTotalFyjine  string `json:"consumeRebateTotalFyjine" binding:""`  // 消费返利总金额
	ConsumeRebateFyjine       string `json:"consumeRebateFyjine" binding:""`       // 消费返利未返金额
	ConsumeRebateTotalFyjifen string `json:"consumeRebateTotalFyjifen" binding:""` // 消费返利总积分
	ConsumeRebateFyjifen      string `json:"consumeRebateFyjifen" binding:""`      // 消费返利未返积分
	PickMch                   int    `json:"pickMch" binding:""`                   // 当前选择店铺id（关联mch表的自增id）
	BalancePassword           string `json:"balancePassword" binding:""`           // 余额支付密码
	IsFreePayment             int    `json:"isFreePayment" binding:""`             // 是否开启免密余额支付
	ProvinceID                int    `json:"provinceId" binding:""`                // 来源省份
	Source                    int    `json:"source" binding:""`                    // 1-小程序 2-公众号 3-app
	Subsidies                 string `json:"subsidies" binding:""`                 // 补贴金
	ConsumeMoney              string `json:"consumeMoney" binding:""`              // 用户累计消费金额
	IsMchShipping             int    `json:"isMchShipping" binding:""`             // 是否是配送员 0否 1是
}

// UpdateCyymallUserByIDRequest request params
type UpdateCyymallUserByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	Type                      int    `json:"type" binding:""` // 用户类型：0=管理员，1=普通用户3员工账号
	Username                  string `json:"username" binding:""`
	Password                  string `json:"password" binding:""`
	AuthKey                   string `json:"authKey" binding:""`
	AccessToken               string `json:"accessToken" binding:""`
	Addtime                   int    `json:"addtime" binding:""`
	IsDelete                  int    `json:"isDelete" binding:""`
	WechatOpenID              string `json:"wechatOpenId" binding:""`  // 微信openid
	WechatUnionID             string `json:"wechatUnionId" binding:""` // 微信用户union id
	Nickname                  string `json:"nickname" binding:""`      // 昵称
	AvatarUrl                 string `json:"avatarUrl" binding:""`
	StoreID                   int    `json:"storeId" binding:""`       // 商城id
	IsDistributor             int    `json:"isDistributor" binding:""` // 是否是分销商 0--不是 1--是 2--申请中
	ParentID                  int    `json:"parentId" binding:""`      // 父级ID
	Time                      int    `json:"time" binding:""`          // 成为分销商的时间
	TotalPrice                string `json:"totalPrice" binding:""`    // 累计佣金
	Price                     string `json:"price" binding:""`         // 可提现佣金
	IsClerk                   int    `json:"isClerk" binding:""`       // 是否是核销员 0--不是 1--是
	We7Uid                    int    `json:"we7Uid" binding:""`        // 微擎账户id
	ShopID                    int    `json:"shopId" binding:""`
	Level                     int    `json:"level" binding:""`                // 会员等级
	Integral                  uint   `json:"integral" binding:""`             // 用户当前积分
	TotalIntegral             uint   `json:"totalIntegral" binding:""`        // 用户总获得积分
	Money                     string `json:"money" binding:""`                // 余额
	ContactWay                string `json:"contactWay" binding:""`           // 联系方式
	Comments                  string `json:"comments" binding:""`             // 备注
	Binding                   string `json:"binding" binding:""`              // 授权手机号
	WechatPlatformOpenID      string `json:"wechatPlatformOpenId" binding:""` // 微信公众号openid
	Platform                  int    `json:"platform" binding:""`             // 小程序平台 0 => 微信, 1 => 支付宝
	Blacklist                 int    `json:"blacklist" binding:""`            // 黑名单 0.否 | 1.是
	ParentUserID              int    `json:"parentUserId" binding:""`         // 可能成为上级的ID
	Appcid                    string `json:"appcid" binding:""`
	IsAdmin                   uint   `json:"isAdmin" binding:""`
	TuanPrice                 string `json:"tuanPrice" binding:""`
	IsDelivery                int    `json:"isDelivery" binding:""` // 1是配送员0不是
	MchID                     int    `json:"mchId" binding:""`
	RealName                  string `json:"realName" binding:""`       // 真实姓名
	RealCode                  string `json:"realCode" binding:""`       // 身份证号
	RealJustPic               string `json:"realJustPic" binding:""`    // 身份证正面照
	RealBackPic               string `json:"realBackPic" binding:""`    // 身份证反面照
	IsReal                    int    `json:"isReal" binding:""`         // 是否已实名注册 2审核中3审核未通过1审核通过
	LivePrice                 string `json:"livePrice" binding:""`      // 直播佣金
	Fyjine                    string `json:"fyjine" binding:""`         // 待返总金额
	Yifan                     string `json:"yifan" binding:""`          // 已返金额
	DeliveryOffice            int    `json:"deliveryOffice" binding:""` // 骑手是否在线（0、在线，1离线）
	SubsidyPrice              string `json:"subsidyPrice" binding:""`   // 补贴金额
	DeliveryMchID             int    `json:"deliveryMchId" binding:""`
	ClerkMchID                uint   `json:"clerkMchId" binding:""`                // 核销员入驻商id
	IsGroupCentre             int    `json:"isGroupCentre" binding:""`             // 是否团长身份 0--否，1--是
	Sex                       int    `json:"sex" binding:""`                       // 性别0-未知,1男,2女
	BirthDate                 string `json:"birthDate" binding:""`                 // 出生日期
	IsVgoods                  int    `json:"isVgoods" binding:""`                  // 是否允许发布抖品：0=否，1=是
	ShareCatID                int    `json:"shareCatId" binding:""`                // 分销商分组id
	UserLabel                 int    `json:"userLabel" binding:""`                 // 用户标签
	ConsumeRebateTotalFyjine  string `json:"consumeRebateTotalFyjine" binding:""`  // 消费返利总金额
	ConsumeRebateFyjine       string `json:"consumeRebateFyjine" binding:""`       // 消费返利未返金额
	ConsumeRebateTotalFyjifen string `json:"consumeRebateTotalFyjifen" binding:""` // 消费返利总积分
	ConsumeRebateFyjifen      string `json:"consumeRebateFyjifen" binding:""`      // 消费返利未返积分
	PickMch                   int    `json:"pickMch" binding:""`                   // 当前选择店铺id（关联mch表的自增id）
	BalancePassword           string `json:"balancePassword" binding:""`           // 余额支付密码
	IsFreePayment             int    `json:"isFreePayment" binding:""`             // 是否开启免密余额支付
	ProvinceID                int    `json:"provinceId" binding:""`                // 来源省份
	Source                    int    `json:"source" binding:""`                    // 1-小程序 2-公众号 3-app
	Subsidies                 string `json:"subsidies" binding:""`                 // 补贴金
	ConsumeMoney              string `json:"consumeMoney" binding:""`              // 用户累计消费金额
	IsMchShipping             int    `json:"isMchShipping" binding:""`             // 是否是配送员 0否 1是
}

// CyymallUserObjDetail detail
type CyymallUserObjDetail struct {
	ID string `json:"id"` // convert to string id

	Type                      int    `json:"type"` // 用户类型：0=管理员，1=普通用户3员工账号
	Username                  string `json:"username"`
	Password                  string `json:"password"`
	AuthKey                   string `json:"authKey"`
	AccessToken               string `json:"accessToken"`
	Addtime                   int    `json:"addtime"`
	IsDelete                  int    `json:"isDelete"`
	WechatOpenID              string `json:"wechatOpenId"`  // 微信openid
	WechatUnionID             string `json:"wechatUnionId"` // 微信用户union id
	Nickname                  string `json:"nickname"`      // 昵称
	AvatarUrl                 string `json:"avatarUrl"`
	StoreID                   int    `json:"storeId"`       // 商城id
	IsDistributor             int    `json:"isDistributor"` // 是否是分销商 0--不是 1--是 2--申请中
	ParentID                  int    `json:"parentId"`      // 父级ID
	Time                      int    `json:"time"`          // 成为分销商的时间
	TotalPrice                string `json:"totalPrice"`    // 累计佣金
	Price                     string `json:"price"`         // 可提现佣金
	IsClerk                   int    `json:"isClerk"`       // 是否是核销员 0--不是 1--是
	We7Uid                    int    `json:"we7Uid"`        // 微擎账户id
	ShopID                    int    `json:"shopId"`
	Level                     int    `json:"level"`                // 会员等级
	Integral                  uint   `json:"integral"`             // 用户当前积分
	TotalIntegral             uint   `json:"totalIntegral"`        // 用户总获得积分
	Money                     string `json:"money"`                // 余额
	ContactWay                string `json:"contactWay"`           // 联系方式
	Comments                  string `json:"comments"`             // 备注
	Binding                   string `json:"binding"`              // 授权手机号
	WechatPlatformOpenID      string `json:"wechatPlatformOpenId"` // 微信公众号openid
	Platform                  int    `json:"platform"`             // 小程序平台 0 => 微信, 1 => 支付宝
	Blacklist                 int    `json:"blacklist"`            // 黑名单 0.否 | 1.是
	ParentUserID              int    `json:"parentUserId"`         // 可能成为上级的ID
	Appcid                    string `json:"appcid"`
	IsAdmin                   uint   `json:"isAdmin"`
	TuanPrice                 string `json:"tuanPrice"`
	IsDelivery                int    `json:"isDelivery"` // 1是配送员0不是
	MchID                     int    `json:"mchId"`
	RealName                  string `json:"realName"`       // 真实姓名
	RealCode                  string `json:"realCode"`       // 身份证号
	RealJustPic               string `json:"realJustPic"`    // 身份证正面照
	RealBackPic               string `json:"realBackPic"`    // 身份证反面照
	IsReal                    int    `json:"isReal"`         // 是否已实名注册 2审核中3审核未通过1审核通过
	LivePrice                 string `json:"livePrice"`      // 直播佣金
	Fyjine                    string `json:"fyjine"`         // 待返总金额
	Yifan                     string `json:"yifan"`          // 已返金额
	DeliveryOffice            int    `json:"deliveryOffice"` // 骑手是否在线（0、在线，1离线）
	SubsidyPrice              string `json:"subsidyPrice"`   // 补贴金额
	DeliveryMchID             int    `json:"deliveryMchId"`
	ClerkMchID                uint   `json:"clerkMchId"`                // 核销员入驻商id
	IsGroupCentre             int    `json:"isGroupCentre"`             // 是否团长身份 0--否，1--是
	Sex                       int    `json:"sex"`                       // 性别0-未知,1男,2女
	BirthDate                 string `json:"birthDate"`                 // 出生日期
	IsVgoods                  int    `json:"isVgoods"`                  // 是否允许发布抖品：0=否，1=是
	ShareCatID                int    `json:"shareCatId"`                // 分销商分组id
	UserLabel                 int    `json:"userLabel"`                 // 用户标签
	ConsumeRebateTotalFyjine  string `json:"consumeRebateTotalFyjine"`  // 消费返利总金额
	ConsumeRebateFyjine       string `json:"consumeRebateFyjine"`       // 消费返利未返金额
	ConsumeRebateTotalFyjifen string `json:"consumeRebateTotalFyjifen"` // 消费返利总积分
	ConsumeRebateFyjifen      string `json:"consumeRebateFyjifen"`      // 消费返利未返积分
	PickMch                   int    `json:"pickMch"`                   // 当前选择店铺id（关联mch表的自增id）
	BalancePassword           string `json:"balancePassword"`           // 余额支付密码
	IsFreePayment             int    `json:"isFreePayment"`             // 是否开启免密余额支付
	ProvinceID                int    `json:"provinceId"`                // 来源省份
	Source                    int    `json:"source"`                    // 1-小程序 2-公众号 3-app
	Subsidies                 string `json:"subsidies"`                 // 补贴金
	ConsumeMoney              string `json:"consumeMoney"`              // 用户累计消费金额
	IsMchShipping             int    `json:"isMchShipping"`             // 是否是配送员 0否 1是
}

// CreateCyymallUserRespond only for api docs
type CreateCyymallUserRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateCyymallUserByIDRespond only for api docs
type UpdateCyymallUserByIDRespond struct {
	Result
}

// GetCyymallUserByIDRespond only for api docs
type GetCyymallUserByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		CyymallUser CyymallUserObjDetail `json:"cyymallUser"`
	} `json:"data"` // return data
}

// DeleteCyymallUserByIDRespond only for api docs
type DeleteCyymallUserByIDRespond struct {
	Result
}

// DeleteCyymallUsersByIDsRequest request params
type DeleteCyymallUsersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteCyymallUsersByIDsRespond only for api docs
type DeleteCyymallUsersByIDsRespond struct {
	Result
}

// GetCyymallUserByConditionRequest request params
type GetCyymallUserByConditionRequest struct {
	query.Conditions
}

// GetCyymallUserByConditionRespond only for api docs
type GetCyymallUserByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		CyymallUser CyymallUserObjDetail `json:"cyymallUser"`
	} `json:"data"` // return data
}

// ListCyymallUsersByIDsRequest request params
type ListCyymallUsersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListCyymallUsersByIDsRespond only for api docs
type ListCyymallUsersByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		CyymallUsers []CyymallUserObjDetail `json:"cyymallUsers"`
	} `json:"data"` // return data
}

// ListCyymallUsersRequest request params
type ListCyymallUsersRequest struct {
	query.Params
}

// ListCyymallUsersRespond only for api docs
type ListCyymallUsersRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		CyymallUsers []CyymallUserObjDetail `json:"cyymallUsers"`
	} `json:"data"` // return data
}
