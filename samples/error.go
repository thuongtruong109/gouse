package samples

import (
	"fmt"

	"github.com/thuongtruong109/gouse"
)

func ErrorDetect() {
	err1 := fmt.Errorf("this is an error")
	fmt.Println("Error 1:", gouse.DetectErr(err1))

	err2 := "this is a string error"
	fmt.Println("Error 2:", gouse.DetectErr(err2))

	err3 := 12345
	fmt.Println("Error 3:", gouse.DetectErr(err3))

	var err4 error = nil
	fmt.Println("Error 4:", gouse.DetectErr(err4))
}

func ErrorFormat() {
	err1 := fmt.Errorf("this is an error")
	if e := gouse.Err(err1); e != nil {
		fmt.Println("Error 1:", e)
	}

	err2 := "this is a string error"
	if e := gouse.Err(err2); e != nil {
		fmt.Println("Error 2:", e)
	}

	err3 := 12345
	if e := gouse.Err(err3); e != nil {
		fmt.Println("Error 3:", e)
	}

	var err4 error = nil
	if e := gouse.Err(err4); e != nil {
		fmt.Println("Error 4:", e)
	} else {
		fmt.Println("Error 4: No error")
	}
}

func ErrorMsg() {
	err1 := fmt.Errorf("file not found")
	customErr1 := gouse.ErrMsg("Failed to open file", err1)
	fmt.Println("Custom Error 1:", customErr1)

	err2 := fmt.Errorf("user not authorized")
	customErr2 := gouse.ErrMsg("Authentication failed", err2)
	fmt.Println("Custom Error 2:", customErr2)

	customErr3 := gouse.ErrMsg("Unexpected error occurred", nil)
	fmt.Println("Custom Error 3:", customErr3)

	customErr4 := gouse.ErrMsg("Operation failed", fmt.Errorf(""))
	fmt.Println("Custom Error 4:", customErr4)
}

func ErrorPanic() {
	err1 := "something went wrong"
	gouse.Panic(fmt.Errorf("critical error"), err1)

	err2 := fmt.Errorf("unable to connect to the database")
	gouse.Panic(fmt.Errorf("database error"), err2)

	err3 := fmt.Errorf("network timeout")
	err4 := "retries exceeded"
	gouse.Panic(fmt.Errorf("network error"), err3, err4)

	gouse.Panic(fmt.Errorf("no error scenario"))
}

func ErrorFatal() {
	err1 := "file not found"
	gouse.Fatal(fmt.Errorf("file error"), err1)

	err2 := fmt.Errorf("unable to connect to the database")
	gouse.Fatal(fmt.Errorf("database error"), err2)

	err3 := fmt.Errorf("network timeout")
	err4 := "retries exceeded"
	gouse.Fatal(fmt.Errorf("network error"), err3, err4)

	gouse.Fatal(fmt.Errorf("no error scenario"))
}
