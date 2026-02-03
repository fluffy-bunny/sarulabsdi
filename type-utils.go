package di

import (
	"reflect"
	"strings"
)

// AddTransient adds a simple transient type
func AddTransient(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddTransientWithImplementedTypes(builder, rt, implementedTypes...)
}
func AddTransientT[T any](builder *Builder, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddTransient(builder, reflectType, implementedTypes...)
}

// AddTransientWithMetadata adds a simple transient type
func AddTransientWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddTransientWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}
func AddTransientWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddTransientWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddTransientWithImplementedTypes adds a type and its implemented interfaces
func AddTransientWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
	}
	builder.Add(def)
	return def
}

// AddTransientWithImplementedTypes adds a type and its implemented interfaces
func AddTransientWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddTransientWithImplementedTypesWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddTransientWithImplementedTypesWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddTransientWithImplementedTypesByFunc adds a type and its implemented interfaces
func AddTransientWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		Build:            build,
	}
	builder.Add(def)
	return def
}
func AddTransientWithImplementedTypesByFuncT[T any](builder *Builder, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddTransientWithImplementedTypesByFunc(builder, reflectType, build, implementedTypes...)
}

// AddTransientWithImplementedTypesByFuncWithMetadata adds a type and its implemented interfaces
func AddTransientWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddTransientWithImplementedTypesByFuncWithMetadataT[T any](builder *Builder, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddTransientWithImplementedTypesByFuncWithMetadata(builder, reflectType, build, metaData, implementedTypes...)
}

// AddScoped adds a simple scoped type
func AddScoped(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddScopedWithImplementedTypes(builder, rt, implementedTypes...)
}
func AddScopedT[T any](builder *Builder, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScoped(builder, reflectType, implementedTypes...)
}

// AddScopedWithMetadata adds a simple scoped type
func AddScopedWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddScopedWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}
func AddScopedWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScopedWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddScopedWithImplementedTypes adds a type and its implemented interfaces
func AddScopedWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
	}
	builder.Add(def)
	return def
}
func AddScopedWithImplementedTypesT[T any](builder *Builder, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScopedWithImplementedTypes(builder, reflectType, implementedTypes...)
}

// AddScopedWithImplementedTypesWithMetadata adds a type and its implemented interfaces
func AddScopedWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddScopedWithImplementedTypesWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScopedWithImplementedTypesWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddScopedWithImplementedTypesByFunc adds a type and its implemented interfaces
func AddScopedWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		Build:            build,
	}
	builder.Add(def)
	return def
}

func AddScopedWithImplementedTypesByFuncT[T any](builder *Builder, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScopedWithImplementedTypesByFunc(builder, reflectType, build, implementedTypes...)
}

// AddScopedWithImplementedTypesByFuncWithMetadata adds a type and its implemented interfaces
func AddScopedWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddScopedWithImplementedTypesByFuncWithMetadataT[T any](builder *Builder, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddScopedWithImplementedTypesByFuncWithMetadata(builder, reflectType, build, metaData, implementedTypes...)
}

// AddFunc adds a simple singleton type
func AddFunc(builder *Builder, fnc interface{}) Def {
	rt := reflect.TypeOf(fnc)
	implementedTypes2 := NewTypeSet()
	implementedTypes2.Add(rt)
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Func:             fnc,
	}
	builder.Add(def)
	return def
}

// AddFunc adds a simple singleton type
func AddFuncWithMetadata(builder *Builder, fnc interface{}, metaData map[string]interface{}) Def {
	rt := reflect.TypeOf(fnc)
	implementedTypes2 := NewTypeSet()
	implementedTypes2.Add(rt)
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Func:             fnc,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddSingleton adds a simple singleton type
func AddSingleton(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddSingletonWithImplementedTypes(builder, rt, implementedTypes...)
}
func AddSingletonT[T any](builder *Builder, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingleton(builder, reflectType, implementedTypes...)
}

// AddSingletonWithMetadata adds a simple singleton type
func AddSingletonWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddSingletonWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}

func AddSingletonWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddSingletonWithImplementedTypes adds a prebuilt obj
func AddSingletonWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
	}
	builder.Add(def)
	return def
}

func AddSingletonWithImplementedTypesT[T any](builder *Builder, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypes(builder, reflectType, implementedTypes...)
}

// AddSingletonWithImplementedTypesWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddSingletonWithImplementedTypesWithMetadataT[T any](builder *Builder, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypesWithMetadata(builder, reflectType, metaData, implementedTypes...)
}

// AddSingletonWithImplementedTypesByFunc adds a prebuilt obj
func AddSingletonWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build:            build,
	}
	builder.Add(def)
	return def
}

func AddSingletonWithImplementedTypesByFuncT[T any](builder *Builder, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypesByFunc(builder, reflectType, build, implementedTypes...)
}

// AddSingletonWithImplementedTypesByFuncWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}
func AddSingletonWithImplementedTypesByFuncWithMetadataT[T any](builder *Builder, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypesByFuncWithMetadata(builder, reflectType, build, metaData, implementedTypes...)
}

// AddSingletonWithImplementedTypesByObj adds a prebuilt obj
func AddSingletonWithImplementedTypesByObj(builder *Builder, obj interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             reflect.TypeOf(obj),
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	}
	builder.Add(def)
	return def
}

func AddSingletonWithImplementedTypesByObjT[T any](builder *Builder, obj interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypesByObj(builder, obj, implementedTypes...)
}

// AddSingletonWithImplementedTypesByObjWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesByObjWithMetadata(builder *Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             reflect.TypeOf(obj),
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
		MetaData: metaData,
	}
	builder.Add(def)
	return def
}

func AddSingletonWithImplementedTypesByObjWithMetadataT[T any](builder *Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	implementedTypes = append(implementedTypes, reflectType)
	return AddSingletonWithImplementedTypesByObjWithMetadata(builder, obj, metaData, implementedTypes...)
}

// AddSingletonTypeByObj adds a prebuilt object by its type
func AddSingletonTypeByObj(builder *Builder, obj interface{}) Def {
	def := Def{
		Scope: App, // Singleton
		Type:  reflect.TypeOf(obj),
		// SafeInject MUST be true because the following CAN be nil
		// Dialer kafkaContracts.IKafkaDialer `inject:""`
		Unshared: false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	}
	builder.Add(def)
	return def
}

// AddSingletonTypeByObjWithMetadata adds a prebuilt object by its type
func AddSingletonTypeByObjWithMetadata(builder *Builder, obj interface{}, metaData map[string]interface{}) Def {
	def := Def{
		Scope: App, // Singleton
		Type:  reflect.TypeOf(obj),
		// SafeInject MUST be true because the following CAN be nil
		// Dialer kafkaContracts.IKafkaDialer `inject:""`
		Unshared: false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
		MetaData: metaData,
	}
	builder.Add(def)
	return def
}

func GetFromContainer[T any](ctn Container) T {
	reflectType := GetInterfaceReflectType((*T)(nil))
	return ctn.GetByType(reflectType).(T)
}

func GetManyFromContainerT[T any](ctn Container) []T {
	reflectType := GetInterfaceReflectType((*T)(nil))
	objs := ctn.GetManyByType(reflectType)
	result := make([]T, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(T))
	}
	return result
}

func SafeGetFromContainerT[T any](ctn Container) (T, error) {
	reflectType := GetInterfaceReflectType((*T)(nil))
	obj, err := ctn.SafeGetByType(reflectType)
	if err != nil {
		var zero T
		return zero, err
	}
	return obj.(T), nil
}

func GetSetterDefinitionT[T any](ctn Container) *Def {
	reflectType := GetInterfaceReflectType((*T)(nil))
	def := ctn.GetDefinitionByType(reflectType)
	return def
}

func SafeGetManyFromContainerT[T any](ctn Container) ([]T, error) {
	reflectType := GetInterfaceReflectType((*T)(nil))
	objs, err := ctn.SafeGetManyByType(reflectType)
	if err != nil {
		return nil, err
	}
	result := make([]T, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(T))
	}
	return result, nil
}
func GetImplementedNames(implementedTypes ...reflect.Type) string {
	builder := strings.Builder{}
	for idx, implementedType := range implementedTypes {
		builder.WriteString(implementedType.Name())
		if idx < len(implementedTypes)-1 {
			builder.WriteString(", ")
		}
	}
	return builder.String()
}
