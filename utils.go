package faker

import "reflect"

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
