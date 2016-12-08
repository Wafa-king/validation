// MIT License

// Copyright (c) 2016 rutcode-go

package validation

const (
	emailPattern = "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"
)

type Email struct{}

var (
	email *Email
)

func NewEmail() *Email {
	return email
}

func (p *Email) IsValid(v interface{}, _ string) (err error) {
	err = NewRegex().IsValid(v, emailPattern)
	if err != nil {
		return ErrEmail
	}
	return
}
