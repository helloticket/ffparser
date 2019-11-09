package ffparser

import (
	"reflect"
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
		return leftPadToLen(value, "0", r.Size), nil
	}

	paddingCharacter := " "
	if r.PaddingChar != "" {
		paddingCharacter = r.PaddingChar
	}

	if r.PaddingAlign == "left" {
		return leftPadToLen(value, paddingCharacter, r.Size), nil
	}

	return rightPadToLen(value, paddingCharacter, r.Size), nil
}

func (r *RecordField) isNumericField() bool {
	if reflect.Int == r.FieldType ||
		reflect.Int16 == r.FieldType ||
		reflect.Int32 == r.FieldType ||
		reflect.Int8 == r.FieldType ||
		reflect.Int64 == r.FieldType ||
		reflect.Float32 == r.FieldType ||
		reflect.Float64 == r.FieldType {
		return true
	}

	return false
}
