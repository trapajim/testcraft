package testcraft

import (
	"errors"
	"fmt"
	"reflect"
)

func randomize[T any](f T, valuer Valuer) (any, error) {
	t := reflect.TypeOf(f)
	v := reflect.New(t)
	fields := reflect.VisibleFields(t)
	var errs []error
	for _, field := range fields {
		sField := v.Elem().FieldByName(field.Name)
		if !sField.CanSet() {
			continue
		}
		if field.Type.Kind() == reflect.Slice {
			err := randomizeSlice(field, valuer, sField)
			errs = appendError(errs, err)
			continue
		}
		sField = createPointerField(sField)
		val := getValueForField(valuer, field.Type.Kind(), field)
		err := fieldAssign(sField, val)
		errs = appendError(errs, err)
	}
	return v.Interface(), errors.Join(errs...)
}

func randomizeSlice(field reflect.StructField, valuer Valuer, sField reflect.Value) error {
	slice, err := createSlice(field, valuer)
	if err != nil {
		return err
	}
	return fieldAssign(sField, slice)
}

func fieldAssign(sField reflect.Value, val reflect.Value) error {
	if !val.IsValid() {
		return fmt.Errorf("cannot set value to field %s", sField.Type().String())
	}
	vType := sField.Type()
	tValue := val.Type()
	if tValue.AssignableTo(vType) {
		sField.Set(val)
	} else {
		sField.Set(val.Convert(vType))
	}
	return nil
}

func createPointerField(sField reflect.Value) reflect.Value {
	if sField.Type().Kind() == reflect.Ptr {
		if sField.IsNil() {
			sField.Set(reflect.New(sField.Type().Elem()))
		}
		sField = sField.Elem()
	}
	return sField
}

func getValueForField(valuer Valuer, kind reflect.Kind, field reflect.StructField) reflect.Value {

	if v, ok := valuer[kind.String()]; ok {
		return v(field)
	}
	if v, ok := valuer[field.Type.String()]; ok {
		return v(field)
	}

	switch kind {
	case reflect.Struct:
		return handleStruct(field, valuer)
	case reflect.Ptr:
		return getValueForField(valuer, field.Type.Elem().Kind(), field)
	default:

		return reflect.Value{}
	}
}

func handleStruct(field reflect.StructField, valuer Valuer) reflect.Value {
	value := reflect.Zero(field.Type)
	ai := value.Interface()

	isPtr := field.Type.Kind() == reflect.Ptr
	if isPtr {
		ai = reflect.Indirect(reflect.New(field.Type.Elem())).Interface()
	}
	val, _ := randomize(ai, valuer)
	return reflect.ValueOf(val).Elem()
}

func appendError(errs []error, err error) []error {
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}

func createSlice(field reflect.StructField, valuer Valuer) (reflect.Value, error) {
	k := field.Type.Elem().Kind()
	if k == reflect.Ptr {
		k = field.Type.Elem().Elem().Kind()
	}
	if k == reflect.Struct {
		return handleStructSlice(field, valuer)
	}
	return handlePrimitiveSlice(field, valuer)
}

func handleStructSlice(field reflect.StructField, valuer Valuer) (reflect.Value, error) {
	value := reflect.Zero(field.Type.Elem())
	ai := value.Interface()
	isPtr := field.Type.Elem().Kind() == reflect.Ptr
	if isPtr {
		ai = reflect.Indirect(reflect.New(value.Type().Elem())).Interface()
	}
	slice := reflect.MakeSlice(reflect.SliceOf(field.Type.Elem()), 0, 3)
	var errs []error
	for i := 0; i < 3; i++ {
		val, err := randomize(ai, valuer)
		if err != nil {
			errs = appendError(errs, err)
			continue
		}
		if isPtr {
			slice = reflect.Append(slice, reflect.ValueOf(val))
		} else {
			slice = reflect.Append(slice, reflect.Indirect(reflect.ValueOf(val)))
		}
	}
	return slice, errors.Join(errs...)
}

func handlePrimitiveSlice(field reflect.StructField, valuer Valuer) (reflect.Value, error) {
	k := field.Type.Elem().Kind()
	orig := k

	if k == reflect.Ptr {
		k = field.Type.Elem().Elem().Kind()
	}

	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(reflect.New(field.Type.Elem()).Elem().Interface())), 0, 3)

	for i := 0; i < 3; i++ {
		val := getValueForField(valuer, k, field)
		if !val.IsValid() {
			return reflect.Value{}, fmt.Errorf("cannot set value to slice %s", field.Name)
		}
		if orig == reflect.Ptr {
			ptr := reflect.New(reflect.TypeOf(val.Interface()))
			ptr.Elem().Set(val)
			val = ptr
		}
		slice = reflect.Append(slice, val)
	}
	return slice, nil
}
