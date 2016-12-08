// MIT License

// Copyright (c) 2016 rutcode-go

package validation

const (
	cellphone86Pattern = `^1[3|4|5|8][0-9]\d{4,8}$`
)

type Cellphone86 struct{}

var (
	cellphone86 *Cellphone86
)

func NewCellphone86() *Cellphone86 {
	return cellphone86
}

func (*Cellphone86) IsValid(v interface{}, _ string) (err error) {
	err = NewRegex().IsValid(v, cellphone86Pattern)
	if err != nil {
		return ErrCellphone
	}
	return
}
