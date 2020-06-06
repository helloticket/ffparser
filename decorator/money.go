package decorator

import (
	"fmt"
	"strings"

	"github.com/helloticket/ffparser/helper"
)

type BrazilMoneyDecorator struct {
}

func (i *BrazilMoneyDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(float64); ok {
		src := helper.Float64ToString(float64(int(value * 100)))
		return strings.Replace(src, ".", "", -1), nil
	}

	return "", nil
}

func (i *BrazilMoneyDecorator) FromString(field string) (interface{}, error) {
	if strings.TrimSpace(field) == "" {
		return float64(0), nil
	}
	decimalPart := field[len(field)-2:]
	integerPart := field[:len(field)-2]
	value := fmt.Sprintf("%s.%s", integerPart, decimalPart)
	return helper.ToFloat64(value), nil
}
