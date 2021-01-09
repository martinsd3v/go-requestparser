package parser

import (
	"reflect"
	"strconv"
	"time"
)

var emptyField = reflect.StructField{}

//trySetValue responsible for trying to assign value to a string
func trySetValue(item reflect.Value, insert FormSlice, field string) {
	valueInsert, ok := insert[field]
	if ok {
		switch item.Kind() {

		case reflect.Slice:
			setSlice(valueInsert, item, emptyField)
			return
		default:
			setWithProperType(valueInsert[0], item, emptyField)
		}
	}
}

func setWithProperType(val string, value reflect.Value, field reflect.StructField) {
	switch value.Kind() {
	case reflect.Int:
		setIntField(val, 0, value)
		return
	case reflect.Int8:
		setIntField(val, 8, value)
		return
	case reflect.Int16:
		setIntField(val, 16, value)
		return
	case reflect.Int32:
		setIntField(val, 32, value)
		return
	case reflect.Int64:
		switch value.Interface().(type) {
		case time.Duration:
			setTimeDuration(val, value, field)
			return
		}
		setIntField(val, 64, value)
		return
	case reflect.Uint:
		setUintField(val, 0, value)
		return
	case reflect.Uint8:
		setUintField(val, 8, value)
		return
	case reflect.Uint16:
		setUintField(val, 16, value)
		return
	case reflect.Uint32:
		setUintField(val, 32, value)
		return
	case reflect.Uint64:
		setUintField(val, 64, value)
		return
	case reflect.Bool:
		setBoolField(val, value)
		return
	case reflect.Float32:
		setFloatField(val, 32, value)
		return
	case reflect.Float64:
		setFloatField(val, 64, value)
		return
	case reflect.String:
		value.SetString(val)
	}
}

func setIntField(val string, bitSize int, field reflect.Value) {
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
}

func setUintField(val string, bitSize int, field reflect.Value) {
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
}

func setBoolField(val string, field reflect.Value) {
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
}

func setFloatField(val string, bitSize int, field reflect.Value) {
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
}

func setArray(vals []string, value reflect.Value, field reflect.StructField) {
	for i, s := range vals {
		setWithProperType(s, value.Index(i), field)
	}
}

func setSlice(vals []string, value reflect.Value, field reflect.StructField) {
	slice := reflect.MakeSlice(value.Type(), len(vals), len(vals))
	setArray(vals, slice, field)
	value.Set(slice)
}

func setTimeDuration(val string, value reflect.Value, field reflect.StructField) {
	d, err := time.ParseDuration(val)
	if err == nil {
		value.Set(reflect.ValueOf(d))
	}
}
