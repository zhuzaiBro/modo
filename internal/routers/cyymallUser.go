package routers

import (
	"github.com/gin-gonic/gin"

	"cyy/internal/handler"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		cyymallUserRouter(group, handler.NewCyymallUserHandler())
	})
}

func cyymallUserRouter(group *gin.RouterGroup, h handler.CyymallUserHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	// or group.Use(middleware.Auth(middleware.WithVerify(verify))) // token authentication

	group.POST("/cyymallUser", h.Create)
	group.DELETE("/cyymallUser/:id", h.DeleteByID)
	group.POST("/cyymallUser/delete/ids", h.DeleteByIDs)
	group.PUT("/cyymallUser/:id", h.UpdateByID)
	group.GET("/cyymallUser/:id", h.GetByID)
	group.POST("/cyymallUser/condition", h.GetByCondition)
	group.POST("/cyymallUser/list/ids", h.ListByIDs)
	group.GET("/cyymallUser/list", h.ListByLastID)
	group.POST("/cyymallUser/list", h.List)
}
