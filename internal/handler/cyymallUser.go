package handler

import (
	"errors"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/gin/response"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	"cyy/internal/cache"
	"cyy/internal/dao"
	"cyy/internal/ecode"
	"cyy/internal/model"
	"cyy/internal/types"
)

var _ CyymallUserHandler = (*cyymallUserHandler)(nil)

// CyymallUserHandler defining the handler interface
type CyymallUserHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	DeleteByIDs(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	GetByCondition(c *gin.Context)
	ListByIDs(c *gin.Context)
	ListByLastID(c *gin.Context)
	List(c *gin.Context)
}

type cyymallUserHandler struct {
	iDao dao.CyymallUserDao
}

// NewCyymallUserHandler creating the handler interface
func NewCyymallUserHandler() CyymallUserHandler {
	return &cyymallUserHandler{
		iDao: dao.NewCyymallUserDao(
			model.GetDB(),
			cache.NewCyymallUserCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create cyymallUser
// @Description submit information to create cyymallUser
// @Tags cyymallUser
// @accept json
// @Produce json
// @Param data body types.CreateCyymallUserRequest true "cyymallUser information"
// @Success 200 {object} types.CreateCyymallUserRespond{}
// @Router /api/v1/cyymallUser [post]
func (h *cyymallUserHandler) Create(c *gin.Context) {
	form := &types.CreateCyymallUserRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	cyymallUser := &model.CyymallUser{}
	err = copier.Copy(cyymallUser, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateCyymallUser)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, cyymallUser)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": cyymallUser.ID})
}

// DeleteByID delete a record by id
// @Summary delete cyymallUser
// @Description delete cyymallUser by id
// @Tags cyymallUser
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteCyymallUserByIDRespond{}
// @Router /api/v1/cyymallUser/{id} [delete]
func (h *cyymallUserHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getCyymallUserIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// DeleteByIDs delete records by batch id
// @Summary delete cyymallUsers
// @Description delete cyymallUsers by batch id
// @Tags cyymallUser
// @Param data body types.DeleteCyymallUsersByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteCyymallUsersByIDsRespond{}
// @Router /api/v1/cyymallUser/delete/ids [post]
func (h *cyymallUserHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteCyymallUsersByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.DeleteByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// UpdateByID update information by id
// @Summary update cyymallUser
// @Description update cyymallUser information by id
// @Tags cyymallUser
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateCyymallUserByIDRequest true "cyymallUser information"
// @Success 200 {object} types.UpdateCyymallUserByIDRespond{}
// @Router /api/v1/cyymallUser/{id} [put]
func (h *cyymallUserHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getCyymallUserIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateCyymallUserByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	cyymallUser := &model.CyymallUser{}
	err = copier.Copy(cyymallUser, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDCyymallUser)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, cyymallUser)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get cyymallUser detail
// @Description get cyymallUser detail by id
// @Tags cyymallUser
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetCyymallUserByIDRespond{}
// @Router /api/v1/cyymallUser/{id} [get]
func (h *cyymallUserHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getCyymallUserIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	cyymallUser, err := h.iDao.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID not found", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.CyymallUserObjDetail{}
	err = copier.Copy(data, cyymallUser)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDCyymallUser)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"cyymallUser": data})
}

// GetByCondition get a record by condition
// @Summary get cyymallUser by condition
// @Description get cyymallUser by condition
// @Tags cyymallUser
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetCyymallUserByConditionRespond{}
// @Router /api/v1/cyymallUser/condition [post]
func (h *cyymallUserHandler) GetByCondition(c *gin.Context) {
	form := &types.GetCyymallUserByConditionRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	err = form.Conditions.CheckValid()
	if err != nil {
		logger.Warn("Parameters error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	cyymallUser, err := h.iDao.GetByCondition(ctx, &form.Conditions)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByCondition not found", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByCondition error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.CyymallUserObjDetail{}
	err = copier.Copy(data, cyymallUser)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDCyymallUser)
		return
	}
	data.ID = utils.Uint64ToStr(cyymallUser.ID)

	response.Success(c, gin.H{"cyymallUser": data})
}

// ListByIDs list of records by batch id
// @Summary list of cyymallUsers by batch id
// @Description list of cyymallUsers by batch id
// @Tags cyymallUser
// @Param data body types.ListCyymallUsersByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListCyymallUsersByIDsRespond{}
// @Router /api/v1/cyymallUser/list/ids [post]
func (h *cyymallUserHandler) ListByIDs(c *gin.Context) {
	form := &types.ListCyymallUsersByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	cyymallUserMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	cyymallUsers := []*types.CyymallUserObjDetail{}
	for _, id := range form.IDs {
		if v, ok := cyymallUserMap[id]; ok {
			record, err := convertCyymallUser(v)
			if err != nil {
				response.Error(c, ecode.ErrListCyymallUser)
				return
			}
			cyymallUsers = append(cyymallUsers, record)
		}
	}

	response.Success(c, gin.H{
		"cyymallUsers": cyymallUsers,
	})
}

// ListByLastID get records by last id and limit
// @Summary list of cyymallUsers by last id and limit
// @Description list of cyymallUsers by last id and limit
// @Tags cyymallUser
// @accept json
// @Produce json
// @Param lastID query int true "last id, default is MaxInt32" default(0)
// @Param limit query int false "size in each page" default(10)
// @Param sort query string false "sort by column name of table, and the "-" sign before column name indicates reverse order" default(-id)
// @Success 200 {object} types.ListCyymallUsersRespond{}
// @Router /api/v1/cyymallUser/list [get]
func (h *cyymallUserHandler) ListByLastID(c *gin.Context) {
	lastID := utils.StrToUint64(c.Query("lastID"))
	if lastID == 0 {
		lastID = math.MaxInt32
	}
	limit := utils.StrToInt(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	sort := c.Query("sort")

	ctx := middleware.WrapCtx(c)
	cyymallUsers, err := h.iDao.GetByLastID(ctx, lastID, limit, sort)
	if err != nil {
		logger.Error("GetByLastID error", logger.Err(err), logger.Uint64("latsID", lastID), logger.Int("limit", limit), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertCyymallUsers(cyymallUsers)
	if err != nil {
		response.Error(c, ecode.ErrListByLastIDCyymallUser)
		return
	}

	response.Success(c, gin.H{
		"cyymallUsers": data,
	})
}

// List of records by query parameters
// @Summary list of cyymallUsers by query parameters
// @Description list of cyymallUsers by paging and conditions
// @Tags cyymallUser
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListCyymallUsersRespond{}
// @Router /api/v1/cyymallUser/list [post]
func (h *cyymallUserHandler) List(c *gin.Context) {
	form := &types.ListCyymallUsersRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	cyymallUsers, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertCyymallUsers(cyymallUsers)
	if err != nil {
		response.Error(c, ecode.ErrListCyymallUser)
		return
	}

	response.Success(c, gin.H{
		"cyymallUsers": data,
		"total":        total,
	})
}

func getCyymallUserIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertCyymallUser(cyymallUser *model.CyymallUser) (*types.CyymallUserObjDetail, error) {
	data := &types.CyymallUserObjDetail{}
	err := copier.Copy(data, cyymallUser)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(cyymallUser.ID)
	return data, nil
}

func convertCyymallUsers(fromValues []*model.CyymallUser) ([]*types.CyymallUserObjDetail, error) {
	toValues := []*types.CyymallUserObjDetail{}
	for _, v := range fromValues {
		data, err := convertCyymallUser(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
