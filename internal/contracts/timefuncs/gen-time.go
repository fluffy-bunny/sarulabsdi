// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package timefuncs

import (
	"reflect"

	di "github.com/fluffy-bunny/sarulabsdi"
)

// ReflectTypeITime used when your service claims to implement ITime
var ReflectTypeITime = di.GetInterfaceReflectType((*ITime)(nil))

// AddSingletonITimeByObj adds a prebuilt obj
func AddSingletonITimeByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeITime)
}

// AddSingletonITime adds a type that implements ITime
func AddSingletonITime(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeITime)
}

// AddSingletonITimeByFunc adds a type by a custom func
func AddSingletonITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeITime)
}

// AddTransientITime adds a type that implements ITime
func AddTransientITime(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeITime)
}

// AddTransientITimeByFunc adds a type by a custom func
func AddTransientITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeITime)
}

// AddScopedITime adds a type that implements ITime
func AddScopedITime(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeITime)
}

// AddScopedITimeByFunc adds a type by a custom func
func AddScopedITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeITime)
}

// RemoveAllITime removes all ITime from the DI
func RemoveAllITime(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeITime)
}

// GetITimeFromContainer alternative to SafeGetITimeFromContainer but panics of object is not present
func GetITimeFromContainer(ctn di.Container) ITime {
	return ctn.GetByType(ReflectTypeITime).(ITime)
}

// GetManyITimeFromContainer alternative to SafeGetManyITimeFromContainer but panics of object is not present
func GetManyITimeFromContainer(ctn di.Container) []ITime {
	objs := ctn.GetManyByType(ReflectTypeITime)
	var results []ITime
	for _, obj := range objs {
		results = append(results, obj.(ITime))
	}
	return results
}

// SafeGetITimeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetITimeFromContainer(ctn di.Container) (ITime, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeITime)
	if err != nil {
		return nil, err
	}
	return obj.(ITime), nil
}

// SafeGetManyITimeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyITimeFromContainer(ctn di.Container) ([]ITime, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeITime)
	if err != nil {
		return nil, err
	}
	var results []ITime
	for _, obj := range objs {
		results = append(results, obj.(ITime))
	}
	return results, nil
}

// ReflectTypeTimeNow used when your service claims to implement TimeNow
var ReflectTypeTimeNow = di.GetInterfaceReflectType((*TimeNow)(nil))

// AddSingletonTimeNowByObj adds a prebuilt obj
func AddSingletonTimeNowByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeTimeNow)
}

// AddSingletonTimeNow adds a type that implements TimeNow
func AddSingletonTimeNow(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeTimeNow)
}

// AddSingletonTimeNowByFunc adds a type by a custom func
func AddSingletonTimeNowByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeTimeNow)
}

// AddTransientTimeNow adds a type that implements TimeNow
func AddTransientTimeNow(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeTimeNow)
}

// AddTransientTimeNowByFunc adds a type by a custom func
func AddTransientTimeNowByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeTimeNow)
}

// AddScopedTimeNow adds a type that implements TimeNow
func AddScopedTimeNow(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeTimeNow)
}

// AddScopedTimeNowByFunc adds a type by a custom func
func AddScopedTimeNowByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeTimeNow)
}

// RemoveAllTimeNow removes all TimeNow from the DI
func RemoveAllTimeNow(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeTimeNow)
}

// GetTimeNowFromContainer alternative to SafeGetTimeNowFromContainer but panics of object is not present
func GetTimeNowFromContainer(ctn di.Container) TimeNow {
	return ctn.GetByType(ReflectTypeTimeNow).(TimeNow)
}

// GetManyTimeNowFromContainer alternative to SafeGetManyTimeNowFromContainer but panics of object is not present
func GetManyTimeNowFromContainer(ctn di.Container) []TimeNow {
	objs := ctn.GetManyByType(ReflectTypeTimeNow)
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results
}

// SafeGetTimeNowFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetTimeNowFromContainer(ctn di.Container) (TimeNow, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeTimeNow)
	if err != nil {
		return nil, err
	}
	return obj.(TimeNow), nil
}

// SafeGetManyTimeNowFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyTimeNowFromContainer(ctn di.Container) ([]TimeNow, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeTimeNow)
	if err != nil {
		return nil, err
	}
	var results []TimeNow
	for _, obj := range objs {
		results = append(results, obj.(TimeNow))
	}
	return results, nil
}
