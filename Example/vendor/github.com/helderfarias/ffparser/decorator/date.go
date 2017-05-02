package decorator

import (
	"time"
)

const (
	DDMMYYYY       = "02012006"
	DDMMYY         = "020106"
	DDMMYYYYHHMMSS = "02/01/2006 15:04:05"
)

type DateTimeDecorator struct {
}

type BrazilDateDecorator struct {
}

type BrazilSmallDateDecorator struct {
}

func (i *BrazilDateDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		return value.Format(DDMMYYYY), nil
	}
	return "", nil
}

func (i *BrazilSmallDateDecorator) FromString(field string) (interface{}, error) {
	value, err := time.Parse(DDMMYYYY, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (i *BrazilSmallDateDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		return value.Format(DDMMYY), nil
	}
	return "", nil
}

func (i *BrazilDateDecorator) FromString(field string) (interface{}, error) {
	value, err := time.Parse(DDMMYY, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (i *DateTimeDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		return value.Format(DDMMYYYYHHMMSS), nil
	}
	return "", nil
}

func (i *DateTimeDecorator) FromString(field string) (interface{}, error) {
	value, err := time.Parse(DDMMYYYYHHMMSS, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}
