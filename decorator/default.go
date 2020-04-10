package decorator

import (
	"fmt"
	"reflect"
	"strings"
)

type DefaultDecorator struct {
}

func (i *DefaultDecorator) ToString(field interface{}) (string, error) {
	return fmt.Sprintf("%v", field), nil
}

func (i *DefaultDecorator) FromString(field string) (interface{}, error) {
	if reflect.TypeOf(field).String() == "string" {
		return strings.TrimSpace(field), nil
	}
	return field, nil
}
