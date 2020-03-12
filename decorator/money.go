package decorator

import (
	"strings"

	"github.com/helderfarias/ffparser/helper"
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
	value := strings.Replace(field, ".", "", -1)
	value = strings.Replace(value, ",", ".", -1)
	return helper.ToFloat64(value), nil
}
