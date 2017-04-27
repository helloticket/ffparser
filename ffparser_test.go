package ffparser

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

type TestPojo1 struct {
    Name        string `record:"start=1,end=10"`
    Address     string `record:"start=11,end=30"`
    PhoneNumber string `record:"start=31,end=40"`
    Other       string
}

type TestPojo2 struct {
    Name        string `record:"start=1,end=10,delimiter=*"`
    Address     string `record:"start=11,end=30,padalign=left"`
    PhoneNumber string `record:"start=31,end=50,delimiter=0,padalign=left"`
    Other       string
}

type TestDecorator1 struct {
    NumInt   int   `record:"start=1,end=5,decorator=IntDecorator"`
    NumInt64 int64 `record:"start=6,end=10,decorator=Int64Decorator"`
}

func TestShouldParseTextToStruct(t *testing.T) {
    parser := NewSimpleParser()

    var expectedResult TestPojo1

    err := parser.CreateFromText(&expectedResult, `name111111222222222222222222223333333333`)

    assert.Nil(t, err)
    assert.Equal(t, TestPojo1{Name: "name111111", Address: "22222222222222222222", PhoneNumber: "3333333333"}, expectedResult)
}

func TestShouldParseTextToStructWithIntegerDecorator(t *testing.T) {
    parser := NewSimpleParser()

    var expectedResult TestDecorator1

    err := parser.CreateFromText(&expectedResult, "1111122222")

    assert.Nil(t, err)
    assert.Equal(t, TestDecorator1{NumInt: 11111, NumInt64: 22222}, expectedResult)
}

func TestShouldParseStructToText(t *testing.T) {
    parser := NewSimpleParser()

    result, err := parser.ParseToText(&TestPojo1{Name: "name", Address: "Address", PhoneNumber: "PhoneNumber"})

    assert.Nil(t, err)
    assert.Equal(t, "name      Address             PhoneNumbe", result)
}

func TestShouldParseStructToTextWithDelimiter(t *testing.T) {
    parser := NewSimpleParser()

    result, err := parser.ParseToText(&TestPojo2{Name: "name", Address: "Address", PhoneNumber: "8032342345"})

    assert.Nil(t, err)
    assert.Equal(t, "name******             Address00000000008032342345", result)
}
