// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"github.com/go-trellis/formats"
	"reflect"
	"unicode/utf8"
)

type Max struct{}

var (
	max *Max
)

func NewMax() *Max {
	return max
}

func (p *Max) IsValid(v interface{}, param string) (err error) {
	vT := reflect.TypeOf(v)
	switch vT.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var val, num int64
		if num, err = formats.ToInt64(param); err != nil {
			return ErrUnsupportedType
		}

		if val, err = formats.ToInt64(v); err != nil {
			return ErrUnsupportedType
		}
		if num < val {
			return ErrMinNumber
		}
		return nil

	case reflect.String:
		var val, num int

		//val = len([]rune(v.(string)))
		val = utf8.RuneCountInString(v.(string))

		if num, err = formats.ToInt(param); err != nil {
			return ErrUnsupportedType
		}

		if num < val {
			return ErrMinNumber
		}
		return nil

	case reflect.Slice, reflect.Map, reflect.Array:
		var val, num int
		val = reflect.ValueOf(v).Len()

		if num, err = formats.ToInt(param); err != nil {
			return ErrUnsupportedType
		}

		if num < val {
			return ErrMinNumber
		}
		return nil
	default:
		return ErrUnsupportedType
	}

}

/*
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
*/
