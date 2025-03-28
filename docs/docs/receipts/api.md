
# <Badge style='font-size: 1.8rem; text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); padding: 0.35rem 0.75rem 0.35rem 0;' type='info' text='🔖 Api' />


```go
import (
	"context"
	"fmt"
	"net/http"
	"github.com/thuongtruong109/gouse"
)
```

## 1. Validate

Description: Validate request data<br>Input params: (ctxBind func() error, ctxReq func() context.Context, req any)<br>

```go
func Validate() {
	type UserRequest struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	req := UserRequest{
		Name:  "John",
		Email: "invalid-email",
	}

	ctxBind := func() error {
		return nil
	}

	ctxReq := func() context.Context {
		return context.Background()
	}

	err := gouse.Validate(ctxBind, ctxReq, req)
	if err != nil {
		gouse.Println("Validation failed:", err)
	} else {
		gouse.Println("Validation succeeded!")
	}
}
```

## 2. Upload file

Description: Upload file to server<br>Input params: (servePath, storagePath string)<br>

```go
func UploadFile() {
	http.HandleFunc("/upload-single", func(w http.ResponseWriter, r *http.Request) {
		gouse.UploadSingle(w, r, "http://localhost:8080/img", "./public")
	})
	http.HandleFunc("/upload-multi", func(w http.ResponseWriter, r *http.Request) {
		gouse.UploadMulti(w, r, "http://localhost:8080/img", "./public")
	})

	http.ListenAndServe(":8080", nil)
}
```

## 3. Header

Description: Return the header of the request<br>Input params: (w http.ResponseWriter, r *http.Request)<br>

```go
func Header() {
	http.HandleFunc("/header", gouse.Header)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
```

## 4. Pagination

Description: Setup pagination<br>Input params: (page, size int)<br>

```go
func Pagination() {
	pagination := gouse.NewPagination(10, 1)

	pagination.SetPage("2")
	pagination.SetSize("15")
	pagination.SetOrderBy("desc")

	pagination.Total = 100

	fmt.Println("Pagination Info:")
	fmt.Println("Page:", pagination.GetPage())
	fmt.Println("Size per Page:", pagination.GetSize())
	fmt.Println("Order By:", pagination.GetOrderBy())
	fmt.Println("Total Pages:", pagination.GetTotalPages())
	fmt.Println("Next Page:", pagination.GetNextPage())
	fmt.Println("Prev Page:", pagination.GetPrevPage())
	fmt.Println("Has More Pages:", pagination.GetHasMore(pagination.Total))
	fmt.Println("Offset:", pagination.GetOffset())
	fmt.Println("First Page:", pagination.GetFirstPage())
	fmt.Println("Last Page:", pagination.GetLastPage())
	fmt.Println("Total Items:", pagination.Total)
	fmt.Println("Query String:", pagination.GetQueryString())
}
```
