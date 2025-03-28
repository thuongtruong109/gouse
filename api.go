package gouse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func Validate(ctxBind func() error, ctxReq func() context.Context, req any) error {
	validate := validator.New()

	if err := ctxBind(); err != nil {
		return err
	}

	ctx := ctxReq()

	return validate.StructCtx(ctx, req)
}

func httpError(w http.ResponseWriter, msg string, statusCode int) {
	http.Error(w, msg, statusCode)
}

func generateFileName(originalFileName string) string {
	fileExt := filepath.Ext(originalFileName)
	originalName := strings.TrimSuffix(filepath.Base(originalFileName), fileExt)
	now := time.Now().UnixNano()
	filename := fmt.Sprintf("%s-%d%s", strings.ReplaceAll(strings.ToLower(originalName), " ", "-"), now, fileExt)
	return filename
}

func setJSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func UploadSingle(w http.ResponseWriter, r *http.Request, servePath, storePath string) {
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		httpError(w, fmt.Sprintf("Error parsing form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		httpError(w, fmt.Sprintf("Error retrieving file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := generateFileName(header.Filename)
	filePath := filepath.Join(servePath, filename)

	out, err := os.Create(filepath.Join(storePath, filename))
	if err != nil {
		httpError(w, fmt.Sprintf("Unable to create file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		httpError(w, fmt.Sprintf("Error writing file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	setJSONHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(fmt.Appendf(nil, `{"filepath": "%s"}`, filePath))
}

func UploadMulti(w http.ResponseWriter, r *http.Request, servePath, storePath string) {
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		httpError(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		httpError(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	var filePaths []string

	for _, file := range files {
		filename := generateFileName(file.Filename)
		filePath := filepath.Join(servePath, filename)
		filePaths = append(filePaths, filePath)

		out, err := os.Create(filepath.Join(storePath, filename))
		if err != nil {
			httpError(w, fmt.Sprintf("Error saving file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer out.Close()

		readerFile, err := file.Open()
		if err != nil {
			httpError(w, fmt.Sprintf("Error opening file: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer readerFile.Close()

		_, err = io.Copy(out, readerFile)
		if err != nil {
			httpError(w, fmt.Sprintf("Error copying file content: %s", err.Error()), http.StatusInternalServerError)
			return
		}
	}

	setJSONHeader(w)
	w.WriteHeader(http.StatusOK)
	w.Write(fmt.Appendf(nil, `{"filepath": %v}`, filePaths))
}

type IHeader struct {
	Method        string
	Addr          string
	URL           string
	Proto         string
	Host          string
	StatusCode    int
	ContentLength int
	ContentType   string
}

func Header(w http.ResponseWriter, r *http.Request) {
	setJSONHeader(w)

	for k, v := range r.Header {
		fmt.Fprintf(w, "%q: %q\n\n", k, v)
	}

	p := &IHeader{
		Method:        r.Method,
		Addr:          r.RemoteAddr,
		URL:           r.URL.String(),
		Proto:         r.Proto,
		Host:          r.Host,
		StatusCode:    200,
		ContentLength: 200,
		ContentType:   "application/json",
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(p)
}

const DefaultPageSize = 10

type Pagination struct {
	Size    int    // Items per page
	Page    int    // Current page
	OrderBy string // Sorting order (e.g., "asc", "desc")
	Total   int    // Total number of items available
}

func NewPagination(size, page int) *Pagination {
	if size <= 0 {
		size = DefaultPageSize
	}
	if page <= 0 {
		page = 1
	}
	return &Pagination{
		Page: page,
		Size: size,
	}
}

func (p *Pagination) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		p.Size = DefaultPageSize
		return nil
	}
	size, err := strconv.Atoi(sizeQuery)
	if err != nil || size <= 0 {
		return fmt.Errorf("invalid page size: %v", err)
	}
	p.Size = size
	return nil
}

func (p *Pagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		p.Page = 1
		return nil
	}
	page, err := strconv.Atoi(pageQuery)
	if err != nil || page <= 0 {
		return fmt.Errorf("invalid page number: %v", err)
	}
	p.Page = page
	return nil
}

func (p *Pagination) SetOrderBy(orderByQuery string) {
	p.OrderBy = orderByQuery
}

func (p *Pagination) GetTotalPages() int {
	if p.Size == 0 {
		return 0
	}
	return int(math.Ceil(float64(p.Total) / float64(p.Size)))
}

func (p *Pagination) GetNextPage() int {
	if p.Page < p.GetTotalPages() {
		return p.Page + 1
	}
	return p.Page
}

func (p *Pagination) GetPrevPage() int {
	if p.Page > 1 {
		return p.Page - 1
	}
	return p.Page
}

func (p *Pagination) GetFirstPage() int {
	return 1
}

func (p *Pagination) GetLastPage() int {
	return p.GetTotalPages()
}

func (p *Pagination) GetPage() int {
	return p.Page
}

func (p *Pagination) GetOrderBy() string {
	return p.OrderBy
}

func (p *Pagination) GetOffset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *Pagination) GetSize() int {
	return p.Size
}

func (p *Pagination) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", p.GetPage(), p.GetSize(), p.GetOrderBy())
}

func (p *Pagination) GetHasMore(totalCount int) bool {
	totalPages := int(math.Ceil(float64(totalCount) / float64(p.Size)))
	return p.Page < totalPages
}
