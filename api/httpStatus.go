package api

type Error struct {
	Message string `json:"message"`
}

type APIError struct {
	Status Status `json:"status"`
	Error  Error  `json:"error"`
}

func CreateError(code int, message string) APIError {
	return APIError{
		Status: Status{
			Code: code,
			Ok:   false,
		},
		Error: Error{
			Message: message,
		},
	}
}

func CreateSuccess(code int) Status {
	return Status{
		Code: code,
		Ok:   true,
	}
}

func OK() Status {
	return CreateSuccess(200)
}

func Created() Status {
	return CreateSuccess(201)
}

func Accepted() Status {
	return CreateSuccess(202)
}

func NonAuthoritativeInformation() Status {
	return CreateSuccess(203)
}

func NoContent() Status {
	return CreateSuccess(204)
}

func ResetContent() Status {
	return CreateSuccess(205)
}

func PartialContent() Status {
	return CreateSuccess(206)
}

func MultiStatus() Status {
	return CreateSuccess(207)
}

func AlreadyReported() Status {
	return CreateSuccess(208)
}

func IMUsed() Status {
	return CreateSuccess(226)
}

func BadRequest(message *string) APIError {
	code := 400

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Bad request")
}

func Unauthorized(message *string) APIError {
	code := 401

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Unauthorized")
}

// Unlike 401 Unauthorized, the client's identity is known to the server.
func Forbidden(message *string) APIError {
	code := 403

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Forbidden")
}

func NotFound(message *string) APIError {
	code := 404

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(404, "Not found")
}

func MethodNotAllowed(message *string) APIError {
	code := 405

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Method Not Allowed")
}

func Conflict(message *string) APIError {
	code := 409

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Conflict")
}

func UnsupportedMediaType(message *string) APIError {
	code := 415

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Unsupported Media Type")
}

func IAmTeapot(message *string) APIError {
	code := 418

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "I'm a teapot")
}

func Locked(message *string) APIError {
	code := 423

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Locked")
}

func TooManyRequests(message *string) APIError {
	code := 429

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Too Many Requests")
}

func Internal(message *string) APIError {
	code := 500

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Internal")
}

func BadGateway(message *string) APIError {
	code := 502

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Bad Gateway")
}

func ServiceUnavailable(message *string) APIError {
	code := 503

	if message != nil {
		return CreateError(code, *message)
	}

	return CreateError(code, "Service Unavailable")
}
