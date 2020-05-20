package faker

import (
	"io"
	"reflect"
)

func UniqueSlice(a interface{}) interface{} {
	aReflect := reflect.ValueOf(a)
	aType := reflect.TypeOf(a)
	aMap := map[interface{}]reflect.Value{}

	if aReflect.Kind() != reflect.Slice {
		return nil
	}

	cReflect := reflect.MakeSlice(aType, 0, 0)

	for i := 0; i < aReflect.Len(); i++ {
		if !aMap[aReflect.Index(i).Interface()].IsValid() {
			aMap[aReflect.Index(i).Interface()] = aReflect.Index(i)
			cReflect = reflect.Append(cReflect, aReflect.Index(i))
		}
	}

	return cReflect.Interface()
}

func UniqueSliceStr(a []string) []string {
	return UniqueSlice(a).([]string)
}

func MakeReflectNew(ref reflect.Value) interface{} {
	switch ref.Kind() {
	case reflect.String:
		return ""
	case reflect.Bool:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return 0
	case reflect.Float32, reflect.Float64:
		return 0.0
	case reflect.Map:
		return reflect.MakeMap(ref.Type())
	case reflect.Slice:
		return reflect.MakeSlice(ref.Type(), 0, 0)
	case reflect.Chan:
		return reflect.MakeChan(ref.Type(), 0)
	case reflect.Func:
		return reflect.MakeFunc(ref.Type(), func([]reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf(io.EOF)}
		})
	default:
		//Uintptr
		//Complex64
		//Complex128
		//Array
		//Interface
		//Ptr
		//Struct
		//UnsafePointer
		return nil
	}
}
