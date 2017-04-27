package decorator

import (
	"fmt"

	"github.com/helderfarias/ffparser/helper"
)

type IntDecorator struct {
}

type Int64Decorator struct {
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
