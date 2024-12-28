package utils

func CreateResponse(
	code string,
	message string,
	data interface{},
	errorCode *string,
	errorMessage *string,
	errorDetails *string,
	pagination *Pagination,
) Response {
	var errorObj *Error
	if errorCode != nil || errorMessage != nil || errorDetails != nil {
		errorObj = &Error{
			Code:    errorCode,
			Message: errorMessage,
			Details: errorDetails,
		}
	}

	return Response{
		Code:       code,
		Message:    message,
		Data:       data,
		Error:      errorObj,
		Pagination: pagination,
	}
}
