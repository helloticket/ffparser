package decorator

import "fmt"

type DefaultDecorator struct {
}

func (i *DefaultDecorator) ToString(field interface{}) (string, error) {
	return fmt.Sprintf("%v", field), nil
}

func (i *DefaultDecorator) FromString(field string) (interface{}, error) {
	return field, nil
}
