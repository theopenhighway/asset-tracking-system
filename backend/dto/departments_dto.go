package dto

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
}

type GetDepartmentbyIdRequest struct {
	Id int `json:"id" binding:"required"`
}

type DepartmentResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
