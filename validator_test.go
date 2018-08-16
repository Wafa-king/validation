// MIT License

// Copyright (c) 2016 rutcode-go

package validation_test

import (
	"fmt"
	"testing"

	"github.com/go-trellis/validation"
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

	cp1.Digit = "我"
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

	type CPStr struct {
		Str string `valid:"min=1"`
	}

	type CPSli struct {
		Sli []int `valid:"min=1"`
	}

	cpInt := CPInt{Digit: 1}
	cpStr := CPStr{Str: "a"}
	cpSli := CPSli{Sli: []int{1, 2}}

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

	err = validation.Validate(cpStr)
	if err != nil{
		t.Error(err)
	}

	cpStr.Str = "abc"
	err = validation.Validate(cpStr)
	if err != nil{
		t.Error(err)
	}

	cpStr.Str = ""
	err = validation.Validate(cpStr)
	if err == nil{
		t.Error(err)
	}

	err = validation.Validate(cpSli)
	if err != nil{
		t.Error(err)
	}

	cpSli.Sli = []int{}
	err = validation.Validate(cpSli)
	if err == nil{
		t.Error(err)
	}

}

func TestFloatMin(t *testing.T) {

	type CPTest struct {
		Digit1 float32 `valid:"fmin=1"`
		Digit2 int     `valid:"fmin=1"`
	}

	cpTest := CPTest{Digit1: 1.0, Digit2: 1}

	err := validation.Validate(cpTest)
	if err != nil {
		t.Error(err)
	}

	cpTest.Digit1 = 2
	err = validation.Validate(cpTest)
	if err != nil {
		t.Error(err)
	}

	cpTest.Digit1 = 0.9
	err = validation.Validate(cpTest)
	if err == nil {
		t.Error("expect:" + validation.ErrMinNumber.Error())
	}
}

func TestFloatMax(t *testing.T) {

	type CPTest struct {
		Digit1 float32 `valid:"fmax=1"`
		Digit2 int     `valid:"fmax=1"`
	}

	cpTest := CPTest{Digit1: 1, Digit2: 1}

	err := validation.Validate(cpTest)
	if err != nil {
		t.Error(err)
	}

	cpTest.Digit1 = 0.8
	err = validation.Validate(cpTest)
	if err != nil {
		t.Error(err)
	}

	cpTest.Digit1 = 2.0
	err = validation.Validate(cpTest)
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

func TestRequired(t *testing.T) {

	type CPRequired struct {
		Email string `valid:"required"`
		Bool  bool   `valid:"required"`
	}
	cpRequired := CPRequired{Bool: true}

	err := validation.Validate(cpRequired)
	if err != validation.ErrNil {
		t.Error("expect:" + validation.ErrNil.Error())
	}

	cpRequired.Email = "hhh@rutcode.com"
	cpRequired.Bool = false
	err = validation.Validate(cpRequired)
	if err != validation.ErrNil {
		t.Error("expect:" + validation.ErrNil.Error())
	}

	cpRequired.Bool = true
	err = validation.Validate(cpRequired)
	if err != nil {
		t.Error(err)
	}
}

type TestNewValidation struct{}

func (*TestNewValidation) IsValid(v interface{}, _ string) error {
	if v.(int64) < 1000 {
		return fmt.Errorf("test below 1000")
	}
	return nil
}

func NewTestNewValidation() *TestNewValidation {
	return new(TestNewValidation)
}

func TestRegistCustomerValidator(t *testing.T) {

	type CPValidation1 struct {
		Digit int64 `valid:"test1"`
	}

	validation.RegistCustomerValidator("test1", NewTestNewValidation()).
		RegistCustomerValidator("test2", NewTestNewValidation())

	err := validation.Validate(CPValidation1{Digit: 1})
	if err == nil {
		t.Error("expect: test below 1000")
	}

	type CPValidation2 struct {
		Digit int64 `valid:"test2"`
	}

	err = validation.Validate(CPValidation2{Digit: 1001})
	if err != nil {
		t.Error(err)
	}

}
