package ffparser

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"time"

	"github.com/helloticket/ffparser/decorator"
)

//FFParser flat file parser
type FFParser struct {
	decorators map[string]decorator.FieldDecorator
}

func NewSimpleParser() *FFParser {
	instance := FFParser{decorators: map[string]decorator.FieldDecorator{}}
	instance.decorators["Default"] = &decorator.DefaultDecorator{}
	instance.decorators["IntDecorator"] = &decorator.IntDecorator{}
	instance.decorators["Int64Decorator"] = &decorator.Int64Decorator{}
	instance.decorators["Float64Decorator"] = &decorator.Float64Decorator{}
	instance.decorators["BrazilDateDecorator"] = &decorator.BrazilDateDecorator{}
	instance.decorators["BrazilSmallDateDecorator"] = &decorator.BrazilSmallDateDecorator{}
	instance.decorators["BrazilMoneyDecorator"] = &decorator.BrazilMoneyDecorator{}
	instance.decorators["DateTimeDecorator"] = &decorator.DateTimeDecorator{}
	return &instance
}

func (f *FFParser) ParseToText(src interface{}) (string, error) {
	mapFields, _ := tags(src, "record")
	if len(mapFields) <= 0 {
		return "", fmt.Errorf("Could not fields public")
	}

	var buffer bytes.Buffer
	recordsField, _ := f.handlerRecordFieldsAndSort(src, mapFields)

	for _, record := range recordsField {
		content, err := getField(src, record.FieldName)
		if err != nil {
			return "", err
		}

		decorator, err := f.factoryDecorator(content, record)
		if err != nil {
			return "", err
		}

		value, err := decorator.ToString(content)
		if err != nil {
			return "", err
		}

		result, err := record.ApplyFormat(value)
		if err != nil {
			return "", err
		}

		buffer.WriteString(result)
	}

	return buffer.String(), nil
}

func (f *FFParser) CreateFromText(src interface{}, text string) error {
	mapFields, _ := tags(src, "record")
	if len(mapFields) <= 0 {
		return fmt.Errorf("Could not fields public")
	}

	recordsField, _ := f.handlerRecordFieldsAndSort(src, mapFields)

	for _, record := range recordsField {
		decorator, err := f.factoryDecorator(src, record)
		if err != nil {
			return err
		}

		value, err := decorator.FromString(strings.TrimSpace(text[record.Start:record.End]))
		if err != nil {
			return err
		}

		if err := setField(src, record.FieldName, value); err != nil {
			return err
		}
	}

	return nil
}

func (f *FFParser) mapperTags(tagName string) map[string]string {
	tags := map[string]string{}

	for _, tagValue := range strings.Split(tagName, ",") {
		entry := strings.Split(tagValue, "=")

		if len(entry) >= 2 {
			tags[entry[0]] = entry[1]
		} else {
			tags[entry[0]] = ""
		}
	}

	return tags
}

func (f *FFParser) factoryDecorator(obj interface{}, record RecordField) (decorator.FieldDecorator, error) {
	if record.Decorator != "" {
		if decorator := f.decorators[record.Decorator]; decorator != nil {
			return decorator, nil
		}
	}

	if record.FieldType == reflect.Int ||
		record.FieldType == reflect.Int8 ||
		record.FieldType == reflect.Int16 ||
		record.FieldType == reflect.Int32 {
		return f.decorators["IntDecorator"], nil
	}

	if record.FieldType == reflect.Int64 {
		return f.decorators["Int64Decorator"], nil
	}

	if record.FieldType == reflect.Float32 ||
		record.FieldType == reflect.Float64 {
		return f.decorators["Float64Decorator"], nil
	}

	if record.FieldType == reflect.Struct {
		if _, ok := obj.(time.Time); ok {
			return f.decorators["DateTimeDecorator"], nil
		}
	}

	return f.decorators["Default"], nil
}

func (f *FFParser) handlerRecordFieldsAndSort(src interface{}, fields map[string]string) ([]RecordField, error) {
	records := []RecordField{}

	for fieldName, tagName := range fields {
		if tagName == "" {
			continue
		}

		tags := f.mapperTags(tagName)
		start := toInteger(tags["start"]) - 1
		end := toInteger(tags["end"])
		size := (end - start)
		decorator := tags["decorator"]
		padChar := ""
		padAlign := ""
		fieldType, _ := getFieldKind(src, fieldName)

		if tags["padchar"] != "" {
			padChar = tags["padchar"]
		}

		if tags["padalign"] != "" {
			padAlign = tags["padalign"]
		}

		records = append(records, RecordField{
			FieldName:    fieldName,
			Start:        start,
			End:          end,
			Size:         size,
			Decorator:    decorator,
			PaddingChar:  padChar,
			PaddingAlign: padAlign,
			FieldType:    fieldType,
		})
	}

	sort.Sort(RecordFieldSorted(records))

	return records, nil
}
