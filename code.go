package berror

import "net/http"

type Code int

const (
	CodeUnknown Code = iota + 1
	CodeUnauthenticated
	CodeInvalidPayload
	CodeInvalidField
	CodeInternal
	CodeDuplicate
	CodeNotExists
	CodeRequiredEntityNotExists
	CodeDisabled
	CodeAuthenticationFailed
)

var codeHttpStatusCodeMap = map[Code]int{
	CodeUnknown:                 http.StatusInternalServerError,
	CodeUnauthenticated:         http.StatusUnauthorized,
	CodeInvalidField:            http.StatusBadRequest,
	CodeInternal:                http.StatusInternalServerError,
	CodeDuplicate:               http.StatusConflict,
	CodeInvalidPayload:          http.StatusBadRequest,
	CodeNotExists:               http.StatusNotFound,
	CodeRequiredEntityNotExists: http.StatusConflict,
	CodeDisabled:                http.StatusNotFound,
	CodeAuthenticationFailed:    http.StatusUnauthorized,
}
