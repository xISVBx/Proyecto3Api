package dtos

type CityRequestDto struct {
	CityID       *int   `form:"city_id"`       // Cambia "json" por "form"
	DepartmentID *int   `form:"department_id"` // Cambia "json" por "form"
	CityName     string `form:"city_name"`
}
