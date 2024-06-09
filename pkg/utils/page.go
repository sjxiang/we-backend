package utils

import (
	"errors"
	"strconv"
	
	"we-backend/pkg/validate"
)

// Page represents the requested page and rows per page.
type Page struct {
	number int
	rows   int
}

// Parse parses the strings and validates the values are in reason.
func Parse(page string, rowsPerPage string) (*Page, error) {
	// 默认值
	number := 1
	if page != "" {
		var err error
		number, err = strconv.Atoi(page)  // 畸形
		if err != nil {
			return nil, validate.NewFieldsError("page", err)
		}
	}

	// 默认值
	rows := 10
	if rowsPerPage != "" {
		var err error
		rows, err = strconv.Atoi(rowsPerPage)
		if err != nil {
			return nil, validate.NewFieldsError("rows", err)
		}
	}

	if number <= 0 {  // 不合理、无效 
		return nil, validate.NewFieldsError("page", errors.New("page value too small, must be larger than 0"))
	}

	if rows <= 0 {
		return nil, validate.NewFieldsError("rows", errors.New("rows value too small, must be larger than 0"))
	}

	if rows > 100 {
		return nil, validate.NewFieldsError("rows", errors.New("rows value too large, must be less than 100"))
	}

	p := Page{
		number: number,
		rows:   rows,
	}

	return &p, nil
}


// Number returns the page number.
func (p Page) Number() int {
	return p.number
}

// RowsPerPage returns the rows per page.
func (p Page) RowsPerPage() int {
	return p.rows
}
