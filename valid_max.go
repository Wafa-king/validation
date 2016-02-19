package validation

import (
	"github.com/go-rut/format"
)

type Max struct{}

var (
	max *Max
)

func NewMax() *Max {
	return max
}

func (p *Max) IsValid(v interface{}, param string) (err error) {

	var val int64
	if val, err = format.ToInt64(v); err != nil {
		return ErrUnsupportedType
	}

	var num int64
	if num, err = format.ToInt64(param); err != nil {
		return ErrUnsupportedType
	}

	if num < val {
		return ErrMaxNumber
	}
	return
}
