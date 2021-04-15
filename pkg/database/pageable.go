package database

type (
	Pageable struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	Page struct {
		Pageable      Pageable    `json:"pageable"`
		TotalPages    int         `json:"totalPages"`
		TotalElements int         `json:"totalElements"`
		Content       interface{} `json:"content"`
	}
)

func NewPageable(page, size int) *Pageable {
	return &Pageable{Page: page, Size: size}
}
