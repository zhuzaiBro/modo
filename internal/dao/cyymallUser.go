package dao

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	cacheBase "github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"cyy/internal/cache"
	"cyy/internal/model"
)

var _ CyymallUserDao = (*cyymallUserDao)(nil)

// CyymallUserDao defining the dao interface
type CyymallUserDao interface {
	Create(ctx context.Context, table *model.CyymallUser) error
	DeleteByID(ctx context.Context, id uint64) error
	DeleteByIDs(ctx context.Context, ids []uint64) error
	UpdateByID(ctx context.Context, table *model.CyymallUser) error
	GetByID(ctx context.Context, id uint64) (*model.CyymallUser, error)
	GetByCondition(ctx context.Context, condition *query.Conditions) (*model.CyymallUser, error)
	GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.CyymallUser, error)
	GetByLastID(ctx context.Context, lastID uint64, limit int, sort string) ([]*model.CyymallUser, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.CyymallUser, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.CyymallUser) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.CyymallUser) error
}

type cyymallUserDao struct {
	db    *gorm.DB
	cache cache.CyymallUserCache // if nil, the cache is not used.
	sfg   *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewCyymallUserDao creating the dao interface
func NewCyymallUserDao(db *gorm.DB, xCache cache.CyymallUserCache) CyymallUserDao {
	if xCache == nil {
		return &cyymallUserDao{db: db}
	}
	return &cyymallUserDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *cyymallUserDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create a record, insert the record and the id value is written back to the table
func (d *cyymallUserDao) Create(ctx context.Context, table *model.CyymallUser) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID delete a record by id
func (d *cyymallUserDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.CyymallUser{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// DeleteByIDs delete records by batch id
func (d *cyymallUserDao) DeleteByIDs(ctx context.Context, ids []uint64) error {
	err := d.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&model.CyymallUser{}).Error
	if err != nil {
		return err
	}

	// delete cache
	for _, id := range ids {
		_ = d.deleteCache(ctx, id)
	}

	return nil
}

// UpdateByID update a record by id
func (d *cyymallUserDao) UpdateByID(ctx context.Context, table *model.CyymallUser) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

func (d *cyymallUserDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.CyymallUser) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.Type != 0 {
		update["type"] = table.Type
	}
	if table.Username != "" {
		update["username"] = table.Username
	}
	if table.Password != "" {
		update["password"] = table.Password
	}
	if table.AuthKey != "" {
		update["auth_key"] = table.AuthKey
	}
	if table.AccessToken != "" {
		update["access_token"] = table.AccessToken
	}
	if table.Addtime != 0 {
		update["addtime"] = table.Addtime
	}
	if table.IsDelete != 0 {
		update["is_delete"] = table.IsDelete
	}
	if table.WechatOpenID != "" {
		update["wechat_open_id"] = table.WechatOpenID
	}
	if table.WechatUnionID != "" {
		update["wechat_union_id"] = table.WechatUnionID
	}
	if table.Nickname != "" {
		update["nickname"] = table.Nickname
	}
	if table.AvatarUrl != "" {
		update["avatar_url"] = table.AvatarUrl
	}
	if table.StoreID != 0 {
		update["store_id"] = table.StoreID
	}
	if table.IsDistributor != 0 {
		update["is_distributor"] = table.IsDistributor
	}
	if table.ParentID != 0 {
		update["parent_id"] = table.ParentID
	}
	if table.Time != 0 {
		update["time"] = table.Time
	}
	if table.TotalPrice != "" {
		update["total_price"] = table.TotalPrice
	}
	if table.Price != "" {
		update["price"] = table.Price
	}
	if table.IsClerk != 0 {
		update["is_clerk"] = table.IsClerk
	}
	if table.We7Uid != 0 {
		update["we7_uid"] = table.We7Uid
	}
	if table.ShopID != 0 {
		update["shop_id"] = table.ShopID
	}
	if table.Level != 0 {
		update["level"] = table.Level
	}
	if table.Integral != 0 {
		update["integral"] = table.Integral
	}
	if table.TotalIntegral != 0 {
		update["total_integral"] = table.TotalIntegral
	}
	if table.Money != "" {
		update["money"] = table.Money
	}
	if table.ContactWay != "" {
		update["contact_way"] = table.ContactWay
	}
	if table.Comments != "" {
		update["comments"] = table.Comments
	}
	if table.Binding != "" {
		update["binding"] = table.Binding
	}
	if table.WechatPlatformOpenID != "" {
		update["wechat_platform_open_id"] = table.WechatPlatformOpenID
	}
	if table.Platform != 0 {
		update["platform"] = table.Platform
	}
	if table.Blacklist != 0 {
		update["blacklist"] = table.Blacklist
	}
	if table.ParentUserID != 0 {
		update["parent_user_id"] = table.ParentUserID
	}
	if table.Appcid != "" {
		update["appcid"] = table.Appcid
	}
	if table.IsAdmin != 0 {
		update["is_admin"] = table.IsAdmin
	}
	if table.TuanPrice != "" {
		update["tuan_price"] = table.TuanPrice
	}
	if table.IsDelivery != 0 {
		update["is_delivery"] = table.IsDelivery
	}
	if table.MchID != 0 {
		update["mch_id"] = table.MchID
	}
	if table.RealName != "" {
		update["real_name"] = table.RealName
	}
	if table.RealCode != "" {
		update["real_code"] = table.RealCode
	}
	if table.RealJustPic != "" {
		update["real_just_pic"] = table.RealJustPic
	}
	if table.RealBackPic != "" {
		update["real_back_pic"] = table.RealBackPic
	}
	if table.IsReal != 0 {
		update["is_real"] = table.IsReal
	}
	if table.LivePrice != "" {
		update["live_price"] = table.LivePrice
	}
	if table.Fyjine != "" {
		update["fyjine"] = table.Fyjine
	}
	if table.Yifan != "" {
		update["yifan"] = table.Yifan
	}
	if table.DeliveryOffice != 0 {
		update["delivery_office"] = table.DeliveryOffice
	}
	if table.SubsidyPrice != "" {
		update["subsidy_price"] = table.SubsidyPrice
	}
	if table.DeliveryMchID != 0 {
		update["delivery_mch_id"] = table.DeliveryMchID
	}
	if table.ClerkMchID != 0 {
		update["clerk_mch_id"] = table.ClerkMchID
	}
	if table.IsGroupCentre != 0 {
		update["is_group_centre"] = table.IsGroupCentre
	}
	if table.Sex != 0 {
		update["sex"] = table.Sex
	}
	if table.BirthDate != "" {
		update["birth_date"] = table.BirthDate
	}
	if table.IsVgoods != 0 {
		update["is_vgoods"] = table.IsVgoods
	}
	if table.ShareCatID != 0 {
		update["share_cat_id"] = table.ShareCatID
	}
	if table.UserLabel != 0 {
		update["user_label"] = table.UserLabel
	}
	if table.ConsumeRebateTotalFyjine != "" {
		update["consume_rebate_total_fyjine"] = table.ConsumeRebateTotalFyjine
	}
	if table.ConsumeRebateFyjine != "" {
		update["consume_rebate_fyjine"] = table.ConsumeRebateFyjine
	}
	if table.ConsumeRebateTotalFyjifen != "" {
		update["consume_rebate_total_fyjifen"] = table.ConsumeRebateTotalFyjifen
	}
	if table.ConsumeRebateFyjifen != "" {
		update["consume_rebate_fyjifen"] = table.ConsumeRebateFyjifen
	}
	if table.PickMch != 0 {
		update["pick_mch"] = table.PickMch
	}
	if table.BalancePassword != "" {
		update["balance_password"] = table.BalancePassword
	}
	if table.IsFreePayment != 0 {
		update["is_free_payment"] = table.IsFreePayment
	}
	if table.ProvinceID != 0 {
		update["province_id"] = table.ProvinceID
	}
	if table.Source != 0 {
		update["source"] = table.Source
	}
	if table.Subsidies != "" {
		update["subsidies"] = table.Subsidies
	}
	if table.ConsumeMoney != "" {
		update["consume_money"] = table.ConsumeMoney
	}
	if table.IsMchShipping != 0 {
		update["is_mch_shipping"] = table.IsMchShipping
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *cyymallUserDao) GetByID(ctx context.Context, id uint64) (*model.CyymallUser, error) {
	// no cache
	if d.cache == nil {
		record := &model.CyymallUser{}
		err := d.db.WithContext(ctx).Where("id = ?", id).First(record).Error
		return record, err
	}

	// get from cache or database
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	if errors.Is(err, model.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.CyymallUser{}
			err = d.db.WithContext(ctx).Where("id = ?", id).First(table).Error
			if err != nil {
				// if data is empty, set not found cache to prevent cache penetration, default expiration time 10 minutes
				if errors.Is(err, model.ErrRecordNotFound) {
					err = d.cache.SetCacheWithNotFound(ctx, id)
					if err != nil {
						return nil, err
					}
					return nil, model.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			err = d.cache.Set(ctx, id, table, cache.CyymallUserExpireTime)
			if err != nil {
				return nil, fmt.Errorf("cache.Set error: %v, id=%d", err, id)
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.CyymallUser)
		if !ok {
			return nil, model.ErrRecordNotFound
		}
		return table, nil
	} else if errors.Is(err, cacheBase.ErrPlaceholder) {
		return nil, model.ErrRecordNotFound
	}

	// fail fast, if cache error return, don't request to db
	return nil, err
}

// GetByCondition get a record by condition
// query conditions:
//
//	name: column name
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in
//	value: column value, if exp=in, multiple values are separated by commas
//	logic: logical type, defaults to and when value is null, only &(and), ||(or)
//
// example: find a male aged 20
//
//	condition = &query.Conditions{
//	    Columns: []query.Column{
//		{
//			Name:    "age",
//			Value:   20,
//		},
//		{
//			Name:  "gender",
//			Value: "male",
//		},
//	}
func (d *cyymallUserDao) GetByCondition(ctx context.Context, c *query.Conditions) (*model.CyymallUser, error) {
	queryStr, args, err := c.ConvertToGorm()
	if err != nil {
		return nil, err
	}

	table := &model.CyymallUser{}
	err = d.db.WithContext(ctx).Where(queryStr, args...).First(table).Error
	if err != nil {
		return nil, err
	}

	return table, nil
}

// GetByIDs get records by batch id
func (d *cyymallUserDao) GetByIDs(ctx context.Context, ids []uint64) (map[uint64]*model.CyymallUser, error) {
	// no cache
	if d.cache == nil {
		var records []*model.CyymallUser
		err := d.db.WithContext(ctx).Where("id IN (?)", ids).Find(&records).Error
		if err != nil {
			return nil, err
		}
		itemMap := make(map[uint64]*model.CyymallUser)
		for _, record := range records {
			itemMap[record.ID] = record
		}
		return itemMap, nil
	}

	// get form cache or database
	itemMap, err := d.cache.MultiGet(ctx, ids)
	if err != nil {
		return nil, err
	}

	var missedIDs []uint64
	for _, id := range ids {
		_, ok := itemMap[id]
		if !ok {
			missedIDs = append(missedIDs, id)
			continue
		}
	}

	// get missed data
	if len(missedIDs) > 0 {
		// find the id of an active placeholder, i.e. an id that does not exist in database
		var realMissedIDs []uint64
		for _, id := range missedIDs {
			_, err = d.cache.Get(ctx, id)
			if errors.Is(err, cacheBase.ErrPlaceholder) {
				continue
			}
			realMissedIDs = append(realMissedIDs, id)
		}

		if len(realMissedIDs) > 0 {
			var missedData []*model.CyymallUser
			err = d.db.WithContext(ctx).Where("id IN (?)", realMissedIDs).Find(&missedData).Error
			if err != nil {
				return nil, err
			}

			if len(missedData) > 0 {
				for _, data := range missedData {
					itemMap[data.ID] = data
				}
				err = d.cache.MultiSet(ctx, missedData, cache.CyymallUserExpireTime)
				if err != nil {
					return nil, err
				}
			} else {
				for _, id := range realMissedIDs {
					_ = d.cache.SetCacheWithNotFound(ctx, id)
				}
			}
		}
	}

	return itemMap, nil
}

// GetByLastID get paging records by last id and limit
func (d *cyymallUserDao) GetByLastID(ctx context.Context, lastID uint64, limit int, sort string) ([]*model.CyymallUser, error) {
	page := query.NewPage(0, limit, sort)

	records := []*model.CyymallUser{}
	err := d.db.WithContext(ctx).Order(page.Sort()).Limit(page.Size()).Where("id < ?", lastID).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

// GetByColumns get paging records by column information,
// Note: query performance degrades when table rows are very large because of the use of offset.
//
// params includes paging parameters and query parameters
// paging parameters (required):
//
//	page: page number, starting from 0
//	size: lines per page
//	sort: sort fields, default is id backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
//
// query parameters (not required):
//
//	name: column name
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in
//	value: column value, if exp=in, multiple values are separated by commas
//	logic: logical type, defaults to and when value is null, only &(and), ||(or)
//
// example: search for a male over 20 years of age
//
//	params = &query.Params{
//	    Page: 0,
//	    Size: 20,
//	    Columns: []query.Column{
//		{
//			Name:    "age",
//			Exp: ">",
//			Value:   20,
//		},
//		{
//			Name:  "gender",
//			Value: "male",
//		},
//	}
func (d *cyymallUserDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.CyymallUser, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.CyymallUser{}).Select([]string{"id"}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.CyymallUser{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *cyymallUserDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.CyymallUser) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *cyymallUserDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	update := map[string]interface{}{
		"deleted_at": time.Now(),
	}
	err := tx.WithContext(ctx).Model(&model.CyymallUser{}).Where("id = ?", id).Updates(update).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *cyymallUserDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.CyymallUser) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}
