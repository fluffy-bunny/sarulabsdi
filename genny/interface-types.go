package genny

import (
	"context"
	"reflect"
	"strings"

	"github.com/cheekybits/genny/generic"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog"
)

// InterfaceType ...
type InterfaceType generic.Type

// ReflectTypeInterfaceType used when your service claims to implement InterfaceType
var ReflectTypeInterfaceType = di.GetInterfaceReflectType((*InterfaceType)(nil))

// AddSingletonInterfaceType adds a type that implements InterfaceType
func AddSingletonInterfaceType(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddSingleton(builder, implType, implementedTypes...)
}

// AddSingletonInterfaceTypeWithMetadata adds a type that implements InterfaceType
func AddSingletonInterfaceTypeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddSingletonWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddSingletonInterfaceTypeByObj adds a prebuilt obj
func AddSingletonInterfaceTypeByObj(builder *di.Builder, obj interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", reflect.TypeOf(obj), _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "obj",
		})
	di.AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonInterfaceTypeByObjWithMetadata adds a prebuilt obj
func AddSingletonInterfaceTypeByObjWithMetadata(builder *di.Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", reflect.TypeOf(obj), _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "obj",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonInterfaceTypeByFunc adds a type by a custom func
func AddSingletonInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddSingletonWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddSingletonInterfaceTypeByFuncWithMetadata adds a type by a custom func
func AddSingletonInterfaceTypeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SINGLETON", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddSingletonWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddTransientInterfaceType adds a type that implements InterfaceType
func AddTransientInterfaceType(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("TRANSIENT", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		})

	di.AddTransientWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddTransientInterfaceTypeWithMetadata adds a type that implements InterfaceType
func AddTransientInterfaceTypeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("TRANSIENT", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddTransientInterfaceTypeByFunc adds a type by a custom func
func AddTransientInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("TRANSIENT", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		})

	di.AddTransientWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddTransientInterfaceTypeByFuncWithMetadata adds a type by a custom func
func AddTransientInterfaceTypeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("TRANSIENT", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddTransientWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
}

// AddScopedInterfaceType adds a type that implements InterfaceType
func AddScopedInterfaceType(builder *di.Builder, implType reflect.Type, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SCOPED", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		})
	di.AddScopedWithImplementedTypes(builder, implType, implementedTypes...)
}

// AddScopedInterfaceTypeWithMetadata adds a type that implements InterfaceType
func AddScopedInterfaceTypeWithMetadata(builder *di.Builder, implType reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SCOPED", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "type",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})
	di.AddScopedWithImplementedTypesWithMetadata(builder, implType, metaData, implementedTypes...)
}

// AddScopedInterfaceTypeByFunc adds a type by a custom func
func AddScopedInterfaceTypeByFunc(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SCOPED", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		})
	di.AddScopedWithImplementedTypesByFunc(builder, implType, build, implementedTypes...)
}

// AddScopedInterfaceTypeByFuncWithMetadata adds a type by a custom func
func AddScopedInterfaceTypeByFuncWithMetadata(builder *di.Builder, implType reflect.Type, build func(ctn di.Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) {
	implementedTypes = append(implementedTypes, ReflectTypeInterfaceType)
	_logAddInterfaceType("SCOPED", implType, _getImplementedInterfaceTypeNames(implementedTypes...),
		_logInterfaceTypeExtra{
			Name:  "DI-BY",
			Value: "func",
		},
		_logInterfaceTypeExtra{
			Name:  "DI-M",
			Value: metaData,
		})

	di.AddScopedWithImplementedTypesByFuncWithMetadata(builder, implType, build, metaData, implementedTypes...)
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

// GetInterfaceTypeDefinition returns that last definition registered that this container can provide
func GetInterfaceTypeDefinition(ctn di.Container) *di.Def {
	def := ctn.GetDefinitionByType(ReflectTypeInterfaceType)
	return def
}

// GetInterfaceTypeDefinitions returns all definitions that this container can provide
func GetInterfaceTypeDefinitions(ctn di.Container) []*di.Def {
	defs := ctn.GetDefinitionsByType(ReflectTypeInterfaceType)
	return defs
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

type _logInterfaceTypeExtra struct {
	Name  string
	Value interface{}
}

func _logAddInterfaceType(scopeType string, implType reflect.Type, interfaces string, extra ..._logInterfaceTypeExtra) {
	log := zerolog.Ctx(context.Background()).With().Logger()
	log = log.With().
		Str("DI", scopeType).
		Str("DI-I", interfaces).
		Str("DI-B", implType.Elem().String()).Logger()

	for _, extra := range extra {
		log = log.With().Interface(extra.Name, extra.Value).Logger()
	}

	log.Info().Send()
}
func _getImplementedInterfaceTypeNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
