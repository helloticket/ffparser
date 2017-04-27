package ffparser

import (
	"reflect"

	"github.com/helderfarias/ffparser/helper"
)

type RecordField struct {
	FieldName    string
	Decorator    string
	PaddingChar  string
	PaddingAlign string
	Start        int
	End          int
	Size         int
	FieldType    reflect.Kind
}

type RecordFieldSorted []RecordField

func (a RecordFieldSorted) Len() int           { return len(a) }
func (a RecordFieldSorted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RecordFieldSorted) Less(i, j int) bool { return a[i].Start < a[j].Start }

func (r *RecordField) ApplyFormat(value string) (string, error) {
	if r.PaddingAlign == "" && r.PaddingChar == "" && r.isNumericField() {
		return helper.LeftPadToLen(value, "0", r.Size), nil
	}

	paddingCharacter := " "
	if r.PaddingChar != "" {
		paddingCharacter = r.PaddingChar
	}

	if r.PaddingAlign == "left" {
		return helper.LeftPadToLen(value, paddingCharacter, r.Size), nil
	}

	return helper.RightPadToLen(value, paddingCharacter, r.Size), nil
}

func (r *RecordField) isNumericField() bool {
	switch r.FieldType {
	case reflect.Int:
	case reflect.Int64:
	case reflect.Float64:
		return true
	}

	return false
}
