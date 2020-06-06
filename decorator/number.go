package decorator

import (
	"fmt"

	"github.com/helloticket/ffparser/helper"
)

type IntDecorator struct {
}

type Int64Decorator struct {
}

type Float64Decorator struct {
}

func (i *IntDecorator) ToString(field interface{}) (string, error) {
	return fmt.Sprintf("%d", field), nil
}

func (i *IntDecorator) FromString(field string) (interface{}, error) {
	return helper.ToInteger(field), nil
}

func (i *Int64Decorator) ToString(field interface{}) (string, error) {
	return fmt.Sprintf("%d", field), nil
}

func (i *Int64Decorator) FromString(field string) (interface{}, error) {
	return helper.ToInteger64(field), nil
}

func (i *Float64Decorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(float64); ok {
		return helper.Float64ToString(value), nil
	}

	return "", nil
}

func (i *Float64Decorator) FromString(field string) (interface{}, error) {
	return helper.ToFloat64(field), nil
}
