package meet

import "fmt"

// Error represents an API error returned by meethayavo.
type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	if e.StatusCode == 0 {
		return e.Message
	}

	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

// Common SDK errors.
var (
	ErrAuthentication = &Error{
		Message: "authentication failed",
	}

	ErrValidation = &Error{
		Message: "validation failed",
	}

	ErrUnauthorized = &Error{
		Message: "unauthorized",
	}

	ErrForbidden = &Error{
		Message: "forbidden",
	}

	ErrNotFound = &Error{
		Message: "resource not found",
	}

	ErrRateLimited = &Error{
		Message: "rate limit exceeded",
	}

	ErrInternalServer = &Error{
		Message: "internal server error",
	}

	ErrNetwork = &Error{
		Message: "network error",
	}

	ErrTimeout = &Error{
		Message: "request timeout",
	}
)
