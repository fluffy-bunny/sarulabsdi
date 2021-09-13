package di

import (
	"reflect"

	"github.com/fatih/structtag"
)

func MakeDefaultBuildByType(rtElem reflect.Type) func(ctn Container) (interface{}, error) {
	objMaker := MakeInjectBuilderFunc(rtElem)
	return func(ctn Container) (interface{}, error) {
		rtValue := reflect.New(rtElem)
		dst := rtValue.Interface()
		return objMaker(ctn, dst)
	}
}

// MakeInjectBuilderFunc is EXPENSIVE consider making direct calls to GetByType and GetManyByType directly
func MakeInjectBuilderFunc(rt reflect.Type) func(ctn Container, dst interface{}) (interface{}, error) {
	setters := []func(ctn Container, dst interface{}){}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tag := field.Tag
		tags, err := structtag.Parse(string(tag))
		if err != nil {
			panic(err)
		}
		if !hasKey("inject", tags) {
			continue
		}
		rtField := field.Type
		fieldName := field.Name
		switch kind := rtField.Kind(); kind {
		case reflect.Array, reflect.Slice:
			sliceElem := rtField.Elem()
			setters = append(setters, func(ctn Container, dst interface{}) {
				v := reflect.ValueOf(dst).Elem()
				f := v.FieldByName(fieldName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						sliceType := reflect.SliceOf(sliceElem)
						sliceV := reflect.New(sliceType).Elem()
						objs := ctn.GetManyByType(sliceElem)
						for _, obj := range objs {
							tsV := reflect.ValueOf(obj)
							sliceV = reflect.Append(sliceV, tsV)
						}
						f.Set(sliceV)
					}
				}
			})
		default:
			setters = append(setters, func(ctn Container, dst interface{}) {
				v := reflect.ValueOf(dst).Elem()
				f := v.FieldByName(fieldName)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						obj := ctn.GetByType(rtField)
						objValue := reflect.ValueOf(obj)
						f.Set(objValue)
					}
				}
			})
		}
	}

	return func(ctn Container, dst interface{}) (interface{}, error) {
		rtDst := reflect.TypeOf(dst)
		switch kind := rtDst.Kind(); kind {
		case reflect.Ptr:
			rtElem := rtDst.Elem()
			if rtElem != rt {
				panic("type mismatch")
			}
		default:
			panic("Must be a ptr to a struct. type mismatch")
		}
		for _, setter := range setters {
			setter(ctn, dst)
		}
		return dst, nil
	}
}

func hasKey(key string, tags *structtag.Tags) bool {
	for _, k := range tags.Keys() {
		if k == key {
			return true
		}
	}
	return false
}
