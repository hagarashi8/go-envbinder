package envbinder

import (
	"fmt"
	"reflect"
	"strings"
)

func Bind(v interface{}) error {
	s := reflect.ValueOf(v).Elem()
	t := s.Type()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("Bind accepts only structs")
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}
		parseField(f, t.Field(i))
	}
	return nil
}

func parseField(v reflect.Value, sf reflect.StructField) error {
	t := reflect.StructTag(sf.Tag)
	if sf.Type.Kind() == reflect.Struct {
		return Bind(v.Addr().Interface())
	}

	tagraw, ok := t.Lookup("env")
	tags := strings.Split(tagraw, ",")
	name := tags[0]

	if !ok {
		return fmt.Errorf("")
	}
	switch sf.Type.Kind() {
	case reflect.Bool:
		b, err := lookupBool(name)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(b))
	case reflect.String:
		s, err := lookupString(name)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(s))
	case reflect.Int:
		i, err := lookupInt[int](name, 0)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	case reflect.Int8:
		i, err := lookupInt[int8](name, 8)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	case reflect.Int16:
		i, err := lookupInt[int16](name, 16)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	case reflect.Int32:
		i, err := lookupInt[int32](name, 32)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	case reflect.Int64:
		i, err := lookupInt[int64](name, 64)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
	default:
		return fmt.Errorf("Unsupported type")
	}
	return nil
}
