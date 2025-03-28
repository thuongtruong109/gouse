package gouse

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpError(t *testing.T) {
	_, err := http.NewRequest("GET", "/img", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	httpError(rr, "Not Found", http.StatusNotFound)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, status)
	}

	expected := "Not Found"
	actual := strings.TrimSpace(rr.Body.String())

	if actual != expected {
		t.Errorf("Expected body %q, got %q", expected, actual)
	}
}

func TestPagination(t *testing.T) {
	p := NewPagination(10, 1)

	if p.Size != 10 {
		t.Errorf("expected size 10, got %d", p.Size)
	}
	if p.Page != 1 {
		t.Errorf("expected page 1, got %d", p.Page)
	}

	// Test default values when invalid inputs are provided
	p2 := NewPagination(0, 0)
	if p2.Size != DefaultPageSize {
		t.Errorf("expected default size, got %d", p2.Size)
	}
	if p2.Page != 1 {
		t.Errorf("expected default page 1, got %d", p2.Page)
	}

	tests := []struct {
		sizeQuery string
		expected  int
		expectErr bool
	}{
		{"", DefaultPageSize, false},
		{"20", 20, false},
		{"abc", 0, true},
	}

	for _, tt := range tests {
		p := &Pagination{}
		err := p.SetSize(tt.sizeQuery)

		if (err != nil) != tt.expectErr {
			t.Errorf("SetSize(%v) expected error: %v, got: %v", tt.sizeQuery, tt.expectErr, err)
		}

		if p.Size != tt.expected {
			t.Errorf("SetSize(%v) = %v, want %v", tt.sizeQuery, p.Size, tt.expected)
		}
	}

	tests2 := []struct {
		pageQuery string
		expected  int
		expectErr bool
	}{
		{"", 1, false},
		{"5", 5, false},
		{"abc", 0, true},
	}

	for _, tt := range tests2 {
		p := &Pagination{}
		err := p.SetPage(tt.pageQuery)

		if (err != nil) != tt.expectErr {
			t.Errorf("SetPage(%v) expected error: %v, got: %v", tt.pageQuery, tt.expectErr, err)
		}

		if p.Page != tt.expected {
			t.Errorf("SetPage(%v) = %v, want %v", tt.pageQuery, p.Page, tt.expected)
		}
	}

	tests3 := []struct {
		total    int
		size     int
		expected int
	}{
		{100, 10, 10},
		{105, 10, 11},
		{50, 10, 5},
		{100, 20, 5},
	}

	for _, tt := range tests3 {
		p := &Pagination{Total: tt.total, Size: tt.size}
		actual := p.GetTotalPages()

		if actual != tt.expected {
			t.Errorf("GetTotalPages() = %v, want %v", actual, tt.expected)
		}
	}

	tests4 := []struct {
		page     int
		total    int
		size     int
		expected int
	}{
		{1, 100, 10, 2},
		{1, 100, 20, 2},
		{10, 100, 10, 10}, // Last page case
		{5, 100, 10, 6},
	}

	for _, tt := range tests4 {
		p := &Pagination{Page: tt.page, Total: tt.total, Size: tt.size}
		actual := p.GetNextPage()

		if actual != tt.expected {
			t.Errorf("GetNextPage() = %v, want %v", actual, tt.expected)
		}
	}

	tests5 := []struct {
		page     int
		expected int
	}{
		{1, 1}, // Cannot go below page 1
		{2, 1},
		{10, 9},
	}

	for _, tt := range tests5 {
		p := &Pagination{Page: tt.page}
		actual := p.GetPrevPage()

		if actual != tt.expected {
			t.Errorf("GetPrevPage() = %v, want %v", actual, tt.expected)
		}
	}

	tests6 := []struct {
		page     int
		size     int
		expected int
	}{
		{1, 10, 0},  // First page, offset should be 0
		{2, 10, 10}, // Second page, offset should be size
		{3, 20, 40}, // Third page, offset should be size * 2
	}

	for _, tt := range tests6 {
		p := &Pagination{Page: tt.page, Size: tt.size}
		actual := p.GetOffset()

		if actual != tt.expected {
			t.Errorf("GetOffset() = %v, want %v", actual, tt.expected)
		}
	}

	p7 := &Pagination{
		Page:    2,
		Size:    10,
		OrderBy: "asc",
	}
	expected := "page=2&size=10&orderBy=asc"
	actual := p7.GetQueryString()

	if actual != expected {
		t.Errorf("GetQueryString() = %v, want %v", actual, expected)
	}

	tests8 := []struct {
		pagination      Pagination
		totalCount      int
		expectedHasMore bool
	}{
		{
			pagination:      Pagination{Page: 1, Size: 10},
			totalCount:      25,
			expectedHasMore: true, // More pages exist after page 1
		},
		{
			pagination:      Pagination{Page: 1, Size: 10},
			totalCount:      10,
			expectedHasMore: false, // Only one page available
		},
		// {
		// 	pagination: Pagination{Page: 2, Size: 10},
		// 	totalCount: 25,
		// 	expectedHasMore: false, // No more pages after page 2
		// },
		{
			pagination:      Pagination{Page: 0, Size: 10},
			totalCount:      0,
			expectedHasMore: false, // No data, no pages
		},
		// {
		// 	pagination: Pagination{Page: 1, Size: 5},
		// 	totalCount: 12,
		// 	expectedHasMore: true, // More pages exist after page 1
		// },
	}

	for _, tt := range tests8 {
		t.Run("", func(t *testing.T) {
			hasMore := tt.pagination.GetHasMore(tt.totalCount)
			if hasMore != tt.expectedHasMore {
				t.Errorf("expected %v, got %v", tt.expectedHasMore, hasMore)
			}
		})
	}
}
