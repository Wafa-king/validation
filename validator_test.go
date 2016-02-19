package validation_test

import (
	"testing"

	"github.com/go-rut/validation"
)

func TestCellphone86(t *testing.T) {

	type CPs struct {
		Cellphone86 string `valid:"cellphone86"`
	}

	cp := &CPs{
		Cellphone86: "11",
	}

	var err error

	err = validation.Validate(cp)

	if err == nil {
		t.Error("expect:" + validation.ErrCellphone.Error())
	}

	cp.Cellphone86 = "13000000000"

	err = validation.Validate(cp)
	if err != nil {
		t.Error(err)
	}
}

func TestLength(t *testing.T) {

	type CPs1 struct {
		Digit string `valid:"length=1"`
	}

	type CPs2 struct {
		Cellphone86 string `valid:"length=11;cellphone86"`
	}

	type CPs3 struct {
		Digit       int    `valid:"length=1"`
		Cellphone86 string `valid:"length=11;cellphone86"`
	}

	cp1 := &CPs1{
		Digit: "10",
	}

	var err error

	err = validation.Validate(cp1)
	if err == nil {
		t.Error("expect:" + validation.ErrLength.Error())
	}

	cp1.Digit = "1"
	err = validation.Validate(cp1)
	if err != nil {
		t.Error(err)
	}

	cp2 := &CPs2{
		Cellphone86: "13000000000",
	}

	err = validation.Validate(cp2)
	if err != nil {
		t.Error(err)
	}

	cp3 := &CPs3{
		Digit:       1,
		Cellphone86: "13000000000",
	}

	err = validation.Validate(cp3)
	if err == nil {
		t.Error("expect:" + validation.ErrUnsupportedType.Error())
	}
}

func TestMax(t *testing.T) {

	type CPInt struct {
		Digit int `valid:"max=1"`
	}

	cpInt := CPInt{Digit: 1}

	err := validation.Validate(cpInt)
	if err != nil {
		t.Error(err)
	}

	cpInt.Digit = 0
	err = validation.Validate(cpInt)
	if err != nil {
		t.Error(err)
	}

	cpInt.Digit = 2
	err = validation.Validate(cpInt)
	if err == nil {
		t.Error("expect:" + validation.ErrMaxNumber.Error())
	}
}

func TestMin(t *testing.T) {

	type CPInt struct {
		Digit int `valid:"min=1"`
	}

	cpInt := CPInt{Digit: 1}

	err := validation.Validate(cpInt)
	if err != nil {
		t.Error(err)
	}

	cpInt.Digit = 2
	err = validation.Validate(cpInt)
	if err != nil {
		t.Error(err)
	}

	cpInt.Digit = 0
	err = validation.Validate(cpInt)
	if err == nil {
		t.Error("expect:" + validation.ErrMinNumber.Error())
	}
}

func TestRange(t *testing.T) {

	type CPInt struct {
		Digit int `valid:"range=3,10"`
	}

	cpInt := CPInt{Digit: 1}

	err := validation.Validate(cpInt)
	if err != validation.ErrRangeLower {
		t.Error("expect:" + validation.ErrRangeLower.Error())
	}

	cpInt.Digit = 11
	err = validation.Validate(cpInt)
	if err != validation.ErrRangeUpper {
		t.Error("expect:" + validation.ErrRangeUpper.Error())
	}

	cpInt.Digit = 5
	err = validation.Validate(cpInt)
	if err != nil {
		t.Error(err)
	}

	type CPIntErr struct {
		Digit int `valid:"range=3,a"`
	}

	cpIntErr := CPIntErr{Digit: 5}

	err = validation.Validate(cpIntErr)
	if err == nil {
		t.Error(err)
	}
}

func TestEmail(t *testing.T) {

	type CPEmail struct {
		Email string `valid:"email"`
	}

	cpEmail := CPEmail{Email: "hhh@rutcode"}

	err := validation.Validate(cpEmail)
	if err != validation.ErrEmail {
		t.Error("expect:" + validation.ErrEmail.Error())
	}

	cpEmail.Email = "hhh@rutcode.com"
	err = validation.Validate(cpEmail)
	if err != nil {
		t.Error(err)
	}

}
