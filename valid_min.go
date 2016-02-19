package validation

import (
	"github.com/go-rut/format"
)

type Min struct{}

var (
	min *Min
)

func NewMin() *Min {
	return min
}

func (p *Min) IsValid(v interface{}, param string) (err error) {

	var val int64
	if val, err = format.ToInt64(v); err != nil {
		return ErrBadParameter
	}

	var num int64
	if num, err = format.ToInt64(param); err != nil {
		return ErrBadParameter
	}

	if num > val {
		return ErrMinNumber
	}
	return
}
