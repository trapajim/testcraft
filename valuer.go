package testcraft

import (
	"github.com/trapajim/testcraft/datagen"
	"reflect"
)

type Valuer map[string]ValuerFunc
type ValuerFunc func(reflect.StructField) reflect.Value

func valueString(field reflect.StructField) reflect.Value {
	return reflect.ValueOf(datagen.Words(2))
}

func valueInt(field reflect.StructField) reflect.Value {
	return reflect.ValueOf(datagen.Rand().Int(100)).Convert(field.Type)
}

func valueFloat(field reflect.StructField) reflect.Value {
	return reflect.ValueOf(datagen.Rand().Float64(100)).Convert(field.Type)
}

func valueBool(field reflect.StructField) reflect.Value {
	return reflect.ValueOf(datagen.Bool())
}

func valueTime(field reflect.StructField) reflect.Value {
	return reflect.ValueOf(datagen.Time())
}

func defaultValuer() Valuer {
	return map[string]ValuerFunc{
		reflect.String.String():  valueString,
		reflect.Int.String():     valueInt,
		reflect.Int8.String():    valueInt,
		reflect.Int16.String():   valueInt,
		reflect.Int32.String():   valueInt,
		reflect.Int64.String():   valueInt,
		reflect.Uint.String():    valueInt,
		reflect.Uint8.String():   valueInt,
		reflect.Uint16.String():  valueInt,
		reflect.Uint32.String():  valueInt,
		reflect.Uint64.String():  valueInt,
		reflect.Float32.String(): valueFloat,
		reflect.Float64.String(): valueFloat,
		reflect.Bool.String():    valueBool,
		"time.Time":              valueTime,
		"*time.Time":             valueTime,
	}
}
