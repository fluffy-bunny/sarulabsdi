package genny

import (
	"reflect"

	"github.com/cheekybits/genny/generic"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// FuncType ...
type FuncType generic.Type

// ReflectTypeFuncType used when your service claims to implement FuncType
var ReflectTypeFuncType = reflect.TypeOf(FuncType(nil))

// AddSingletonFuncTypeFunc adds a func to the DI
func AddFuncTypeFunc(builder *di.Builder, fnc FuncType) {
	di.AddFunc(builder, fnc)
}

// RemoveAllFuncTypeFunc removes all FuncType functions from the DI
func RemoveAllFuncTypeFunc(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeFuncType)
}

// GetFuncTypeFromContainer alternative to SafeGetFuncTypeFromContainer but panics of object is not present
func GetFuncTypeFromContainer(ctn di.Container) FuncType {
	return ctn.GetByType(ReflectTypeFuncType).(FuncType)
}

// GetManyFuncTypeFromContainer alternative to SafeGetManyFuncTypeFromContainer but panics of object is not present
func GetManyFuncTypeFromContainer(ctn di.Container) []FuncType {
	objs := ctn.GetManyByType(ReflectTypeFuncType)
	var results []FuncType
	for _, obj := range objs {
		results = append(results, obj.(FuncType))
	}
	return results
}

// SafeGetFuncTypeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetFuncTypeFromContainer(ctn di.Container) (FuncType, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeFuncType)
	if err != nil {
		return nil, err
	}
	return obj.(FuncType), nil
}

// SafeGetManyFuncTypeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyFuncTypeFromContainer(ctn di.Container) ([]FuncType, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeFuncType)
	if err != nil {
		return nil, err
	}
	var results []FuncType
	for _, obj := range objs {
		results = append(results, obj.(FuncType))
	}
	return results, nil
}
