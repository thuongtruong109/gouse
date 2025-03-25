package gouse

import (
	"fmt"
	"log"
)

func DetectErr(err any) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return fmt.Sprintf("%v", err)
	}
}

func ToErr(err any) error {
	switch e := err.(type) {
	case error:
		return e
	case string:
		return fmt.Errorf("%s", e)
	default:
		return fmt.Errorf("%v", err)
	}
}

func Panic(msg error, err ...any) {
	if err != nil {
		er := DetectErr(err)
		panic(fmt.Sprintf("%s: %s\n", msg, er))
	}
}

func Fatal(msg error, err ...any) {
	if err != nil {
		er := DetectErr(err)
		log.Fatalf("%s: %s\n", msg, er)
	}
}
