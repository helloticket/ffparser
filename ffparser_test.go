package ffparser

import (
	"testing"
	"time"

	"github.com/prometheus/common/log"
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
	NumInt       int       `record:"start=1,end=5,decorator=IntDecorator"`
	NumInt64     int64     `record:"start=6,end=10,decorator=Int64Decorator"`
	NumFloat64   float64   `record:"start=11,end=25,decorator=BrazilMoneyDecorator"`
	NumFloat64_2 float64   `record:"start=26,end=40,decorator=BrazilMoneyDecorator"`
	DataHora     time.Time `record:"start=41,end=54,decorator=DateTimeDecorator"`
	DataHoraBR   time.Time `record:"start=55,end=62,decorator=BrazilDateDecorator"`
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
		Value3: 45678,
		Value4: 4567833.22,
		Value5: "be happy",
	})

	assert.Equal(t, "10/05/2017 00:00:00  000000124020000004567804567833.22be happy   ", result)
}

func TestShouldCreateFromTextToStructWithAutoDetectFieldType(t *testing.T) {
	textoParaConversao := "be happy  "
	parser := NewSimpleParser()

	type AutoDetectStruct struct {
		Value string `record:"start=1,end=10"`
	}

	var objeto AutoDetectStruct
	err := parser.CreateFromText(&objeto, textoParaConversao)

	if err != nil {
		log.Errorf("Erro encontrado: %v", err)
	}

	assert.Equal(t, "be happy", objeto.Value)
}

func TestShouldParseTextToStructWithAutoDetectFieldTypeAndDecorators(t *testing.T) {
	parser := NewSimpleParser()

	result, _ := parser.ParseToText(&TestDecorator1{
		NumInt:       1,
		NumInt64:     int64(12402),
		NumFloat64:   178.5,
		NumFloat64_2: 167.0,
		DataHora:     time.Date(2020, 3, 31, 18, 55, 2, 0, time.UTC),
	})

	assert.Equal(t, "00001124020000000000178500000000000167003103202018550200000000", result)
}

func TestShouldCreateFromTextAutoDetectFieldTypeAndDecorators(t *testing.T) {
	texto := "00001124020000000000178500000000000167003103202018550200000000"
	var estrutura TestDecorator1
	parser := NewSimpleParser()
	parser.CreateFromText(&estrutura, texto)

	dataHoraEsperada := time.Date(2020, 3, 31, 18, 55, 2, 0, time.UTC)

	var dataBrEsperada time.Time

	assert.Equal(t, 1, estrutura.NumInt)
	assert.Equal(t, int64(12402), estrutura.NumInt64)
	assert.Equal(t, 178.5, estrutura.NumFloat64)
	assert.Equal(t, 167.0, estrutura.NumFloat64_2)
	assert.Equal(t, dataHoraEsperada.String(), estrutura.DataHora.String())
	assert.Equal(t, dataBrEsperada, estrutura.DataHoraBR)
}
