package dto

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetDepartmentbyIdRequest struct {
	Id int `json:"id" binding:"required"`
}

type DepartmentResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateDepartmentRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
