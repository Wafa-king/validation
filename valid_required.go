// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"reflect"
)

type Required struct{}

var required *Required

func NewRequired() *Required {
	return required
}

func (*Required) IsValid(value interface{}, _ string) error {

	if IsZero(value) {
		return ErrNil
	}
	return nil
}

func IsZero(value interface{}) bool {

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}

	return false
}
