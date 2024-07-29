package net

import (
	"fmt"
	"io"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var counter uint64

func handleRequest(urls []string) gin.HandlerFunc {
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
