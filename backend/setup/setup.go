package setup

import (
	"asset-tracking-system/handler"
	"asset-tracking-system/repository"
	"asset-tracking-system/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type HandlerOps struct {
	deptHandler *handler.DeptHandler
}

func SetupHandlers(db *sql.DB) *HandlerOps {
	deptRepository := repository.NewDeptRepository(db)
	deptService := service.NewDeptService(deptRepository)
	deptHandler := handler.NewDeptHandler(deptService)

	return &HandlerOps{
		deptHandler: deptHandler,
	}
}

func SetupRoutes(h *HandlerOps) *gin.Engine {
	g := gin.New()

	g.GET("/all-dept", h.deptHandler.GetAllDept)
	g.GET("/dept-id/{id}", h.deptHandler.GetDeptById)
	g.POST("/add-dept", h.deptHandler.AddDept)
	g.PATCH("/update-dept", h.deptHandler.UpdateDept)

	return g
}
