// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package something

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeISomething used when your service claims to implement ISomething
var ReflectTypeISomething = di.GetInterfaceReflectType((*ISomething)(nil))

// AddSingletonISomethingByObj adds a prebuilt obj
func AddSingletonISomethingByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeISomething)
}

// AddSingletonISomething adds a type that implements ISomething
func AddSingletonISomething(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeISomething)
}

// AddSingletonISomethingByFunc adds a type by a custom func
func AddSingletonISomethingByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething)
}

// AddTransientISomething adds a type that implements ISomething
func AddTransientISomething(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeISomething)
}

// AddTransientISomethingByFunc adds a type by a custom func
func AddTransientISomethingByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething)
}

// AddScopedISomething adds a type that implements ISomething
func AddScopedISomething(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeISomething)
}

// AddScopedISomethingByFunc adds a type by a custom func
func AddScopedISomethingByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething)
}

// RemoveAllISomething removes all ISomething from the DI
func RemoveAllISomething(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeISomething)
}

// GetISomethingFromContainer alternative to SafeGetISomethingFromContainer but panics of object is not present
func GetISomethingFromContainer(ctn di.Container) ISomething {
	return ctn.GetByType(ReflectTypeISomething).(ISomething)
}

// GetManyISomethingFromContainer alternative to SafeGetManyISomethingFromContainer but panics of object is not present
func GetManyISomethingFromContainer(ctn di.Container) []ISomething {
	objs := ctn.GetManyByType(ReflectTypeISomething)
	var results []ISomething
	for _, obj := range objs {
		results = append(results, obj.(ISomething))
	}
	return results
}

// SafeGetISomethingFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetISomethingFromContainer(ctn di.Container) (ISomething, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeISomething)
	if err != nil {
		return nil, err
	}
	return obj.(ISomething), nil
}

// SafeGetManyISomethingFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyISomethingFromContainer(ctn di.Container) ([]ISomething, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeISomething)
	if err != nil {
		return nil, err
	}
	var results []ISomething
	for _, obj := range objs {
		results = append(results, obj.(ISomething))
	}
	return results, nil
}

// ReflectTypeISomething2 used when your service claims to implement ISomething2
var ReflectTypeISomething2 = di.GetInterfaceReflectType((*ISomething2)(nil))

// AddSingletonISomething2ByObj adds a prebuilt obj
func AddSingletonISomething2ByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeISomething2)
}

// AddSingletonISomething2 adds a type that implements ISomething2
func AddSingletonISomething2(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeISomething2)
}

// AddSingletonISomething2ByFunc adds a type by a custom func
func AddSingletonISomething2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething2)
}

// AddTransientISomething2 adds a type that implements ISomething2
func AddTransientISomething2(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeISomething2)
}

// AddTransientISomething2ByFunc adds a type by a custom func
func AddTransientISomething2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething2)
}

// AddScopedISomething2 adds a type that implements ISomething2
func AddScopedISomething2(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeISomething2)
}

// AddScopedISomething2ByFunc adds a type by a custom func
func AddScopedISomething2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething2)
}

// RemoveAllISomething2 removes all ISomething2 from the DI
func RemoveAllISomething2(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeISomething2)
}

// GetISomething2FromContainer alternative to SafeGetISomething2FromContainer but panics of object is not present
func GetISomething2FromContainer(ctn di.Container) ISomething2 {
	return ctn.GetByType(ReflectTypeISomething2).(ISomething2)
}

// GetManyISomething2FromContainer alternative to SafeGetManyISomething2FromContainer but panics of object is not present
func GetManyISomething2FromContainer(ctn di.Container) []ISomething2 {
	objs := ctn.GetManyByType(ReflectTypeISomething2)
	var results []ISomething2
	for _, obj := range objs {
		results = append(results, obj.(ISomething2))
	}
	return results
}

// SafeGetISomething2FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetISomething2FromContainer(ctn di.Container) (ISomething2, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeISomething2)
	if err != nil {
		return nil, err
	}
	return obj.(ISomething2), nil
}

// SafeGetManyISomething2FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyISomething2FromContainer(ctn di.Container) ([]ISomething2, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeISomething2)
	if err != nil {
		return nil, err
	}
	var results []ISomething2
	for _, obj := range objs {
		results = append(results, obj.(ISomething2))
	}
	return results, nil
}

// ReflectTypeISomething3 used when your service claims to implement ISomething3
var ReflectTypeISomething3 = di.GetInterfaceReflectType((*ISomething3)(nil))

// AddSingletonISomething3ByObj adds a prebuilt obj
func AddSingletonISomething3ByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeISomething3)
}

// AddSingletonISomething3 adds a type that implements ISomething3
func AddSingletonISomething3(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeISomething3)
}

// AddSingletonISomething3ByFunc adds a type by a custom func
func AddSingletonISomething3ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething3)
}

// AddTransientISomething3 adds a type that implements ISomething3
func AddTransientISomething3(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeISomething3)
}

// AddTransientISomething3ByFunc adds a type by a custom func
func AddTransientISomething3ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething3)
}

// AddScopedISomething3 adds a type that implements ISomething3
func AddScopedISomething3(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeISomething3)
}

// AddScopedISomething3ByFunc adds a type by a custom func
func AddScopedISomething3ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeISomething3)
}

// RemoveAllISomething3 removes all ISomething3 from the DI
func RemoveAllISomething3(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeISomething3)
}

// GetISomething3FromContainer alternative to SafeGetISomething3FromContainer but panics of object is not present
func GetISomething3FromContainer(ctn di.Container) ISomething3 {
	return ctn.GetByType(ReflectTypeISomething3).(ISomething3)
}

// GetManyISomething3FromContainer alternative to SafeGetManyISomething3FromContainer but panics of object is not present
func GetManyISomething3FromContainer(ctn di.Container) []ISomething3 {
	objs := ctn.GetManyByType(ReflectTypeISomething3)
	var results []ISomething3
	for _, obj := range objs {
		results = append(results, obj.(ISomething3))
	}
	return results
}

// SafeGetISomething3FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetISomething3FromContainer(ctn di.Container) (ISomething3, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeISomething3)
	if err != nil {
		return nil, err
	}
	return obj.(ISomething3), nil
}

// SafeGetManyISomething3FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyISomething3FromContainer(ctn di.Container) ([]ISomething3, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeISomething3)
	if err != nil {
		return nil, err
	}
	var results []ISomething3
	for _, obj := range objs {
		results = append(results, obj.(ISomething3))
	}
	return results, nil
}
