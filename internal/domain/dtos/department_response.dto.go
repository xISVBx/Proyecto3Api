package dtos

import "col-moda/internal/domain/entities"

type DepartmentResponse struct {
	ID   int   `json:"id"`
	Name string `json:"name"`
}

func DepartmentsResponseFromEntitie(departments []entities.Department) []DepartmentResponse {
	var departmentResponse []DepartmentResponse

	for _, department := range departments {
		departmentResponse = append(departmentResponse, DepartmentResponse{
			ID:   department.ID,
			Name: department.Name,
		})
	}

	return departmentResponse
}