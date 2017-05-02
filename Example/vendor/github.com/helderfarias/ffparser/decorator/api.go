package decorator

type FieldDecorator interface {
	ToString(field interface{}) (string, error)
	FromString(field string) (interface{}, error)
}
