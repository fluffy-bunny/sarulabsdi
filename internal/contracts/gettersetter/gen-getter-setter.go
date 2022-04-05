// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package gettersetter

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeIGetterSetter used when your service claims to implement IGetterSetter
var ReflectTypeIGetterSetter = di.GetInterfaceReflectType((*IGetterSetter)(nil))

// AddSingletonIGetterSetter adds a type that implements IGetterSetter
func AddSingletonIGetterSetter(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIGetterSetterWithMetadata adds a type that implements IGetterSetter
func AddSingletonIGetterSetterWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIGetterSetterByObj adds a prebuilt obj
func AddSingletonIGetterSetterByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", reflect.TypeOf(obj), _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIGetterSetterByObjWithMetadata adds a prebuilt obj
func AddSingletonIGetterSetterByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", reflect.TypeOf(obj), _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIGetterSetterByFunc adds a type by a custom func
func AddSingletonIGetterSetterByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIGetterSetterByFuncWithMetadata adds a type by a custom func
func AddSingletonIGetterSetterByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SINGLETON", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIGetterSetter adds a type that implements IGetterSetter
func AddTransientIGetterSetter(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("TRANSIENT", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIGetterSetterWithMetadata adds a type that implements IGetterSetter
func AddTransientIGetterSetterWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("TRANSIENT", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIGetterSetterByFunc adds a type by a custom func
func AddTransientIGetterSetterByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("TRANSIENT", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIGetterSetterByFuncWithMetadata adds a type by a custom func
func AddTransientIGetterSetterByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("TRANSIENT", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIGetterSetter adds a type that implements IGetterSetter
func AddScopedIGetterSetter(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SCOPED", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIGetterSetterWithMetadata adds a type that implements IGetterSetter
func AddScopedIGetterSetterWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SCOPED", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIGetterSetterByFunc adds a type by a custom func
func AddScopedIGetterSetterByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SCOPED", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIGetterSetterByFuncWithMetadata adds a type by a custom func
func AddScopedIGetterSetterByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter)
	_logAddIGetterSetter("SCOPED", implType, _getImplementedIGetterSetterNames(implementedTypes...),
		_logIGetterSetterExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetterExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIGetterSetter removes all IGetterSetter from the DI
func RemoveAllIGetterSetter(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIGetterSetter)
}

// GetIGetterSetterFromContainer alternative to SafeGetIGetterSetterFromContainer but panics of object is not present
func GetIGetterSetterFromContainer(ctn di.Container) IGetterSetter {
	return ctn.GetByType(ReflectTypeIGetterSetter).(IGetterSetter)
}

// GetManyIGetterSetterFromContainer alternative to SafeGetManyIGetterSetterFromContainer but panics of object is not present
func GetManyIGetterSetterFromContainer(ctn di.Container) []IGetterSetter {
	objs := ctn.GetManyByType(ReflectTypeIGetterSetter)
	var results []IGetterSetter
	for _, obj := range objs {
		results = append(results, obj.(IGetterSetter))
	}
	return results
}

// SafeGetIGetterSetterFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIGetterSetterFromContainer(ctn di.Container) (IGetterSetter, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGetterSetter)
	if err != nil {
		return nil, err
	}
	return obj.(IGetterSetter), nil
}

// GetIGetterSetterDefinition returns that last definition registered that this container can provide
func GetIGetterSetterDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIGetterSetter)
	return def
}

// GetIGetterSetterDefinitions returns all definitions that this container can provide
func GetIGetterSetterDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIGetterSetter)
	return defs
}

// SafeGetManyIGetterSetterFromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIGetterSetterFromContainer(ctn di.Container) ([]IGetterSetter, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIGetterSetter)
	if err != nil {
		return nil, err
	}
	var results []IGetterSetter
	for _, obj := range objs {
		results = append(results, obj.(IGetterSetter))
	}
	return results, nil
}

type _logIGetterSetterExtra struct {
	Name  string
	Value interface{}
}

func _logAddIGetterSetter(scopeType string, implType reflect.Type, interfaces string, extra ..._logIGetterSetterExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIGetterSetterNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}

// ReflectTypeIGetterSetter2 used when your service claims to implement IGetterSetter2
var ReflectTypeIGetterSetter2 = di.GetInterfaceReflectType((*IGetterSetter2)(nil))

// AddSingletonIGetterSetter2 adds a type that implements IGetterSetter2
func AddSingletonIGetterSetter2(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonIGetterSetter2WithMetadata adds a type that implements IGetterSetter2
func AddSingletonIGetterSetter2WithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonIGetterSetter2ByObj adds a prebuilt obj
func AddSingletonIGetterSetter2ByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", reflect.TypeOf(obj), _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonIGetterSetter2ByObjWithMetadata adds a prebuilt obj
func AddSingletonIGetterSetter2ByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", reflect.TypeOf(obj), _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonIGetterSetter2ByFunc adds a type by a custom func
func AddSingletonIGetterSetter2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonIGetterSetter2ByFuncWithMetadata adds a type by a custom func
func AddSingletonIGetterSetter2ByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SINGLETON", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientIGetterSetter2 adds a type that implements IGetterSetter2
func AddTransientIGetterSetter2(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("TRANSIENT", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientIGetterSetter2WithMetadata adds a type that implements IGetterSetter2
func AddTransientIGetterSetter2WithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("TRANSIENT", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientIGetterSetter2ByFunc adds a type by a custom func
func AddTransientIGetterSetter2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("TRANSIENT", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientIGetterSetter2ByFuncWithMetadata adds a type by a custom func
func AddTransientIGetterSetter2ByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("TRANSIENT", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedIGetterSetter2 adds a type that implements IGetterSetter2
func AddScopedIGetterSetter2(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SCOPED", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedIGetterSetter2WithMetadata adds a type that implements IGetterSetter2
func AddScopedIGetterSetter2WithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SCOPED", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedIGetterSetter2ByFunc adds a type by a custom func
func AddScopedIGetterSetter2ByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SCOPED", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedIGetterSetter2ByFuncWithMetadata adds a type by a custom func
func AddScopedIGetterSetter2ByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeIGetterSetter2)
	_logAddIGetterSetter2("SCOPED", implType, _getImplementedIGetterSetter2Names(implementedTypes...),
		_logIGetterSetter2Extra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logIGetterSetter2Extra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// RemoveAllIGetterSetter2 removes all IGetterSetter2 from the DI
func RemoveAllIGetterSetter2(builder *di.Builder) {
	builder.RemoveAllByType(ReflectTypeIGetterSetter2)
}

// GetIGetterSetter2FromContainer alternative to SafeGetIGetterSetter2FromContainer but panics of object is not present
func GetIGetterSetter2FromContainer(ctn di.Container) IGetterSetter2 {
	return ctn.GetByType(ReflectTypeIGetterSetter2).(IGetterSetter2)
}

// GetManyIGetterSetter2FromContainer alternative to SafeGetManyIGetterSetter2FromContainer but panics of object is not present
func GetManyIGetterSetter2FromContainer(ctn di.Container) []IGetterSetter2 {
	objs := ctn.GetManyByType(ReflectTypeIGetterSetter2)
	var results []IGetterSetter2
	for _, obj := range objs {
		results = append(results, obj.(IGetterSetter2))
	}
	return results
}

// SafeGetIGetterSetter2FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetIGetterSetter2FromContainer(ctn di.Container) (IGetterSetter2, error) {
	obj, err := ctn.SafeGetByType(ReflectTypeIGetterSetter2)
	if err != nil {
		return nil, err
	}
	return obj.(IGetterSetter2), nil
}

// GetIGetterSetter2Definition returns that last definition registered that this container can provide
func GetIGetterSetter2Definition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeIGetterSetter2)
	return def
}

// GetIGetterSetter2Definitions returns all definitions that this container can provide
func GetIGetterSetter2Definitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeIGetterSetter2)
	return defs
}

// SafeGetManyIGetterSetter2FromContainer trys to get the object by type, will not panic, returns nil and error
func SafeGetManyIGetterSetter2FromContainer(ctn di.Container) ([]IGetterSetter2, error) {
	objs, err := ctn.SafeGetManyByType(ReflectTypeIGetterSetter2)
	if err != nil {
		return nil, err
	}
	var results []IGetterSetter2
	for _, obj := range objs {
		results = append(results, obj.(IGetterSetter2))
	}
	return results, nil
}

type _logIGetterSetter2Extra struct {
	Name  string
	Value interface{}
}

func _logAddIGetterSetter2(scopeType string, implType reflect.Type, interfaces string, extra ..._logIGetterSetter2Extra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedIGetterSetter2Names(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
