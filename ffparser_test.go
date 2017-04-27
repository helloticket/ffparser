package ffparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestPojo1 struct {
	Name        string `record:"start=1,end=10"`
	Address     string `record:"start=11,end=30"`
	PhoneNumber string `record:"start=31,end=40"`
	Other       string
}

type TestPojo2 struct {
	Name        string `record:"start=1,end=10,padchar=*"`
	Address     string `record:"start=11,end=30,padalign=left"`
	PhoneNumber string `record:"start=31,end=50,padchar=0,padalign=left"`
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

func TestShouldParseTextToStructWithBrazilDateDecorator(t *testing.T) {
	parser := NewSimpleParser()

	type DateStruct struct {
		ValueDate time.Time `record:"start=1,end=10,decorator=BrazilDateDecorator"`
	}

	result, err := parser.ParseToText(&DateStruct{ValueDate: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)})

	assert.Nil(t, err)
	assert.Equal(t, "01012017  ", result)
}

func TestShouldParseTextToStructWithBrazilSmallDateDecorator(t *testing.T) {
	parser := NewSimpleParser()

	type DateStruct struct {
		ValueDate time.Time `record:"start=1,end=8,decorator=BrazilSmallDateDecorator"`
	}

	result, err := parser.ParseToText(&DateStruct{ValueDate: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)})

	assert.Nil(t, err)
	assert.Equal(t, "010117  ", result)
}

func TestShouldParseTextToStructWithMoneyDecorator(t *testing.T) {
	parser := NewSimpleParser()

	type MoneyStruct struct {
		Value float64 `record:"start=1,end=12,decorator=BrazilMoneyDecorator"`
	}

	result1, _ := parser.ParseToText(&MoneyStruct{Value: 1.999})
	result2, _ := parser.ParseToText(&MoneyStruct{Value: 10.999})
	result3, _ := parser.ParseToText(&MoneyStruct{Value: 120.999})
	result4, _ := parser.ParseToText(&MoneyStruct{Value: 1230.999})
	result5, _ := parser.ParseToText(&MoneyStruct{Value: 5222230.999})

	assert.Equal(t, "000000001999", result1)
	assert.Equal(t, "000000010999", result2)
	assert.Equal(t, "000000120999", result3)
	assert.Equal(t, "000001230999", result4)
	assert.Equal(t, "005222230999", result5)
}

func TestShouldParseTextToStructWithAutoDetectFieldType(t *testing.T) {
	parser := NewSimpleParser()

	type AutoDetectStruct struct {
		Value1 time.Time `record:"start=1,end=21"`
		Value2 int       `record:"start=22,end=32"`
		Value3 int64     `record:"start=33,end=43"`
		Value4 float64   `record:"start=44,end=54"`
		Value5 string    `record:"start=55,end=65"`
	}

	result, _ := parser.ParseToText(&AutoDetectStruct{
		Value1: time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC),
		Value2: 12402,
		Value3: 4567822222,
		Value4: 4567833.22,
		Value5: "be happy",
	})

	assert.Equal(t, "10/05/2017 00:00:00  12402      4567822222 04567833.22be happy   ", result)
}
