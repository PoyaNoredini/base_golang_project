package errors

import "net/http"

type AppError struct {
    Status  int
    Message string
}

func (e *AppError) Error() string {
    return e.Message
}

func BadRequest(msg string) *AppError  {
     return &AppError{
        Status: http.StatusBadRequest,
        Message: msg,
     }
}
func NotFound(msg string) *AppError    {
     return &AppError{
        Status: http.StatusNotFound,
        Message: msg,
     }
}
func Unauthorized(msg string) *AppError {
     return &AppError{
        Status: http.StatusUnauthorized,
        Message: msg,
     }
}
func Internal(msg string) *AppError    {
     return &AppError{
        Status: http.StatusInternalServerError,
        Message: msg,
     }

func Forbidden(msg string) *AppError {
   return &AppError{
      Status: http.StatusForbidden,
      Message: msg,
   }
}
}