package model

// CyymallUser 用户
type CyymallUser struct {
	ID                        uint64 `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Type                      int    `gorm:"column:type;type:smallint(6);default:1;NOT NULL" json:"type"` // 用户类型：0=管理员，1=普通用户3员工账号
	Username                  string `gorm:"column:username;type:varchar(255);NOT NULL" json:"username"`
	Password                  string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`
	AuthKey                   string `gorm:"column:auth_key;type:varchar(255);NOT NULL" json:"authKey"`
	AccessToken               string `gorm:"column:access_token;type:varchar(255);NOT NULL" json:"accessToken"`
	Addtime                   int    `gorm:"column:addtime;type:int(11);default:0;NOT NULL" json:"addtime"`
	IsDelete                  int    `gorm:"column:is_delete;type:smallint(6);default:0;NOT NULL" json:"isDelete"`
	WechatOpenID              string `gorm:"column:wechat_open_id;type:varchar(255);NOT NULL" json:"wechatOpenId"`   // 微信openid
	WechatUnionID             string `gorm:"column:wechat_union_id;type:varchar(255);NOT NULL" json:"wechatUnionId"` // 微信用户union id
	Nickname                  string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                      // 昵称
	AvatarUrl                 string `gorm:"column:avatar_url;type:varchar(2048)" json:"avatarUrl"`
	StoreID                   int    `gorm:"column:store_id;type:int(11);default:0;NOT NULL" json:"storeId"`                // 商城id
	IsDistributor             int    `gorm:"column:is_distributor;type:int(11);default:0;NOT NULL" json:"isDistributor"`    // 是否是分销商 0--不是 1--是 2--申请中
	ParentID                  int    `gorm:"column:parent_id;type:int(11);default:0;NOT NULL" json:"parentId"`              // 父级ID
	Time                      int    `gorm:"column:time;type:int(11);default:0;NOT NULL" json:"time"`                       // 成为分销商的时间
	TotalPrice                string `gorm:"column:total_price;type:decimal(10,2);default:0.00;NOT NULL" json:"totalPrice"` // 累计佣金
	Price                     string `gorm:"column:price;type:decimal(10,2);default:0.00;NOT NULL" json:"price"`            // 可提现佣金
	IsClerk                   int    `gorm:"column:is_clerk;type:int(11);default:0;NOT NULL" json:"isClerk"`                // 是否是核销员 0--不是 1--是
	We7Uid                    int    `gorm:"column:we7_uid;type:int(11);default:0;NOT NULL" json:"we7Uid"`                  // 微擎账户id
	ShopID                    int    `gorm:"column:shop_id;type:int(11)" json:"shopId"`
	Level                     int    `gorm:"column:level;type:int(11);default:-1" json:"level"`                                    // 会员等级
	Integral                  uint   `gorm:"column:integral;type:int(11) unsigned;default:0;NOT NULL" json:"integral"`             // 用户当前积分
	TotalIntegral             uint   `gorm:"column:total_integral;type:int(11) unsigned;default:0;NOT NULL" json:"totalIntegral"`  // 用户总获得积分
	Money                     string `gorm:"column:money;type:decimal(10,2);default:0.00" json:"money"`                            // 余额
	ContactWay                string `gorm:"column:contact_way;type:varchar(255)" json:"contactWay"`                               // 联系方式
	Comments                  string `gorm:"column:comments;type:varchar(255)" json:"comments"`                                    // 备注
	Binding                   string `gorm:"column:binding;type:char(11)" json:"binding"`                                          // 授权手机号
	WechatPlatformOpenID      string `gorm:"column:wechat_platform_open_id;type:varchar(64);NOT NULL" json:"wechatPlatformOpenId"` // 微信公众号openid
	Platform                  int    `gorm:"column:platform;type:tinyint(4);default:0;NOT NULL" json:"platform"`                   // 小程序平台 0 => 微信, 1 => 支付宝
	Blacklist                 int    `gorm:"column:blacklist;type:tinyint(1);default:0;NOT NULL" json:"blacklist"`                 // 黑名单 0.否 | 1.是
	ParentUserID              int    `gorm:"column:parent_user_id;type:int(11);default:0;NOT NULL" json:"parentUserId"`            // 可能成为上级的ID
	Appcid                    string `gorm:"column:appcid;type:varchar(100)" json:"appcid"`
	IsAdmin                   uint   `gorm:"column:is_admin;type:tinyint(4) unsigned;default:0" json:"isAdmin"`
	TuanPrice                 string `gorm:"column:tuan_price;type:decimal(11,2);default:0.00" json:"tuanPrice"`
	IsDelivery                int    `gorm:"column:is_delivery;type:tinyint(1);default:0" json:"isDelivery"` // 1是配送员0不是
	MchID                     int    `gorm:"column:mch_id;type:int(11)" json:"mchId"`
	RealName                  string `gorm:"column:real_name;type:varchar(255)" json:"realName"`                                // 真实姓名
	RealCode                  string `gorm:"column:real_code;type:varchar(255)" json:"realCode"`                                // 身份证号
	RealJustPic               string `gorm:"column:real_just_pic;type:varchar(255)" json:"realJustPic"`                         // 身份证正面照
	RealBackPic               string `gorm:"column:real_back_pic;type:varchar(255)" json:"realBackPic"`                         // 身份证反面照
	IsReal                    int    `gorm:"column:is_real;type:tinyint(1);default:0" json:"isReal"`                            // 是否已实名注册 2审核中3审核未通过1审核通过
	LivePrice                 string `gorm:"column:live_price;type:decimal(10,2);default:0.00;NOT NULL" json:"livePrice"`       // 直播佣金
	Fyjine                    string `gorm:"column:fyjine;type:varchar(255)" json:"fyjine"`                                     // 待返总金额
	Yifan                     string `gorm:"column:yifan;type:varchar(255)" json:"yifan"`                                       // 已返金额
	DeliveryOffice            int    `gorm:"column:delivery_office;type:tinyint(1);default:0" json:"deliveryOffice"`            // 骑手是否在线（0、在线，1离线）
	SubsidyPrice              string `gorm:"column:subsidy_price;type:decimal(10,2);default:0.00;NOT NULL" json:"subsidyPrice"` // 补贴金额
	DeliveryMchID             int    `gorm:"column:delivery_mch_id;type:int(11);default:0" json:"deliveryMchId"`
	ClerkMchID                uint   `gorm:"column:clerk_mch_id;type:int(11) unsigned;default:0" json:"clerkMchId"`                                // 核销员入驻商id
	IsGroupCentre             int    `gorm:"column:is_group_centre;type:int(11);default:0" json:"isGroupCentre"`                                   // 是否团长身份 0--否，1--是
	Sex                       int    `gorm:"column:sex;type:int(11);default:0" json:"sex"`                                                         // 性别0-未知,1男,2女
	BirthDate                 string `gorm:"column:birth_date;type:varchar(25)" json:"birthDate"`                                                  // 出生日期
	IsVgoods                  int    `gorm:"column:is_vgoods;type:int(11);default:1" json:"isVgoods"`                                              // 是否允许发布抖品：0=否，1=是
	ShareCatID                int    `gorm:"column:share_cat_id;type:int(11);default:0;NOT NULL" json:"shareCatId"`                                // 分销商分组id
	UserLabel                 int    `gorm:"column:user_label;type:int(11);default:0" json:"userLabel"`                                            // 用户标签
	ConsumeRebateTotalFyjine  string `gorm:"column:consume_rebate_total_fyjine;type:decimal(10,2);default:0.00" json:"consumeRebateTotalFyjine"`   // 消费返利总金额
	ConsumeRebateFyjine       string `gorm:"column:consume_rebate_fyjine;type:decimal(10,2);default:0.00" json:"consumeRebateFyjine"`              // 消费返利未返金额
	ConsumeRebateTotalFyjifen string `gorm:"column:consume_rebate_total_fyjifen;type:decimal(10,2);default:0.00" json:"consumeRebateTotalFyjifen"` // 消费返利总积分
	ConsumeRebateFyjifen      string `gorm:"column:consume_rebate_fyjifen;type:decimal(10,2);default:0.00" json:"consumeRebateFyjifen"`            // 消费返利未返积分
	PickMch                   int    `gorm:"column:pick_mch;type:int(11);default:0" json:"pickMch"`                                                // 当前选择店铺id（关联mch表的自增id）
	BalancePassword           string `gorm:"column:balance_password;type:varchar(255)" json:"balancePassword"`                                     // 余额支付密码
	IsFreePayment             int    `gorm:"column:is_free_payment;type:int(11);default:1" json:"isFreePayment"`                                   // 是否开启免密余额支付
	ProvinceID                int    `gorm:"column:province_id;type:int(11);default:0" json:"provinceId"`                                          // 来源省份
	Source                    int    `gorm:"column:source;type:tinyint(1);default:1" json:"source"`                                                // 1-小程序 2-公众号 3-app
	Subsidies                 string `gorm:"column:subsidies;type:decimal(10,2);default:0.00" json:"subsidies"`                                    // 补贴金
	ConsumeMoney              string `gorm:"column:consume_money;type:decimal(10,2);default:0.00" json:"consumeMoney"`                             // 用户累计消费金额
	IsMchShipping             int    `gorm:"column:is_mch_shipping;type:tinyint(4);default:0" json:"isMchShipping"`                                // 是否是配送员 0否 1是
}

// TableName table name
func (m *CyymallUser) TableName() string {
	return "cyymall_user"
}
