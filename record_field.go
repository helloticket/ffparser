package ffparser

import (
    "fmt"

    "github.com/helderfarias/ffparser/helper"
)

type RecordField struct {
    FieldName    string
    Decorator    string
    Delimiter    string
    PaddingAlign string
    Start        int
    End          int
    Size         int
}

type RecordFieldSorted []RecordField

func (a RecordFieldSorted) Len() int           { return len(a) }
func (a RecordFieldSorted) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RecordFieldSorted) Less(i, j int) bool { return a[i].Start < a[j].Start }

func (r *RecordField) ApplyFormat(value string) (string, error) {
    if r.PaddingAlign == "right" {
        return helper.RightPadToLen(value, r.Delimiter, r.Size), nil
    }

    if r.PaddingAlign == "left" {
        return helper.LeftPadToLen(value, r.Delimiter, r.Size), nil
    }

    return "", fmt.Errorf("Padding align invalid")
}
