package gouse

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/* Port */

func ScanPort(protocol, hostname string, start, end int) {
	for i := start; i < end; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial(protocol, hostname+":"+port)
		if err == nil {
			fmt.Println("Port", i, "open")
			conn.Close()
		} else {
			fmt.Println("Port", i, "closed")
		}
	}
}

func CheckPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	return true
}

/* Proxy */

func handleRequest(urls []string) gin.HandlerFunc {
	var counter uint64

	return func(ctx *gin.Context) {
		path := ctx.Param("path")
		if path == "" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Path is required",
			})
			ctx.Done()
			return
		}

		index := atomic.AddUint64(&counter, 1) % uint64(len(urls))
		requestedURL := urls[index] + path[1:]

		fmt.Println("Requested URL: ", requestedURL)

		req, _ := http.NewRequest(ctx.Request.Method, requestedURL, ctx.Request.Body)

		req.Header = ctx.Request.Header.Clone()
		req.Header.Del("origin")
		req.Header.Del("referer")

		queryValues := req.URL.Query()
		for k, v := range ctx.Request.URL.Query() {
			queryValues.Add(k, v[0])
		}
		req.URL.RawQuery = queryValues.Encode()

		response, err1 := http.DefaultClient.Do(req)

		for k, v := range response.Header.Clone() {
			ctx.Header(k, v[0])
		}

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")

		responseBytes, err2 := io.ReadAll(response.Body)

		if err1 != nil || err2 != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to request",
			})
			ctx.Done()
			return
		}

		ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), responseBytes)
	}
}

func Proxy(port string, urls []string) {
	router := gin.Default()

	router.GET("*path", handleRequest(urls))
	router.POST("*path", handleRequest(urls))
	router.PUT("*path", handleRequest(urls))
	router.PATCH("*path", handleRequest(urls))
	router.DELETE("*path", handleRequest(urls))
	router.OPTIONS("*path", handleRequest(urls))
	router.HEAD("*path", handleRequest(urls))

	router.Run(":" + port)
}

/* Validator */

func ReadRequest(ctxBind func() error, ctxReq func() context.Context, req interface{}) error {
	validate := validator.New()

	if err := ctxBind(); err != nil {
		return err
	}

	ctx := ctxReq()

	return validate.StructCtx(ctx, req)
}

/* Href link */

func OpenHref(url string) {
	Cmd("explorer " + url)
}

func EncodeHref(s string) string {
	return url.QueryEscape(s)
}

func DecodeHref(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}

func CheckHref(url string) (bool, error) {
	_, err := http.Get(url)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckHrefStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}

func HrefHeader(url string) (http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

func HrefConnectTime(url string) (float64, error) {
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	elapsedTime := time.Since(startTime)

	if err := resp.Body.Close(); err != nil {
		return 0, err
	}
	return elapsedTime.Seconds(), nil
}