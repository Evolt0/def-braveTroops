package errors

import "net/http"

//go:generate  gen-const-msg -t=StatusError -m
type StatusError int

const (
	// 请求参数错误
	BadRequest StatusError = http.StatusBadRequest
)
const (
	// 资源未找到
	NotFound StatusError = http.StatusNotFound
)
const (
	// 内部处理错误
	InternalServerError StatusError = http.StatusInternalServerError
)

const (
	// 状态冲突
	Conflict StatusError = http.StatusConflict
)
const (
	// 禁止访问
	Forbidden StatusError = http.StatusForbidden
)
const (
	// 未授权
	Unauthorized StatusError = http.StatusUnauthorized
)
