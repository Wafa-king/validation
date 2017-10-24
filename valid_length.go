// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"reflect"

	"github.com/go-trellis/formats"
)

type Length struct{}

var (
	length *Length
)

func NewLength() *Length {
	return length
}

func (*Length) IsValid(v interface{}, param string) (err error) {

	var num int64
	var valid bool
	if num, err = formats.ToInt64(param); err != nil {
		return ErrUnsupportedType
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.String:
		valid = int64(len(val.String())) == num
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = int64(val.Len()) == num
	default:
		return ErrUnsupportedType
	}

	if !valid {
		return ErrLength
	}

	return
}
