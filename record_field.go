package ffparser

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
