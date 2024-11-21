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
	ErrorRequiredUuid error = errors.New("uuid is required")
	ErrorInvalidUuid  error = errors.New("uuid is invalid")

	ErrorInvalidEmail error = errors.New("email is invalid")
	ErrorEmailLength  error = errors.New("email must be between 8 and 32 characters long")

	ErrorInvalidUsername error = errors.New("username is invalid, must be between 3 and 20 characters long and only contains letters, numbers and underscore")

	ErrorInvalidPassword      error = errors.New("password is invalid, must be between 8 and 32 characters long and contains at least one lowercase letter, one uppercase letter, one digit and one special character")
	ErrorPasswordLength       error = errors.New("password must be between 8 and 32 characters long")
	ErrorPasswordEmptyLower   error = errors.New("password must contain at least one lowercase letter")
	ErrorPasswordEmptyUpper   error = errors.New("password must contain at least one uppercase letter")
	ErrorPasswordEmptyDigit   error = errors.New("password must contain at least one digit")
	ErrorPasswordEmptySpecial error = errors.New("password must contain at least one special character")

	ErrorHashPassword    error = errors.New("error when encrypt password")
	ErrorComparePassword error = errors.New("password not match")

	ErrorInvalidPhone error = errors.New("phone is invalid")
)
