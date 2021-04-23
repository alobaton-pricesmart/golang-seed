package sorting

import (
	"net/http"
	"strings"
)

type (
	// Sort handle your sorters.
	Sort struct {
		Sorters []Sorter
	}

	// sort refers to a order sentence it contains the field and the direction of the sorting.
	Sorter struct {
		Field     string
		Direction string
	}
)

// NewSort allow us to create a new Sort struct
func NewSort(params []string) Sort {
	var sorters []Sorter
	for _, value := range params {
		s := strings.Split(value, ",")
		f := s[0]
		d := s[1]
		sorters = append(sorters, Sorter{Field: f, Direction: d})
	}
	return Sort{sorters}
}

func Sortr(r *http.Request) Sort {
	sortp := r.URL.Query()["sort"]
	return NewSort(sortp)
}
