package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"cyy/internal/model"
)

func newCyymallUserCache() *gotest.Cache {
	record1 := &model.CyymallUser{}
	record1.ID = 1
	record2 := &model.CyymallUser{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCyymallUserCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_cyymallUserCache_Set(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CyymallUser)
	err := c.ICache.(CyymallUserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CyymallUserCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_cyymallUserCache_Get(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CyymallUser)
	err := c.ICache.(CyymallUserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CyymallUserCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CyymallUserCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_cyymallUserCache_MultiGet(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	var testData []*model.CyymallUser
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CyymallUser))
	}

	err := c.ICache.(CyymallUserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CyymallUserCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.CyymallUser))
	}
}

func Test_cyymallUserCache_MultiSet(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	var testData []*model.CyymallUser
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CyymallUser))
	}

	err := c.ICache.(CyymallUserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_cyymallUserCache_Del(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CyymallUser)
	err := c.ICache.(CyymallUserCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_cyymallUserCache_SetCacheWithNotFound(t *testing.T) {
	c := newCyymallUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CyymallUser)
	err := c.ICache.(CyymallUserCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCyymallUserCache(t *testing.T) {
	c := NewCyymallUserCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewCyymallUserCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewCyymallUserCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
