package database

import (
	"strings"
)

type (
	Sort struct {
		sorters []sort
	}

	sort struct {
		field     string
		direction string
	}
)

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
