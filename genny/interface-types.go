package genny

import (
	"reflect"

	"github.com/cheekybits/genny/generic"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// InterfaceType ...
type InterfaceType generic.Type

// ReflectTypeInterfaceType used when your service claims to implement InterfaceType
var ReflectTypeInterfaceType = di.GetInterfaceReflectType((*InterfaceType)(nil))

// AddSingletonInterfaceTypeByObj adds a prebuilt obj
func AddSingletonInterfaceTypeByObj(builder *di.Builder, obj interface{}) {
	di.AddSingletonWithImplementedTypesByObj(builder, obj, ReflectTypeInterfaceType)
}

// AddSingletonInterfaceType adds a type that implements InterfaceType
func AddSingletonInterfaceType(builder *di.Builder, implType reflect.Type) {
	di.AddSingletonWithImplementedTypes(builder, implType, ReflectTypeInterfaceType)
}

// AddSingletonInterfaceTypeByFunc adds a type by a custom func
func AddSingletonInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, ReflectTypeInterfaceType)
}

// AddTransientInterfaceType adds a type that implements InterfaceType
func AddTransientInterfaceType(builder *di.Builder, implType reflect.Type) {
	di.AddTransientWithImplementedTypes(builder, implType, ReflectTypeInterfaceType)
}

// AddTransientInterfaceTypeByFunc adds a type by a custom func
func AddTransientInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, ReflectTypeInterfaceType)
}

// AddScopedInterfaceType adds a type that implements InterfaceType
func AddScopedInterfaceType(builder *di.Builder, implType reflect.Type) {
	di.AddScopedWithImplementedTypes(builder, implType, ReflectTypeInterfaceType)
}

// AddScopedInterfaceTypeByFunc adds a type by a custom func
func AddScopedInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error)) {
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, ReflectTypeInterfaceType)
}

// RemoveAllInterfaceType removes all InterfaceType from the DI
func RemoveAllInterfaceType(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeInterfaceType)
}

// GetInterfaceTypeFromContainer alternative to SafeGetInterfaceTypeFromContainer but panics of object is not present
func GetInterfaceTypeFromContainer(ctn di.Container) InterfaceType {
	return ctn.GetByType(ReflectTypeInterfaceType).(InterfaceType)
}

// GetManyInterfaceTypeFromContainer alternative to SafeGetManyInterfaceTypeFromContainer but panics of object is not present
func GetManyInterfaceTypeFromContainer(ctn di.Container) []InterfaceType {
	objs := ctn.GetManyByType(ReflectTypeInterfaceType)
	var results []InterfaceType
	for _, obj := range objs {
		results = append(results, obj.(InterfaceType))
	}
	return results
}

// SafeGetInterfaceTypeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetInterfaceTypeFromContainer(ctn di.Container) (InterfaceType, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeInterfaceType)
	if err != nil {
		return nil, err
	}
	return obj.(InterfaceType), nil
}

// SafeGetManyInterfaceTypeFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyInterfaceTypeFromContainer(ctn di.Container) ([]InterfaceType, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeInterfaceType)
	if err != nil {
		return nil, err
	}
	var results []InterfaceType
	for _, obj := range objs {
		results = append(results, obj.(InterfaceType))
	}
	return results, nil
}
