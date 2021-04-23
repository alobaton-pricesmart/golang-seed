package pagination

import (
	"math"
	"net/http"
	"strconv"
)

type (
	// Pageable request
	Pageable struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	// Page response
	Page struct {
		Pageable      Pageable    `json:"pageable"`
		TotalPages    int         `json:"totalPages"`
		TotalElements int         `json:"totalElements"`
		Content       interface{} `json:"content"`
	}
)

// NewPageable allow us to build a new Pageable request
func NewPageable(page, size int) Pageable {
	return Pageable{Page: page, Size: size}
}

// NewPage creates a Page
func NewPage(pageable Pageable, totalElements int, content interface{}) *Page {
	return &Page{
		Pageable:      pageable,
		TotalPages:    totalPages(totalElements, pageable.Size),
		TotalElements: totalElements,
		Content:       content,
	}
}

func Pageabler(r *http.Request) (Pageable, error) {
	def := NewPageable(0, 10)

	if len(r.URL.Query()["page"]) < 1 {
		return def, ErrRequiredPage
	}

	if len(r.URL.Query()["size"]) < 1 {
		return def, ErrRequiredSize
	}

	var err error
	var pagep int
	var sizep int

	pagep, err = strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil {
		return def, ErrInvalidPage
	}

	sizep, err = strconv.Atoi(r.URL.Query()["size"][0])
	if err != nil {
		return def, ErrInvalidSize
	}

	return NewPageable(pagep, sizep), nil
}

func totalPages(totalElements int, size int) int {
	return int(math.Ceil(float64(totalElements) / float64(size)))
}
