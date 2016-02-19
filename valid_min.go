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
		return ErrUnsupportedType
	}

	var num int64
	if num, err = format.ToInt64(param); err != nil {
		return ErrUnsupportedType
	}

	if num > val {
		return ErrMinNumber
	}
	return
}
