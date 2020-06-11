package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrazilMoneyDecoratorValor(t *testing.T) {

	testeNumero := float64(322.15)

	brMoneyDecorator := BrazilMoneyDecorator{}
	strNum, err := brMoneyDecorator.ToString(testeNumero)

	assert.Nil(t, err)
	assert.Equal(t, "32215", strNum)
}

func TestBrazilMoneyDecoratorString(t *testing.T) {

	testeNumeroStr := "0000032215"

	brMoneyDecorator := BrazilMoneyDecorator{}
	numeroFloat64, err := brMoneyDecorator.FromString(testeNumeroStr)
	assert.Nil(t, err)
	assert.Equal(t, float64(322.15), numeroFloat64)
}
