package dtos

type DepartmentRequestDto struct {
	DepartmentID   *int   `form:"department_id"` // Cambia "json" por "form"
	DepartmentName string `form:"department_name"`
}
