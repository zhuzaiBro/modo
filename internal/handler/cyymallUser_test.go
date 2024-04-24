package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/gohttp"
	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"cyy/internal/cache"
	"cyy/internal/dao"
	"cyy/internal/model"
	"cyy/internal/types"
)

func newCyymallUserHandler() *gotest.Handler {
	// todo additional test field information
	testData := &model.CyymallUser{}
	testData.ID = 1
	testData.CreatedAt = time.Now()
	testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewCyymallUserCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewCyymallUserDao(d.DB, c.ICache.(cache.CyymallUserCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &cyymallUserHandler{iDao: d.IDao.(dao.CyymallUserDao)}
	iHandler := h.IHandler.(CyymallUserHandler)

	testFns := []gotest.RouterInfo{
		{
			FuncName:    "Create",
			Method:      http.MethodPost,
			Path:        "/cyymallUser",
			HandlerFunc: iHandler.Create,
		},
		{
			FuncName:    "DeleteByID",
			Method:      http.MethodDelete,
			Path:        "/cyymallUser/:id",
			HandlerFunc: iHandler.DeleteByID,
		},
		{
			FuncName:    "DeleteByIDs",
			Method:      http.MethodPost,
			Path:        "/cyymallUser/delete/ids",
			HandlerFunc: iHandler.DeleteByIDs,
		},
		{
			FuncName:    "UpdateByID",
			Method:      http.MethodPut,
			Path:        "/cyymallUser/:id",
			HandlerFunc: iHandler.UpdateByID,
		},
		{
			FuncName:    "GetByID",
			Method:      http.MethodGet,
			Path:        "/cyymallUser/:id",
			HandlerFunc: iHandler.GetByID,
		},
		{
			FuncName:    "GetByCondition",
			Method:      http.MethodPost,
			Path:        "/cyymallUser/condition",
			HandlerFunc: iHandler.GetByCondition,
		},
		{
			FuncName:    "ListByIDs",
			Method:      http.MethodPost,
			Path:        "/cyymallUser/list/ids",
			HandlerFunc: iHandler.ListByIDs,
		},
		{
			FuncName:    "ListByLastID",
			Method:      http.MethodGet,
			Path:        "/cyymallUser/list",
			HandlerFunc: iHandler.ListByLastID,
		},
		{
			FuncName:    "List",
			Method:      http.MethodPost,
			Path:        "/cyymallUser/list",
			HandlerFunc: iHandler.List,
		},
	}

	h.GoRunHTTPServer(testFns)

	time.Sleep(time.Millisecond * 200)
	return h
}

func Test_cyymallUserHandler_Create(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := &types.CreateCyymallUserRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.CyymallUser))

	h.MockDao.SQLMock.ExpectBegin()
	args := h.MockDao.GetAnyArgs(h.TestData)
	h.MockDao.SQLMock.ExpectExec("INSERT INTO .*").
		WithArgs(args[:len(args)-1]...). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(1, 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("Create"), testData)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

}

func Test_cyymallUserHandler_DeleteByID(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Delete(result, h.GetRequestURL("DeleteByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Delete(result, h.GetRequestURL("DeleteByID", 0))
	assert.NoError(t, err)

	// delete error test
	err = gohttp.Delete(result, h.GetRequestURL("DeleteByID", 111))
	assert.Error(t, err)
}

func Test_cyymallUserHandler_DeleteByIDs(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteCyymallUsersByIDsRequest{IDs: []uint64{testData.ID}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), nil)
	assert.NoError(t, err)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("DeleteByIDs"), &types.DeleteCyymallUsersByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_cyymallUserHandler_UpdateByID(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := &types.UpdateCyymallUserByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.CyymallUser))

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &gohttp.StdResult{}
	err := gohttp.Put(result, h.GetRequestURL("UpdateByID", testData.ID), testData)
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Put(result, h.GetRequestURL("UpdateByID", 0), testData)
	assert.NoError(t, err)

	// update error test
	err = gohttp.Put(result, h.GetRequestURL("UpdateByID", 111), testData)
	assert.Error(t, err)
}

func Test_cyymallUserHandler_GetByID(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").
		WithArgs(testData.ID).
		WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Get(result, h.GetRequestURL("GetByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = gohttp.Get(result, h.GetRequestURL("GetByID", 0))
	assert.NoError(t, err)

	// get error test
	err = gohttp.Get(result, h.GetRequestURL("GetByID", 111))
	assert.Error(t, err)
}

func Test_cyymallUserHandler_GetByCondition(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("GetByCondition"), &types.GetCyymallUserByConditionRequest{
		query.Conditions{
			Columns: []query.Column{
				{
					Name:  "id",
					Value: testData.ID,
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero error test
	err = gohttp.Post(result, h.GetRequestURL("GetByCondition"), nil)
	assert.NoError(t, err)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("GetByCondition"), &types.GetCyymallUserByConditionRequest{
		query.Conditions{
			Columns: []query.Column{
				{
					Name:  "id",
					Value: 2,
				},
			},
		},
	})
	assert.Error(t, err)
}

func Test_cyymallUserHandler_ListByIDs(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.ListCyymallUsersByIDsRequest{IDs: []uint64{testData.ID}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	_ = gohttp.Post(result, h.GetRequestURL("ListByIDs"), nil)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("ListByIDs"), &types.ListCyymallUsersByIDsRequest{IDs: []uint64{111}})
	assert.Error(t, err)
}

func Test_cyymallUserHandler_ListByLastID(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Get(result, h.GetRequestURL("ListByLastID"), gohttp.KV{"lastID": 0, "size": 10})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// error test
	err = gohttp.Get(result, h.GetRequestURL("ListByLastID"), gohttp.KV{"lastID": 0, "size": 10, "sort": "unknown-column"})
	assert.Error(t, err)
}

func Test_cyymallUserHandler_List(t *testing.T) {
	h := newCyymallUserHandler()
	defer h.Close()
	testData := h.TestData.(*model.CyymallUser)

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
		AddRow(testData.ID, testData.CreatedAt, testData.UpdatedAt)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &gohttp.StdResult{}
	err := gohttp.Post(result, h.GetRequestURL("List"), &types.ListCyymallUsersRequest{query.Params{
		Page: 0,
		Size: 10,
		Sort: "ignore count", // ignore test count
	}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// nil params error test
	err = gohttp.Post(result, h.GetRequestURL("List"), nil)
	assert.NoError(t, err)

	// get error test
	err = gohttp.Post(result, h.GetRequestURL("List"), &types.ListCyymallUsersRequest{query.Params{
		Page: 0,
		Size: 10,
		Sort: "unknown-column",
	}})
	assert.Error(t, err)
}

func TestNewCyymallUserHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewCyymallUserHandler()
}
