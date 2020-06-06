package decorator

import (
	"fmt"
	"strconv"
	"strings"
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

func (i *BrazilDateDecorator) FromString(field string) (interface{}, error) {
	if strings.TrimSpace(field) == "" || field == "00000000" {
		var value time.Time
		return value, nil
	}

	value, err := time.Parse(DDMMYYYY, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (i *BrazilDateDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		strValue := value.Format(DDMMYYYY)
		if strValue == "01010001" { // Quando a data não é informada o time.Time inicializa com 01/01/0001
			strValue = "00000000"
		}
		return strValue, nil
	}
	return "", nil
}

func (i *BrazilSmallDateDecorator) FromString(field string) (interface{}, error) {
	if strings.TrimSpace(field) == "" || field == "00000000" {
		var value time.Time
		return value, nil
	}

	value, err := time.Parse(DDMMYYYY, field)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (i *BrazilSmallDateDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		strValue := value.Format(DDMMYY)
		if strValue == "010101" { // Quando a data não é informada o time.Time inicializa com 01/01/0001
			strValue = "000000"
		}
		return strValue, nil
	}
	return "", nil
}

// func (i *DateTimeDecorator) ToString(field interface{}) (string, error) {
// 	if value, ok := field.(time.Time); ok {
// 		return value.Format(DDMMYYYYHHMMSS), nil
// 	}
// 	return "", nil
// }

// func (i *DateTimeDecorator) FromString(field string) (interface{}, error) {
// 	value, err := time.Parse(DDMMYYYYHHMMSS, field)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return value, nil
// }

func (i *DateTimeDecorator) FromString(field string) (interface{}, error) {
	if strings.TrimSpace(field) == "" || field == "00000000000000" {
		var value time.Time
		return value, nil
	}
	year, _ := strconv.Atoi(field[4:8])
	month, _ := strconv.Atoi(field[2:4])
	day, _ := strconv.Atoi(field[:2])
	hour, _ := strconv.Atoi(field[8:10])
	minutes, _ := strconv.Atoi(field[10:12])
	seconds, _ := strconv.Atoi(field[12:])
	value := time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, time.UTC)
	return value, nil
}

func (i *DateTimeDecorator) ToString(field interface{}) (string, error) {
	if value, ok := field.(time.Time); ok {
		strTime := fmt.Sprintf("%02d%02d%04d%02d%02d%02d", value.Day(), value.Month(), value.Year(),
			value.Hour(), value.Minute(), value.Second())
		return strTime, nil
	}
	return "", nil
}
