package constant

// HTTP Status Codes
const (
	SUCCESS            = 200
	ACCEPTED           = 202
	BAD_REQUEST        = 400
	UNAUTHORIZED       = 401
	METHOD_NOT_ALLOWED = 405
	INTERNAL_ERROR     = 500
)

// Success Messages
const (
	SUCCESS_MESSAGE      = "Success"
	CREATED_MESSAGE      = "Resource created successfully"
	UPDATED_MESSAGE      = "Resource updated successfully"
	DELETED_MESSAGE      = "Resource deleted successfully"
	DATA_FORWARD_SUCCESS = "Data forwarded successfully"
)

// Error Messages
const (
	BAD_REQUEST_MESSAGE        = "Invalid request"
	UNAUTHORIZED_MESSAGE       = "unauthorized"
	METHOD_NOT_ALLOWED_MESSAGE = "Method not allowed"
	INTERNAL_ERROR_MESSAGE     = "Internal server error"
	VALIDATION_ERROR_MESSAGE   = "Validation failed"
	DATABASE_ERROR_MESSAGE     = "Database operation failed"
	ACCOUNT_FETCH_FAILED       = "account fetch failed"
	DESTINATION_FETCH_FAILED   = "destination fetch failed"
	BODY_FAILED                = "error reading body"
	INVAILD_JSON               = "invalid JSON data"
	EMAIL_EXIST                = "this email already have a account"
	ACCOUNT_ID_REQUIRED        = "account_id is required for update"
	INVAILD_STATUS_VALUE       = "in valid status value"
	INVALID_ACCOUNT_ID         = "in valid account_id"
)
