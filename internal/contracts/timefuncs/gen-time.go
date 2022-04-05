// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package timefuncs

import (
	"reflect"
	"strings"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// ReflectTypeITime used when your service claims to implement ITime
var ReflectTypeITime = di.GetInterfaceReflectType((*ITime)(nil))

// AddSingletonITime adds a type that implements ITime
func AddSingletonITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonITimeWithMetadata adds a type that implements ITime
func AddSingletonITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonITimeByObj adds a prebuilt obj
func AddSingletonITimeByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", reflect.TypeOf(obj), _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonITimeByObjWithMetadata adds a prebuilt obj
func AddSingletonITimeByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", reflect.TypeOf(obj), _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonITimeByFunc adds a type by a custom func
func AddSingletonITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonITimeByFuncWithMetadata adds a type by a custom func
func AddSingletonITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SINGLETON", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientITime adds a type that implements ITime
func AddTransientITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("TRANSIENT", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientITimeWithMetadata adds a type that implements ITime
func AddTransientITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("TRANSIENT", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientITimeByFunc adds a type by a custom func
func AddTransientITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("TRANSIENT", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientITimeByFuncWithMetadata adds a type by a custom func
func AddTransientITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("TRANSIENT", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedITime adds a type that implements ITime
func AddScopedITime(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SCOPED", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedITimeWithMetadata adds a type that implements ITime
func AddScopedITimeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SCOPED", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedITimeByFunc adds a type by a custom func
func AddScopedITimeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SCOPED", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedITimeByFuncWithMetadata adds a type by a custom func
func AddScopedITimeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeITime)
	_logAddITime("SCOPED", implType, _getImplementedITimeNames(implementedTypes...),
		_logITimeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logITimeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
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

// GetITimeDefinition returns that last definition registered that this container can provide
func GetITimeDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeITime)
	return def
}

// GetITimeDefinitions returns all definitions that this container can provide
func GetITimeDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeITime)
	return defs
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

type _logITimeExtra struct {
	Name  string
	Value interface{}
}

func _logAddITime(scopeType string, implType reflect.Type, interfaces string, extra ..._logITimeExtra) {
	infoEvent := log.Info().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String())

	for _, extra := range extra {
		infoEvent = infoEvent.Interface(extra.Name, extra.Value)
	}

	infoEvent.Send()

}
func _getImplementedITimeNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
