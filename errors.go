package validation

import "errors"

var (
	ErrTagNameNotExists = errors.New("tag name not exists")
	ErrUnsupportedType  = errors.New("unsupported valid type")
	ErrUnknownFuncName  = errors.New("unknown valid function name")

	ErrBadParameter = errors.New("bad parameter")
	ErrRegexCompile = errors.New("compile invalid parameter")

	ErrCellphone  = errors.New("cellphone is invalid")
	ErrEmail      = errors.New("email is invalid")
	ErrCreditCard = errors.New("creadit card is invalid")
	ErrRegexMatch = errors.New("failed validate regex")
	ErrLength     = errors.New("length is invalid")
	ErrMinNumber  = errors.New("min number is invalid")
	ErrMaxNumber  = errors.New("max number is invalid")
	ErrRangeUpper = errors.New("out of range, greater than upper limit")
	ErrRangeLower = errors.New("out of range, less than lower limit")
)
