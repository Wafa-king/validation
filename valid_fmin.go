// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"github.com/go-rut/format"
)

type FloatMin struct{}

var (
	floatMin *FloatMin
)

func NewFloatMin() *FloatMin {
	return floatMin
}

func (p *FloatMin) IsValid(v interface{}, param string) (err error) {

	var val float64
	if val, err = format.ToFloat64(v); err != nil {
		return ErrUnsupportedType
	}

	var num float64
	if num, err = format.ToFloat64(param); err != nil {
		return ErrUnsupportedType
	}

	if num > val {
		return ErrMinNumber
	}
	return
}
