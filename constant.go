package gouse

import "errors"

const (
	INFO_LOG_PREFIX   string = "[INFO]"
	ERROR_LOG_PREFIX  string = "[ERROR]"
	ACCESS_LOG_PREFIX string = "[ACCESS]"
)

const (
	CHAIN_STR = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[]|:<>?/.,';][=-`~"
	CHAIN_NUM = "0123456789"
)

const (
	ERROR_LOG_PATH  = "logs/error.log"
	INFO_LOG_PATH   = "logs/info.log"
	ACCESS_LOG_PATH = "logs/access.log"
)

const (
	DESC_TEST    = "Expected %v but it got %v"
	DESC_CONSOLE = "Unsupported action type"
)

var (
	ERROR_REQUIRED_UUID error = errors.New("uuid is required")
	ERROR_INVALID_UUID  error = errors.New("uuid is invalid")

	ERROR_INVALID_EMAIL error = errors.New("email is invalid")
	ERROR_EMAIL_LENGTH  error = errors.New("email must be between 8 and 32 characters long")

	ERROR_INVALID_USERNAME error = errors.New("username is invalid, must be between 3 and 20 characters long and only contains letters, numbers and underscore")

	ERROR_INVALID_PASSWORD       error = errors.New("password is invalid, must be between 8 and 32 characters long and contains at least one lowercase letter, one uppercase letter, one digit and one special character")
	ERROR_PASSWORD_LENGTH        error = errors.New("password must be between 8 and 32 characters long")
	ERROR_PASSWORD_EMPTY_LOWER   error = errors.New("password must contain at least one lowercase letter")
	ERROR_PASSWORD_EMPTY_UPPER   error = errors.New("password must contain at least one uppercase letter")
	ERROR_PASSWORD_EMPTY_DIGIT   error = errors.New("password must contain at least one digit")
	ERROR_PASSWORD_EMPTY_SPECIAL error = errors.New("password must contain at least one special character")

	ERROR_HASH_PASSWORD    error = errors.New("error when encrypt password")
	ERROR_COMPARE_PASSWORD error = errors.New("password not match")

	ERROR_INVALID_PHONE error = errors.New("phone is invalid")
)
