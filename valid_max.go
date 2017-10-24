// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"github.com/go-trellis/formats"
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
	if val, err = formats.ToInt64(v); err != nil {
		return ErrUnsupportedType
	}

	var num int64
	if num, err = formats.ToInt64(param); err != nil {
		return ErrUnsupportedType
	}

	if num < val {
		return ErrMaxNumber
	}
	return
}
