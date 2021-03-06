# validation


valid go struct values

set "valid:-" flag in struct

## Build

* [![Build Status](https://travis-ci.org/go-rut/validation.png)](https://travis-ci.org/go-rut/validation)

## Usage

### 1. customer validator

**follow the interface**

```go
type Validator interface {
    IsValid(interface{}, string) error
}
```

**customer struct and functions**

```go
type TestNewValidation struct{}

func (*TestNewValidation) IsValid(v interface{}, _ string) error {
    if v.(int64) < 1000 {
        return fmt.Errorf("test below 1000")
    }
    return nil
}

func NewTestNewValidation() *TestNewValidation {
    return (TestNewValidation)(nil)
}
```

**regist customer validator**

```go
func TestRegistCustomerValidator(t *testing.T) {

    type CPValidation1 struct {
        Digit int64 `valid:"test1"`
    }

    // Set customer validator with name
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
```

### 2. validate cellphone86

```go
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
```

### 3. validate length

```go
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
```

### 4. validate maximum number

```go
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
```

### 5. validate minimum number

```go
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
```

### 6. validate range of number

```go
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
```

### 7. validate email

```go
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
```

### 8. validate required values

```go
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
```
