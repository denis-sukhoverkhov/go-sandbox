package sandbox

import (
	"fmt"
	"reflect"
)

func ChangeStructFieldValue(data interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(data)

	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("data must be a non-nil pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a pointer to a struct")
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}

	newVal := reflect.ValueOf(value)
	if field.Type() != newVal.Type() {
		return fmt.Errorf("field %s type mismatch", fieldName)
	}

	field.Set(newVal)

	return nil
}
