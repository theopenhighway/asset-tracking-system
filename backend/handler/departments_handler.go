package handler

import (
	"asset-tracking-system/apperror"
	"asset-tracking-system/dto"
	"asset-tracking-system/service"
	"net/http"

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
	response, err := h.deptService.GetAllDept()

	if err != nil {
		c.Error(apperror.ErrorUnauthAccess)
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Msg: "dept", Data: response})
}

func (h *DeptHandler) GetDeptById(c *gin.Context) {
	deptId := c.Value("id")

	if deptId == nil {
		c.Error(apperror.ErrorUnauthAccess)
		return
	}

	deptIdx := deptId.(int)

	response, err := h.deptService.FindDeptById(deptIdx)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Msg: "dept", Data: response})
}

func (h *DeptHandler) DeleteDept(c *gin.Context) {

}
