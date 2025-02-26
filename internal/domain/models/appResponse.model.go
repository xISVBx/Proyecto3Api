package models

type AppResponse struct {
	Success bool     `json:"success"`
	Data    any      `json:"data"`
	Errors  []string `json:"errors"`
}

func ResSuccess(data any) AppResponse {
	return AppResponse{
		Success: true,
		Data:    data,
		Errors:  make([]string, 0),
	}
}

func ResFailed(errors []string) AppResponse {
	return AppResponse{
		Success: false,
		Data:    nil,
		Errors:  errors,
	}
}

func ResError(error string) AppResponse {
	listErrors := make([]string, 0)

	listErrors = append(listErrors, error)

	return AppResponse{
		Success: false,
		Data:    nil,
		Errors:  listErrors,
	}
}

func ResFromApp(errors *AppError) AppResponse {
	listErrors := make([]string, 0)

	listErrors = append(listErrors, errors.Err.Error())

	return AppResponse{
		Success: false,
		Data:    nil,
		Errors:  listErrors,
	}
}
