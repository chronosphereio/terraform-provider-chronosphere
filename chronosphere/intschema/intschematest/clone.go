package intschematest

import (
	"fmt"
	"reflect"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/tfid"
)

var typeTFID = reflect.TypeOf(tfid.ID{})

// Clone creates a copy of an intschema struct.
// It only supports a limited set of types used by intschema structs.
func Clone[T any](v T) T {
	rv := reflect.ValueOf(v)
	clone := reflect.New(rv.Type()).Elem()

	cloneValue(rv, clone)
	return clone.Interface().(T)
}

func cloneValue(src, dst reflect.Value) {
	if src.Type() == typeTFID {
		// We can't recurse into tfid.ID as it contains unexported fields that
		// reflect can't set. However, it has no reference types, so we can set
		// the value directly without recursing in.
		dst.Set(src)
		return
	}

	switch src.Type().Kind() {
	case reflect.Bool, reflect.String, reflect.Float64, reflect.Int64:
		dst.Set(src)
	case reflect.Pointer:
		clonePtr(src, dst)
	case reflect.Slice:
		cloneSlice(src, dst)
	case reflect.Map:
		cloneMap(src, dst)
	case reflect.Struct:
		cloneStruct(src, dst)
	default:
		panic(fmt.Errorf("unsupported type %v", src.Type()))
	}
}

func clonePtr(src, dst reflect.Value) {
	if src.IsNil() {
		return
	}

	dst.Set(reflect.New(src.Elem().Type()))
	cloneValue(src.Elem(), dst.Elem())
}

func cloneSlice(src, dst reflect.Value) {
	if src.IsNil() {
		return
	}

	dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
	for i := 0; i < src.Len(); i++ {
		cloneValue(src.Index(i), dst.Index(i))
	}
}

func cloneMap(src, dst reflect.Value) {
	if src.IsNil() {
		return
	}

	dst.Set(reflect.MakeMap(src.Type()))
	for _, k := range src.MapKeys() {
		dst.SetMapIndex(k, src.MapIndex(k))
	}
}

func cloneStruct(src, dst reflect.Value) {
	for i := 0; i < src.Type().NumField(); i++ {
		cloneValue(src.Field(i), dst.Field(i))
	}
}
