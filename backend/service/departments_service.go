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

func (s *deptService) CreateDept(dept *dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error) {
	newDept, err := s.deptRepository.CreateDepartment(dept.Name)

	if err != nil {
		return nil, err
	}

	result := dto.DepartmentResponse{
		Id:   newDept.Id,
		Name: newDept.Name,
	}
	return &result, nil
}

func (s *deptService) FindDeptById(deptId *dto.GetDepartmentbyIdRequest) (*dto.DepartmentResponse, error) {
	dept, err := s.deptRepository.GetDepartmentbyId(deptId.Id)

	if err != nil {
		return nil, err
	}
	if dept == nil {
		return nil, nil
	}

	result := dto.DepartmentResponse{
		Id:   dept.Id,
		Name: dept.Name,
	}
	return &result, nil
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
