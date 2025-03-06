package dtos

import (
	"github.com/google/uuid"
)

type CompanyRequestDto struct {
	ID          uuid.UUID `form:"id,omitempty"`
	CompanyName string    `form:"company_name,omitempty"`
}
