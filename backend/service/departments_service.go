package service

import "asset-tracking-system/repository"

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

func (s *deptService) CreateDept() {

}

func (s *deptService) FindDeptById() {

}

func (s *deptService) GetAllDept() {

}

func 
