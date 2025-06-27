package service

import (
	"asset-tracking-system/dto"
	"asset-tracking-system/repository"
)

type DeptService interface {
	CreateDept(dept *dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error)
	FindDeptById(deptId *dto.GetDepartmentbyIdRequest) (*dto.DepartmentResponse, error)
	GetAllDept() (*[]dto.DepartmentResponse, error)
	DeleteDeptById(deptId *dto.GetDepartmentbyIdRequest) error
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
	dept, err := s.deptRepository.GetDepartment()
	deptArr := []dto.DepartmentResponse{}

	if err != nil {
		return nil, err
	}

	return &deptArr, nil
}

func (s *deptService) UpdateDept(deptId *dto.UpdateDepartmentRequest) (*dto.DepartmentResponse, error) {
	dept, err := s.deptRepository.UpdateDepartment(deptId.Id, deptId.Name)

	if err != nil {
		return nil, err
	}
	result := dto.DepartmentResponse{
		Id:   dept.Id,
		Name: dept.Name,
	}
	return &result, nil
}

func (s *deptService) DeleteDeptById(deptId *dto.GetDepartmentbyIdRequest) error {
	err := s.deptRepository.DeleteDepartment(deptId.Id)

	if err != nil {
		return err
	}

	return nil
}
