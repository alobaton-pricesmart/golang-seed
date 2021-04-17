package database

import "math"

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

func totalPages(totalElements int, size int) int {
	return int(math.Ceil(float64(totalElements) / float64(size)))
}
