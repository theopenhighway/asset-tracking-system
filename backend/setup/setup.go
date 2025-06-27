package setup

import (
	"asset-tracking-system/handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type HandlerOps struct {
	deptHandler *handler.DeptHandler
}

func SetupHandlers(db *sql.DB) *HandlerOps {
	return nil
}

func SetupRoutes(h *HandlerOps) *gin.Engine {
	g := gin.New()

	return g
}
