// MIT License

// Copyright (c) 2016 rutcode-go

package validation

import (
	"regexp"
)

type Regex struct{}

var (
	regex *Regex
)

func NewRegex() *Regex {
	return regex
}

func (*Regex) IsValid(v interface{}, param string) error {
	s, ok := v.(string)
	if !ok {
		return ErrUnsupportedType
	}

	re, err := regexp.Compile(param)
	if err != nil {
		return ErrRegexCompile
	}

	if !re.MatchString(s) {
		return ErrRegexMatch
	}

	return nil
}
