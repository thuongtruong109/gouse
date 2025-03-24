package gouse

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"net/url"
)

func ScanPort(protocol, hostname string, start, end int) {
	for i := start; i < end; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial(protocol, hostname+":"+port)
		if err == nil {
			Println("Port", i, "open")
			conn.Close()
		} else {
			Println("Port", i, "closed")
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

func OpenUrl(url string) {
	Cmd("explorer " + url)
}

func EncodeUrl(s string) string {
	return url.QueryEscape(s)
}

func DecodeUrl(s string) string {
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}

func CheckUrl(url string) (bool, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, 0, err
	}

	return true, resp.StatusCode, nil
}

func HeaderUrl(url string) (http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

func ConTimeUrl(url string) (float64, error) {
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
