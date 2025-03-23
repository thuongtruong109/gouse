package gouse

type ITest struct {
	Name    string
	Input   string
	Want    bool
	WantErr bool
}

type Number interface {
	int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

const (
	// These maybe not supported by all terminals
	DEFAULT_CONSOLE string = "\033[0m"
	RED_CONSOLE     string = "\033[31m"
	GREEN_CONSOLE   string = "\033[32m"
	YELLOW_CONSOLE  string = "\033[33m"
	PURPLE_CONSOLE  string = "\033[34m"
	PINK_CONSOLE    string = "\033[35m"
	CYAN_CONSOLE    string = "\033[36m"
	WHITE_CONSOLE   string = "\033[97m"
	ORANGE_CONSOLE  string = "\033[38;5;208m"
	BLUE_CONSOLE    string = "\033[38;5;27m"
	MAGENTA_CONSOLE string = "\033[38;5;13m"
	GRAY_CONSOLE    string = "\033[37m"
)

const (
	// 	EmptyReg   = `^$`
	// 	SpaceReg   = `\s+`
	// 	WordNumReg = `^[a-zA-Z0-9]+$`
	UsernameReg = `^[a-zA-Z0-9_]{3,20}$`
	EmailLenReg = `^.{8,32}$`
	PhoneReg = `^\+\d{1,2}\s?\(\d{1,4}\)\s?\d{1,6}-\d{1,6}$`

	PasswordLenReg     = `^.{8,32}$`
	PasswordLowerReg   = `[a-z]`
	PasswordUpperReg   = `[A-Z]`
	PasswordDigitReg   = `\d`
	PasswordSpecialReg = `[!@#$%^&*]`
)

const (
	DESC_CREATE_FAILED string = "create failed"
	DESC_GET_FAILED    string = "get failed"
	DESC_UPDATE_FAILED string = "update failed"
	DESC_DELETE_FAILED string = "delete failed"

	DESC_CREATE_SUCCESS string = "create successfully"
	DESC_GET_SUCCESS    string = "get successfully"
	DESC_UPDATE_SUCCESS string = "update successfully"
	DESC_DELETE_SUCCESS string = "delete successfully"

	DESC_NOT_FOUND_DATA string = "not found data"
	DESC_EMPTY_DATA     string = "empty data"
	DESC_INVALID_DATA   string = "invalid data"
)

// const (
// 	STATUS_OK              string = "OK"
// 	STATUS_CREATED         string = "Created"
// 	STATUS_ACCEPTED        string = "Accepted"
// 	STATUS_NO_CONTENT      string = "No Content"
// 	STATUS_RESET_CONTENT   string = "Reset Content"
// 	STATUS_PARTIAL_CONTENT string = "Partial Content"
// 	STATUS_MOVED_PERMANENT string = "Moved Permanently"
// 	STATUS_FOUND           string = "Found"
// 	STATUS_SEE_OTHER       string = "See Other"
// 	STATUS_NOT_MODIFIED    string = "Not Modified"
// 	STATUS_USE_PROXY       string = "Use Proxy"
// 	STATUS_TEMP_REDIRECT   string = "Temporary Redirect"

// 	STATUS_BAD_REQUEST          string = "Bad Request"
// 	STATUS_UNAUTHORIZED         string = "Unauthorized"
// 	STATUS_FORBIDDEN            string = "Forbidden"
// 	STATUS_NOT_FOUND            string = "Not Found"
// 	STATUS_METHOD_NOT_ALLOWED   string = "Method Not Allowed"
// 	STATUS_NOT_ACCEPTABLE       string = "Not Acceptable"
// 	STATUS_PROXY_AUTH_REQUIRED  string = "Proxy Authentication Required"
// 	STATUS_REQUEST_TIMEOUT      string = "Request Timeout"
// 	STATUS_CONFLICT             string = "Conflict"
// 	STATUS_GONE                 string = "Gone"
// 	STATUS_LENGTH_REQUIRED      string = "Length Required"
// 	STATUS_PRECONDITION_FAILED  string = "Precondition Failed"
// 	STATUS_REQUEST_ENTITY_LARGE string = "Request Entity Too Large"
// 	STATUS_REQUEST_URI_LARGE    string = "Request-URI Too Large"
// 	STATUS_UNSUPPORTED_MEDIA    string = "Unsupported Media Type"
// 	STATUS_REQUEST_RANGE_INVALID string = "Requested Range Not Satisfiable"
// 	STATUS_EXPECTATION_FAILED    string = "Expectation Failed"

// 	STATUS_INTERNAL_SERVER_ERROR string = "Internal Server Error"
// 	STATUS_NOT_IMPLEMENTED        string = "Not Implemented"
// 	STATUS_BAD_GATEWAY            string = "Bad Gateway"
// 	STATUS_SERVICE_UNAVAILABLE    string = "Service Unavailable"
// 	STATUS_GATEWAY_TIMEOUT        string = "Gateway Timeout"

// 	STATUS_CONTINUE           string = "Continue"
// 	STATUS_SWITCHING_PROTOCOL string = "Switching Protocols"
// 	STATUS_PROCESSING         string = "Processing"
// 	STATUS_EARLY_HINTS        string = "Early Hints"
// )

type IFontDouble map[string][3]string
var DOUBLE_ALPHA = IFontDouble{
	`A`: {
		`╔═╗`,
		`╠═╣`,
		`╩ ╩`,
	},
	`B`: {
		`╔╗ `,
		`╠╩╗`,
		`╚═╝`,
	},
	`C`: {
		`╔═╗`,
		`║  `,
		`╚═╝`,
	},
	`D`: {
		`╦═╗`,
		`║ ║`,
		`╩═╝`,
	},
	`E`: {
		`╔═╗`,
		`║╣ `,
		`╚═╝`,
	},
	`F`: {
		`╔═╗`,
		`║╣ `,
		`╩  `,
	},
	`G`: {
		`╔═╗`,
		`║ ╦`,
		`╚═╝`,
	},
	`H`: {
		`╦ ╦`,
		`╠═╣`,
		`╩ ╩`,
	},
	`I`: {
		` ╦ `,
		` ║ `,
		` ╩ `,
	},
	`J`: {
		`╦═╦`,
		`  ║`,
		` ═╝`,
	},
	`K`: {
		`╦  ╦`,
		`╠═╔╝`,
		`╝ ╚`,
	},
	`L`: {
		`╦  `,
		`║  `,
		`╩═╝`,
	},
	`M`: {
		`╔╗╔╗`,
		`║║║║`,
		`╝╚╝╚`,
	},
	`N`: {
		`╔╗╔`,
		`║║║`,
		`╝╚╝`,
	},
	`O`: {
		`╔═╗`,
		`║ ║`,
		`╚═╝`,
	},
	`P`: {
		`╔═╗`,
		`╠═╝`,
		`╩  `,
	},
	`Q`: {
		`╔═╗`,
		`║ ║`,
		`╚═╚═`,
	},
	`R`: {
		`╔═╗`,
		`╠╔╝`,
		`╝╚╝`,
	},
	`S`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`T`: {
		`╔╦╗`,
		` ║ `,
		` ╩ `,
	},
	`U`: {
		`╦ ╦`,
		`║ ║`,
		`╚═╝`,
	},
	`V`: {
		`╦   ╦`,
		`╚╗ ╔╝`,
		` ╚╔╝`,
	},
	`W`: {
		`╦ ╦`,
		`║║║`,
		`╚╩╝`,
	},
	`X`: {
		`╚╔╝`,
		`╔╝╗`,
	},
	`Y`: {
		`╦ ╦`,
		`╚╔╝`,
		` ╩ `,
	},
	`Z`: {
		`╔═╗`,
		`╔═╝`,
		`╚═╝`,
	},
	`0`: {
		`╔═╗`,
		`║ ║`,
		`╚═╝`,
	},
	`1`: {
		` ╦ `,
		` ║ `,
		` ╩ `,
	},
	`2`: {
		`╔═╗`,
		`╔═╝`,
		`╚══`,
	},
	`3`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`4`: {
		`╦ ╦`,
		`╚═╣`,
		`  ╩`,
	},
	`5`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`6`: {
		`╔═╗`,
		`╠═╗`,
		`╚═╝`,
	},
	`7`: {
		`╔═╗`,
		` ═║`,
		`  ╩`,
	},
	`8`: {
		`╔═╗`,
		`╠═╣`,
		`╚═╝`,
	},
	`9`: {
		`╔═╗`,
		`╚═╣`,
		`╚═╝`,
	},
}

// var zero =
//   `███
//    █ █
//    █ █
//    █ █
//    ███`

// var one =
//   `██
//     █
//     █
//     █
//    ███`

// var two =
//   `███
//      █
//    ███
//    █
//    ███`

// var three =
//   `███
//      █
//    ███
//      █
//    ███`

// var four =
//   `█ █
//    █ █
//    ███
//      █
//      █`

// var five =
//   `███
//    █
//    ███
//      █
//    ███`

// var six =
//   `███
//    █
//    ███
//    █ █
//    ███`

// var seven =
// `███
//     █
//     █
//     █
//     █`

// var eight =
//   `███
//    █ █
//    ███
//    █ █
//    ███`

// var nine =
//   `███
//    █ █
//    ███
//      █
//    ███`

// var colon =
//   `
//     ░

//     ░
//       `
