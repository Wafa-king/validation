package validation

import (
	"reflect"
	"strings"
	"unicode"
)

const (
	TagName = "valid"
)

type Validation struct {
	validFuncs map[string]Validator
}

type tagFunc struct {
	Name  string
	Param string
	Valid ValidFunc
}

type ValidFunc func(data interface{}, param string) (err error)

type Validator interface {
	IsValid(data interface{}, param string) (err error)
}

var (
	validation *Validation
)

func init() {
	validation = new(Validation)
	validation.validFuncs = map[string]Validator{
		"cellphone86": NewCellphone86(),
		"regex":       NewRegex(),
		"length":      NewLength(),
		"max":         NewMax(),
		"min":         NewMin(),
		"range":       NewRange(),
		"email":       NewEmail(),
		"creditcard":  NewCreditCard(),
	}
}

func Validate(s interface{}) error {
	return validation.Valid(s)
}

func (p *Validation) Valid(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		return p.Valid(val.Elem().Interface())
	}
	if val.Kind() != reflect.Struct {
		return ErrUnsupportedType
	}

	st := reflect.TypeOf(s)
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		if f.Kind() == reflect.Ptr && !f.IsNil() {
			f = f.Elem()
		}

		tagStr := strings.Trim(st.Field(i).Tag.Get(TagName), " ")
		if tagStr == "-" {
			continue
		}
		if tagStr == "" {
			if f.Kind() == reflect.Struct {
				return p.Valid(f.Interface())
			}
			continue
		}

		fieldName := st.Field(i).Name
		if !unicode.IsUpper(rune(fieldName[0])) {
			continue
		}

		funcs, err := p.getFuncs(tagStr)
		if err != nil {
			return err
		}

		data := f.Interface()
		for _, v := range funcs {
			if err = v.Valid(data, v.Param); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Validation) getFuncs(tagStr string) (funcs []tagFunc, err error) {
	tFuncs := strings.Split(tagStr, ";")

	for _, v := range tFuncs {
		vs := strings.SplitN(v, "=", 2)

		tf := tagFunc{}

		tf.Name = strings.Trim(vs[0], " ")
		if tf.Name == "" {
			return nil, ErrUnknownFuncName
		}

		if len(vs) > 1 {
			tf.Param = strings.Trim(vs[1], " ")
		}

		validation := p.validFuncs[tf.Name]
		if validation == nil {
			return nil, ErrUnknownFuncName
		}
		tf.Valid = validation.IsValid

		funcs = append(funcs, tf)
	}
	return
}
