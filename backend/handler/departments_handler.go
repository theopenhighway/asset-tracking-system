package handler

import (
	"asset-tracking-system/service"

	"github.com/gin-gonic/gin"
)

type DeptHandler struct {
	deptService service.DeptService
}

func NewDeptHandler(ds service.DeptService) *DeptHandler {
	return &DeptHandler{
		deptService: ds,
	}

}

func (h *DeptHandler) AddDept(c *gin.Context) {

}

func (h *DeptHandler) UpdateDept(c *gin.Context) {

}

func (h *DeptHandler) SearchDept(c *gin.Context) {

}

func (h *DeptHandler) GetAllDept(c *gin.Context) {

}

func (h *DeptHandler) DeleteDept(c *gin.Context) {

}
