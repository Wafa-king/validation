// MIT License

// Copyright (c) 2016 rutcode-go

package validation

const (
	creditCardPattern = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
)

type CreditCard struct{}

var (
	creditCard *CreditCard
)

func NewCreditCard() *CreditCard {
	return creditCard
}

func (p *CreditCard) IsValid(v interface{}, _ string) (err error) {
	err = NewRegex().IsValid(v, creditCardPattern)
	if err != nil {
		return ErrCreditCard
	}
	return
}
