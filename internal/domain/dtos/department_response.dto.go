package dtos

import "col-moda/internal/domain/entities"

type DepartmentResponseDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func DepartmentsResponseFromEntitie(departments []entities.Department) []DepartmentResponseDto {
	var departmentResponse []DepartmentResponseDto

	for _, department := range departments {
		departmentResponse = append(departmentResponse, DepartmentResponseDto{
			ID:   department.ID,
			Name: department.Name,
		})
	}

	return departmentResponse
}
