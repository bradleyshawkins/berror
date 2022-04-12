package berror

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	err         error
	code        Code
	message     string
	userMessage string
	fields      []InvalidField
}

type Reason string

const (
	ReasonMissing = "MISSING"
	ReasonInvalid = "INVALID"
)

type InvalidField struct {
	Field  string
	Reason Reason
}

func New(msg string, options ...Option) *Error {
	return Wrap(errors.New(msg), options...)
}

func Wrap(err error, options ...Option) *Error {
	e := &Error{
		err:         err,
		code:        CodeUnknown,
		userMessage: "Unknown error occurred",
	}

	for _, option := range options {
		option(e)
	}

	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %v, Code: %v, Message: %v", e.err.Error(), e.code, e.message)
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) UserMessage() string {
	return e.userMessage
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) InvalidFields() []InvalidField {
	return e.fields
}

func (e *Error) HttpStatusCode() int {
	code, ok := codeHttpStatusCodeMap[e.code]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}

func (e *Error) WriteAsJson(w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		http.Error(w, "unable to write json response", http.StatusInternalServerError)
	}
}

func NewInternal(msg string) *Error {
	return New(msg, WithInternal())
}

func WrapInternal(err error, msg string) *Error {
	return Wrap(err, WithMessage(msg), WithInternal())
}
