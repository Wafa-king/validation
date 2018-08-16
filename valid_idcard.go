// MIT License

// Copyright (c) 2016 rutcode-go

package validation

const (
	// 15位
	//fivteen = `^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$`
	// 18位
	//eight = `^[1-9]\d{5} (18|19|([23]\d))\d {2}((0[1-9])|(10|11|12)) (([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$"`
	// 15位&&18位
	IdCardPattern = `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$`
)

type IdCard struct{}

var (
	idCard *IdCard
)

func NewIdCard() *IdCard {
	return idCard
}

func (p *IdCard) IsValid(v interface{}, _ string) (err error) {
	err = NewRegex().IsValid(v, IdCardPattern)
	if err != nil {
		return ErrCreditCard
	}
	return
}
