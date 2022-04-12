package berror

const (
	UnauthenticatedMsg         = "User is not authenticated"
	InvalidFieldMsg            = "Invalid field was provided"
	InternalMsg                = "Unexpected internal error occurred"
	DuplicateMsg               = "Entity already exists"
	InvalidPayloadMsg          = "Unable to deserialize payload"
	NotExistsMsg               = "Entity does not exist"
	RequiredEntityNotExistsMsg = "A required entity has not been created"
	DisabledMsg                = "Entity has been disabled"
	AuthenticationFailedMsg    = "Authentication failed"
)

type Option func(e *Error)

func WithUnauthenticated() Option {
	return func(e *Error) {
		e.code = CodeUnauthenticated
		e.userMessage = UnauthenticatedMsg
	}
}

func WithInternal() Option {
	return func(e *Error) {
		e.code = CodeInternal
		e.userMessage = InternalMsg
	}
}

func WithInvalidFields(invalidFields ...InvalidField) Option {
	return func(e *Error) {
		e.code = CodeInvalidField
		e.userMessage = InvalidFieldMsg
		e.fields = invalidFields
	}
}

func WithDuplicate() Option {
	return func(e *Error) {
		e.code = CodeDuplicate
		e.userMessage = DuplicateMsg
	}
}

func WithNotExists() Option {
	return func(e *Error) {
		e.code = CodeNotExists
		e.userMessage = NotExistsMsg
	}
}

func WithInvalidPayload() Option {
	return func(e *Error) {
		e.code = CodeInvalidPayload
		e.userMessage = InvalidPayloadMsg
	}
}

func WithRequiredEntityNotExist() Option {
	return func(e *Error) {
		e.code = CodeRequiredEntityNotExists
		e.userMessage = RequiredEntityNotExistsMsg
	}
}

func WithMessage(message string) Option {
	return func(e *Error) {
		e.message = message
	}
}

func WithEntityDisabled() Option {
	return func(e *Error) {
		e.code = CodeDisabled
		e.userMessage = DisabledMsg
	}
}

func WithAuthenticationFailed() Option {
	return func(e *Error) {
		e.code = CodeAuthenticationFailed
		e.userMessage = AuthenticationFailedMsg
	}
}
