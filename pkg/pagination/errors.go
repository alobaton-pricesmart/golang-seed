package pagination

import "fmt"

var (
	ErrRequiredPage = fmt.Errorf("page is required")
	ErrRequiredSize = fmt.Errorf("size is required")
	ErrInvalidPage  = fmt.Errorf("page is invalid")
	ErrInvalidSize  = fmt.Errorf("size is invalid")
)
