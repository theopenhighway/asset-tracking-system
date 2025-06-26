package service

import (
	"asset-tracking-system/dto"
	"asset-tracking-system/repository"
)

type DeptService interface {
}

type deptService struct {
	deptRepository repository.DeptRepostiory
}

func NewDeptService(dr repository.DeptRepostiory) *deptService {
	return &deptService{
		deptRepository: dr,
	}
}

func (s *deptService) CreateDept(*dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error) {
	return nil, nil
}

func (s *deptService) FindDeptById(*dto.GetDepartmentbyIdRequest) (*dto.DepartmentResponse, error) {
	return nil, nil
}

func (s *deptService) GetAllDept() (*[]dto.DepartmentResponse, error) {
	return nil, nil
}

func (s *deptService) UpdateDept() (*dto.DepartmentResponse, error) {
	return nil, nil
}

func (s *deptService) DeleteDeptById(*dto.GetDepartmentbyIdRequest) error {
	return nil
}
