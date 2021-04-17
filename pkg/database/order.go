package database

import (
	"strings"
)

type (
	// Sort handle your sorters.
	Sort struct {
		sorters []sort
	}

	// sort refers to a order sentence it contains the field and the direction of the sorting.
	sort struct {
		field     string
		direction string
	}
)

// NewSort allow us to create a new Sort struct
func NewSort(params []string) Sort {
	var sorters []sort
	for _, value := range params {
		s := strings.Split(value, ",")
		f := s[0]
		d := s[1]
		sorters = append(sorters, sort{field: f, direction: d})
	}
	return Sort{sorters}
}
