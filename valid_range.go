package validation

import (
	"strings"

	"github.com/go-rut/format"
)

type Range struct{}

var (
	tmpRange *Range
)

func NewRange() *Range {
	return tmpRange
}

func (p *Range) IsValid(v interface{}, param string) (err error) {

	tokens := strings.Split(param, ",")
	if len(tokens) != 2 {
		return ErrBadParameter
	}

	var val int64
	if val, err = format.ToInt64(v); err != nil {
		return ErrUnsupportedType
	}

	var p1, p2 int64
	if p1, err = format.ToInt64(tokens[0]); err != nil {
		return ErrUnsupportedType
	}
	if p2, err = format.ToInt64(tokens[1]); err != nil {
		return ErrUnsupportedType
	}

	if p1 > p2 {
		p1, p2 = p2, p1
	}

	if val < p1 {
		return ErrRangeLower
	}

	if val > p2 {
		return ErrRangeUpper
	}

	return
}
