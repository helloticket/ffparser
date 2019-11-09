package ffparser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	unexported uint64
	Dummy      string `test:"dummytag"`
	Yummy      int    `test:"yummytag"`
}

func TestGetField_on_struct(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	value, err := getField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, value, "test")
}

func TestGetField_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
	}

	value, err := getField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, value, "test")
}

func TestGetField_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := getField(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetField_non_existing_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	_, err := getField(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetField_unexported_field(t *testing.T) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	assert.Panics(t, func() {
		getField(dummyStruct, "unexported")
	})
}

func TestGetFieldKind_on_struct(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := getFieldKind(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.String)

	kind, err = getFieldKind(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.Int)
}

func TestGetFieldKind_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	kind, err := getFieldKind(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.String)

	kind, err = getFieldKind(dummyStruct, "Yummy")
	assert.NoError(t, err)
	assert.Equal(t, kind, reflect.Int)
}

func TestGetFieldKind_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := getFieldKind(dummy, "Dummy")
	assert.Error(t, err)
}

func TestGetFieldKind_non_existing_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	_, err := getFieldKind(dummyStruct, "obladioblada")
	assert.Error(t, err)
}

func TestGetFieldTag_on_struct(t *testing.T) {
	dummyStruct := TestStruct{}

	tag, err := getFieldTag(dummyStruct, "Dummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "dummytag")

	tag, err = getFieldTag(dummyStruct, "Yummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "yummytag")
}

func TestGetFieldTag_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{}

	tag, err := getFieldTag(dummyStruct, "Dummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "dummytag")

	tag, err = getFieldTag(dummyStruct, "Yummy", "test")
	assert.NoError(t, err)
	assert.Equal(t, tag, "yummytag")
}

func TestGetFieldTag_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := getFieldTag(dummy, "Dummy", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_non_existing_field(t *testing.T) {
	dummyStruct := TestStruct{}

	_, err := getFieldTag(dummyStruct, "obladioblada", "test")
	assert.Error(t, err)
}

func TestGetFieldTag_unexported_field(t *testing.T) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}

	_, err := getFieldTag(dummyStruct, "unexported", "test")
	assert.Error(t, err)
}

func TestSetField_on_struct_with_valid_value_type(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := setField(&dummyStruct, "Dummy", "abc")
	assert.NoError(t, err)
	assert.Equal(t, dummyStruct.Dummy, "abc")
}

func TestSetField_non_existing_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := setField(&dummyStruct, "obladioblada", "life goes on")
	assert.Error(t, err)
}

func TestSetField_invalid_value_type(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	err := setField(&dummyStruct, "Yummy", "123")
	assert.Error(t, err)
}

func TestSetField_non_exported_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	assert.Error(t, setField(&dummyStruct, "unexported", "fail, bitch"))
}

func TestFields_on_struct(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestFields_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	fields, err := fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestFields_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := fields(dummy)
	assert.Error(t, err)
}

func TestFields_with_non_exported_fields(t *testing.T) {
	dummyStruct := TestStruct{
		unexported: 6789,
		Dummy:      "test",
		Yummy:      123,
	}

	fields, err := fields(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, fields, []string{"Dummy", "Yummy"})
}

func TestHasField_on_struct_pointer_with_existing_field(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := hasField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.True(t, has)
}

func TestHasField_non_existing_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := hasField(dummyStruct, "Test")
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestHasField_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := hasField(dummy, "Test")
	assert.Error(t, err)
}

func TestHasField_unexported_field(t *testing.T) {
	dummyStruct := TestStruct{
		unexported: 7890,
		Dummy:      "test",
		Yummy:      123,
	}

	has, err := hasField(dummyStruct, "unexported")
	assert.NoError(t, err)
	assert.False(t, has)
}

func TestTags_on_struct(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := tags(dummyStruct, "test")
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func TestTags_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := tags(dummyStruct, "test")
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func TestTags_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := tags(dummy, "test")
	assert.Error(t, err)
}

func TestItems_on_struct(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := items(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func TestItems_on_struct_pointer(t *testing.T) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	tags, err := items(dummyStruct)
	assert.NoError(t, err)
	assert.Equal(t, tags, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func TestItems_on_non_struct(t *testing.T) {
	dummy := "abc 123"

	_, err := items(dummy)
	assert.Error(t, err)
}

func TestSetField_int64_field(t *testing.T) {
	type A struct {
		ID int64
	}

	dummyStruct := A{ID: 10}

	err := setField(&dummyStruct, "ID", int64(20))

	assert.Nil(t, err)
	assert.Equal(t, int64(20), dummyStruct.ID)
}

func TestSetField_pointer_field(t *testing.T) {
	type A struct {
		ID int64
	}

	var ptr *A

	ptr = &A{ID: 10}

	err := setField(&ptr, "ID", int64(20))

	assert.Nil(t, err)
	assert.Equal(t, int64(20), ptr.ID)
}

func TestHasField_on_struct_with_existing_field(t *testing.T) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}

	has, err := hasField(dummyStruct, "Dummy")
	assert.NoError(t, err)
	assert.True(t, has)
}
