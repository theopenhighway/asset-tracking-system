package dto

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetDepartmentbyId struct {
	Id int `json:"id" binding:"required"`
}
