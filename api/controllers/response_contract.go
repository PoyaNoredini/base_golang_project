package controllers

type ResponseContract interface {
    Success(ctx interface{}, data interface{})
    Error(ctx interface{}, code int, message string)
}